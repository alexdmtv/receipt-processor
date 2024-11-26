package memdb

import (
	"sync"
)

type DB struct {
	mu    sync.RWMutex
	store map[string]interface{}
}

func New() *DB {
	return &DB{
		store: make(map[string]interface{}),
	}
}

func (d *DB) Get(key string) (interface{}, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.store[key], nil
}

func (d *DB) Set(key string, value interface{}) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.store[key] = value
	return nil
}

func (d *DB) Delete(key string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.store, key)
	return nil
}

func (d *DB) Close() error {
	return nil
}
