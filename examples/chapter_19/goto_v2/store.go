package main

import (
	"encoding/gob"
	"io"
	"log"
	"os"
	"sync"
)

type URLStore struct {
	urls map[string]string
	mu   sync.RWMutex
	file *os.File
}

type record struct {
	Key, Url string
}

func NewURLStore(filename string) *URLStore {
	s := &URLStore{urls: make(map[string]string)}
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Error opening URLStore:", err)
	}
	s.file = f
	if err = s.load(); err != nil {
		log.Fatal("Error loading URLStore:", err)
	}
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
			if err := u.save(key, url); err != nil {
				log.Println("fail to save url into urlstore:", err)
			}
			return key
		}
	}
	panic("shouldn't get here")
}

func (u *URLStore) load() error {
	if _, err := u.file.Seek(0, 0); err != nil {
		return err
	}
	d := gob.NewDecoder(u.file)
	var err error
	if err == nil {
		var r record
		if err = d.Decode(&r); err == nil {
			u.Set(r.Key, r.Url)
		}
	}

	if err == io.EOF {
		return nil
	}
	return err
}

func (u *URLStore) save(key, url string) error {
	e := gob.NewEncoder(u.file)
	return e.Encode(record{key, url})
}
