package memcache

type StorableItem interface {
	Key() string
	Value() string
	ExpiresIn() int64
}

type storableItem struct {
	key       string
	value     string
	expiresIn int64
}

func (si storableItem) Key() string {
	return si.key
}

func (si storableItem) Value() string {
	return si.value
}

func (si storableItem) ExpiresIn() int64 {
	return si.expiresIn
}

type PersistenStorageAdapter interface {
	NewStorableItem(key, val string, expIn int64) StorableItem
	LoadMemcache() ([]StorableItem, error)
	DumpMemcache([]StorableItem) error
}
