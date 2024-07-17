package kv

import (
	"fmt"
	"sync"
)

type DistributedCache struct {
	mutex       sync.Mutex
	data        map[string][]byte
	valueStored chan<- []byte
}

func NewDistribuetCache() *DistributedCache {
	return &DistributedCache{
		data:        make(map[string][]byte),
		valueStored: make(chan<- []byte),
	}
}

func (d *DistributedCache) Get(key string) []byte {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if value, ok := d.data[key]; ok {
		return value
	}
	return nil
}

func (d *DistributedCache) Set(key string, value []byte) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	d.data[key] = value
	d.valueStored <- value
}

func (d *DistributedCache) Delete(key string) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	delete(d.data, key)
}

func (d *DistributedCache) PrintCache() {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	for key, value := range d.data {
		fmt.Printf("Key = %v - Value = %v", key, value)
	}
}

func (d *DistributedCache) GetAll() {
}
