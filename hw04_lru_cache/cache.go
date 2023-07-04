package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type item struct {
	Key Key
	val interface{}
}

func (lc *lruCache) Set(key Key, value interface{}) bool {
	cacheItem := item{key, value}
	if mapListItem, ok := lc.items[key]; ok {
		mapListItem.Value = cacheItem
		lc.queue.MoveToFront(mapListItem)
		return true
	}

	if lc.queue.Len() == lc.capacity {
		rc := lc.queue.Back()
		lc.queue.Remove(rc)
		delete(lc.items, rc.Value.(item).Key)
	}
	lc.items[key] = lc.queue.PushFront(cacheItem)

	return false
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	if mapListItem, ok := lc.items[key]; ok {
		lc.queue.MoveToFront(mapListItem)
		return mapListItem.Value.(item).val, true
	}
	return nil, false
}

func (lc *lruCache) Clear() {
	lc.queue = NewList()
	lc.items = make(map[Key]*ListItem, lc.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
