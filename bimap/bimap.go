/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/4/26 3:10 下午, by lishanlei, create
*/

/*
DESCRIPTION
BiMap提供了一个双向映射、线程安全的Map结构
*/

package bimap

import "sync"

// BiMap 是一个具备双向映射的map结构
type BiMap struct {
	*sync.RWMutex
	immutable bool
	forward   map[interface{}]interface{}
	inverse   map[interface{}]interface{}
}

// NewBiMap return a empty
func NewBiMap() *BiMap {
	return &BiMap{
		&sync.RWMutex{},
		false,
		make(map[interface{}]interface{}),
		make(map[interface{}]interface{}),
	}
}

// Put 将一个key-value放入双向映射map中
func (b *BiMap) Put(key interface{}, value interface{}) {
	b.RLock()
	if b.immutable {
		panic("Cannot modify immutable map")
	}
	b.RUnlock()

	b.Lock()
	defer b.Unlock()
	b.forward[key] = value
	b.inverse[value] = key
}

// ContainsKey 是否包含一个key
func (b *BiMap) ContainsKey(key interface{}) (found bool) {
	b.RLock()
	defer b.RUnlock()
	_, found = b.forward[key]
	return
}

// ContainsValue 是否包含一个value
func (b *BiMap) ContainsValue(value interface{}) (found bool) {
	b.RLock()
	defer b.RUnlock()
	_, found = b.inverse[value]
	return
}

// Get 返回给定键的值，以及该元素是否存在
func (b *BiMap) Get(key interface{}) (value interface{}, found bool) {
	b.RLock()
	defer b.RUnlock()
	value, found = b.forward[key]
	return
}

// GetInverse 返回BiMap中给定值的键以及该元素是否存在
func (b *BiMap) GetInverse(value interface{}) (interface{}, bool) {
	b.RLock()
	defer b.RUnlock()
	key, found := b.inverse[value]
	return key, found
}

// Remove 根据key删除对应的键值对，如果不存在则返回
func (b *BiMap) Remove(key interface{}) {
	b.RLock()
	if b.immutable {
		panic("Cannot modify immutable map")
	}
	b.RUnlock()

	if !b.ContainsKey(key) {
		return
	}

	value, _ := b.Get(key)
	b.Lock()
	defer b.Unlock()
	delete(b.forward, key)
	delete(b.inverse, value)
}

// RemoveInverse 根据给定的值删除对应的键值对，如果不存在则返回
func (b *BiMap) RemoveInverse(value interface{}) {
	b.RLock()
	if b.immutable {
		panic("Cannot modify immutable map")
	}
	b.RUnlock()

	if !b.ContainsValue(value) {
		return
	}
	b.Lock()
	defer b.Unlock()
	key, _ := b.GetInverse(value)
	delete(b.inverse, value)
	delete(b.forward, key)
}

// Size 返回map的长度
func (b *BiMap) Size() int {
	b.RLock()
	defer b.RUnlock()
	return len(b.forward)
}

// MakeImmutable 冻结BiMap，之后如果写入会返回panic
func (b *BiMap) MakeImmutable() {
	b.Lock()
	defer b.Unlock()
	b.immutable = true
}

// UnMakeImmutable 解除冻结
func (b *BiMap) UnMakeImmutable() {
	b.Lock()
	defer b.Unlock()
	b.immutable = false
}





