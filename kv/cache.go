package kv

import (
	"fmt"
	"sync"
)

type DistributedCache struct {
	mutex       sync.Mutex
	data        map[int][]byte
	valueStored chan<- []byte
}

func NewDistribuetCache() *DistributedCache {
	return &DistributedCache{
		data:        make(map[int][]byte),
		valueStored: make(chan<- []byte),
	}
}

func (d *DistributedCache) Get(key int) []byte {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if value, ok := d.data[key]; ok {
		return value
	}
	return nil
}

func (d *DistributedCache) Set(key int, value []byte) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	d.data[key] = value
	d.valueStored <- value
}

func (d *DistributedCache) Delete(key int) {
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
