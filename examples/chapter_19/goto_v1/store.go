package main

import "sync"

type URLStore struct {
	urls map[string]string
	mu   sync.RWMutex
}

func NewURLStore() *URLStore {
	return &URLStore{urls: make(map[string]string)}
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
			return key
		}
	}
	// shouldn't get here
	return ""
}
