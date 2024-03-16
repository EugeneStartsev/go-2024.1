package lrucache

import (
	"container/list"
	"time"
)

type lruCache struct {
	capacity int
	items    map[int]*list.Element
	queue    *list.List
	time     time.Time
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

	l.queue.MoveToFront(elem)
	l.time = time.Now()
	return elem.Value.(int), true
}

func (l *lruCache) Set(key, value int) {
	if elem, isExist := l.items[key]; isExist == true {
		l.queue.MoveToFront(elem)
		elem.Value = value
		l.time = time.Now()
	}

	if l.queue.Len() == l.capacity {
		l.purge()
	}

	elem := l.queue.PushFront(value)
	l.items[key] = elem
	l.time = time.Now()
}

func (l *lruCache) Range(f func(key, value int) bool) {
	for key, value := range l.items {
		f(key, value.Value.(int))
	}
}

func (l *lruCache) Clear() {
	for len(l.items) > 0 {
		l.purge()
	}
}

func (l *lruCache) purge() {
	if elem := l.queue.Back(); elem != nil {
		item := l.queue.Remove(elem).(int)
		delete(l.items, item)
	}
}
