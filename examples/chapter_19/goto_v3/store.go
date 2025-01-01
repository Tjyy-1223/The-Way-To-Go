package main

import (
	"encoding/gob"
	"io"
	"log"
	"os"
	"sync"
)

const saveQueueLength = 1000

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
		save: make(chan record, saveQueueLength),
	}
	if err := s.load(filename); err != nil {
		log.Fatal("Error loading URLStore:", err)
	}
	go s.saveLoop(filename)
	return s
}

func (u *URLStore) Get(key string) string {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.urls[key]
}

func (u *URLStore) Set(key, url string) bool {
	u.mu.Lock()
	defer u.mu.Unlock()
	if _, present := u.urls[key]; present {
		return false
	}
	u.urls[key] = url
	return true
}

func (u *URLStore) Count() int {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return len(u.urls)
}

func (u *URLStore) Put(url string) string {
	for {
		key := genKey(u.Count()) // generate the short URL
		if ok := u.Set(key, url); ok {
			u.save <- record{key, url}
			return key
		}
	}
	panic("shouldn't get here")
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
	d := gob.NewDecoder(f)
	for err == nil {
		var r record
		if err = d.Decode(&r); err == nil {
			u.Set(r.Key, r.Url)
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
	e := gob.NewEncoder(f)
	for {
		r := <-u.save
		if err := e.Encode(r); err != nil {
			log.Println("Error saving to URLStore: ", err)
		}
	}
}
