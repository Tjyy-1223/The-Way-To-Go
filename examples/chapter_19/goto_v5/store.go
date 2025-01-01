package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/rpc"
	"os"
	"sync"
)

const saveQueueLength = 1000

type Store interface {
	Put(url, key *string) error
	Get(key, url *string) error
}

type ProxyStore struct {
	urls   *URLStore
	client *rpc.Client
}

type URLStore struct {
	urls map[string]string
	mu   sync.RWMutex
	save chan record
}

type record struct {
	Key, Url string
}

func NewURLStore(filename string) *URLStore {
	s := &URLStore{
		urls: make(map[string]string),
	}
	if filename != "" {
		s.save = make(chan record, saveQueueLength)
		if err := s.load(filename); err != nil {
			log.Fatal("Error loading URLStore:", err)
		}
		go s.saveLoop(filename)
	}
	return s
}

func (u *URLStore) Get(key, url *string) error {
	u.mu.RLock()
	defer u.mu.RUnlock()
	println("master get the key", *key)
	if r, ok := u.urls[*key]; ok {
		*url = r
		return nil
	}
	return errors.New("not Found Key")
}

func (u *URLStore) Set(key, url *string) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	if _, present := u.urls[*key]; present {
		return errors.New("key already exists")
	}
	u.urls[*key] = *url
	return nil
}

func (u *URLStore) Put(url, key *string) error {
	for {
		*key = genKey(u.Count()) // generate the short URL
		if err := u.Set(key, url); err == nil {
			break
		}
	}
	if u.save != nil {
		u.save <- record{*key, *url}
	}
	return nil
}

func (u *URLStore) Count() int {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return len(u.urls)
}

func (u *URLStore) load(filename string) error {
	var err error
	var f *os.File
	f, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println("Error opening URLStore:", err)
		return err
	}
	defer f.Close()

	if _, err := f.Seek(0, 0); err != nil {
		return err
	}
	d := json.NewDecoder(f)
	for err == nil {
		var r record
		if err = d.Decode(&r); err == nil {
			u.Set(&r.Key, &r.Url)
		}
	}
	if err == io.EOF {
		return nil
	}
	log.Println("Error decoding URLStore:", err) // map hasn't been read correctly
	return err
}

func (u *URLStore) saveLoop(filename string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Error opening URLStore: ", err)
	}
	defer f.Close()
	e := json.NewEncoder(f)
	for {
		r := <-u.save
		if err := e.Encode(r); err != nil {
			log.Println("Error saving to URLStore: ", err)
		}
	}
}

func NewProxyStore(addr string) *ProxyStore {
	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Println("Error constructing ProxyStore:", err)
	}
	return &ProxyStore{urls: NewURLStore(""), client: client}
}

func (p *ProxyStore) Get(key, url *string) error {
	if err := p.urls.Get(key, url); err == nil {
		return nil
	}
	if err := p.client.Call("Store.Get", key, url); err != nil {
		return err
	}
	// store the url
	p.urls.Set(key, url)
	return nil
}

func (p *ProxyStore) Put(url, key *string) error {
	if err := p.client.Call("Store.Put", url, key); err != nil {
		return err
	}
	p.urls.Set(key, url)
	return nil
}
