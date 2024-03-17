package lrucache

import (
	"container/list"
	"sort"
	"time"
)

type Item struct {
	key       int
	value     int
	timestamp time.Time
}

type lruCache struct {
	capacity int
	items    map[int]*list.Element
	queue    *list.List
}

func New(cap int) Cache {
	return &lruCache{
		capacity: cap,
		items:    make(map[int]*list.Element),
		queue:    list.New(),
	}
}

func (l *lruCache) Get(key int) (int, bool) {
	elem, isExist := l.items[key]

	if isExist == false {
		return 0, false
	}

	elem.Value.(*Item).timestamp = time.Now()
	l.queue.MoveToFront(elem)
	return elem.Value.(*Item).value, true
}

func (l *lruCache) Set(key, value int) {
	if l.capacity > 0 {
		if elem, isExist := l.items[key]; isExist == true {
			l.queue.MoveToFront(elem)
			elem.Value.(*Item).value = value
			elem.Value.(*Item).timestamp = time.Now()
		}

		if l.queue.Len() == l.capacity {
			l.purge()
		}

		item := &Item{
			key:       key,
			value:     value,
			timestamp: time.Now(),
		}

		elem := l.queue.PushFront(item)
		l.items[key] = elem
	}
}

func (l *lruCache) Range(f func(key, value int) bool) {
	sortedStruct := make([]*Item, 0, len(l.items))

	for key, val := range l.items {
		sortedStruct = append(sortedStruct, &Item{
			key:       key,
			value:     val.Value.(*Item).value,
			timestamp: val.Value.(*Item).timestamp,
		})
	}

	sort.Slice(sortedStruct, func(i, j int) bool {
		return sortedStruct[i].timestamp.Before(sortedStruct[j].timestamp)
	})

	for _, item := range sortedStruct {
		f(item.key, item.value)
	}
}

func (l *lruCache) Clear() {
	for len(l.items) > 0 {
		l.purge()
	}
}

func (l *lruCache) purge() {
	if elem := l.queue.Back(); elem != nil {
		item := l.queue.Remove(elem).(*Item)
		delete(l.items, item.key)
	}
}
