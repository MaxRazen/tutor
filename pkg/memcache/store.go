package memcache

import (
	"log"
	"sync"
	"sync/atomic"
	"time"
)

var instance store

const (
	DumpInterval = time.Second * 30
)

type storeItem struct {
	value     string
	expiresIn int64
}

type store struct {
	data     map[string]storeItem
	adapter  PersistenStorageAdapter
	mu       sync.Mutex
	dumpedAt int64
	dumping  atomic.Bool
}

func Init(pse PersistenStorageAdapter) {
	storableItems, err := pse.LoadMemcache()

	if err != nil {
		log.Println("memcache/init:", err)
	}

	data := deserialize(storableItems)

	instance = store{
		adapter:  pse,
		data:     data,
		mu:       sync.Mutex{},
		dumpedAt: time.Now().Unix(),
		dumping:  atomic.Bool{},
	}
	instance.dumping.Store(false)
}

func Get(key string) (string, bool) {
	v, ok := instance.data[key]
	if !ok {
		return "", false
	}
	if v.expiresIn > time.Now().Unix() {
		dump()
		return v.value, true
	}
	removeKey(key)
	dump()
	return "", true
}

func Set(key, value string, ttl time.Duration) {
	v := storeItem{
		value:     value,
		expiresIn: time.Now().Unix() + int64(ttl.Seconds()),
	}
	instance.mu.Lock()
	instance.data[key] = v
	instance.mu.Unlock()
	dump()
}

func removeKey(key string) {
	if _, ok := instance.data[key]; ok {
		instance.mu.Lock()
		delete(instance.data, key)
		instance.mu.Unlock()
	}
}

func dump() {
	if instance.dumpedAt > time.Now().Add(-DumpInterval).Unix() || instance.dumping.Load() {
		return
	}

	instance.dumping.Store(true)
	defer instance.dumping.Store(false)

	data := serialize(instance.data)
	err := instance.adapter.DumpMemcache(data)

	if err != nil {
		log.Println("memcache/dump:", err)
	}

	instance.dumpedAt = time.Now().Unix()
}

func serialize(kv map[string]storeItem) []StorableItem {
	items := make([]StorableItem, 0, len(kv))

	for k, v := range kv {
		items = append(items, instance.adapter.NewStorableItem(
			k,
			v.value,
			v.expiresIn,
		))
	}

	return items
}

func deserialize(storableItems []StorableItem) map[string]storeItem {
	m := make(map[string]storeItem)
	now := time.Now().Unix()

	for _, si := range storableItems {
		if si.ExpiresIn() <= now {
			continue
		}

		m[si.Key()] = storeItem{
			value:     si.Value(),
			expiresIn: si.ExpiresIn(),
		}
	}

	return m
}
