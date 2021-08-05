/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/4/26 2:42 下午, by lishanlei, create
*/

/*
DESCRIPTION
BiMultiMap提供了一个多重映射、双向映射、线程安全的Map结构
*/

package bimultimap

import (
	"icode.baidu.com/baidu/personal-code/go-multibimap/multimap"
	"sync"
)


// BiMultiMap 是一个具备多重映射，双向映射的map结构
type BiMultiMap struct {
	immutable bool
	forward *multimap.MultiMap
	inverse *multimap.MultiMap
	*sync.RWMutex
}

// NewBiMultiMap 返回一个BiMultiMap结构
func NewBiMultiMap() *BiMultiMap {
	return &BiMultiMap{
		false,
		multimap.NewMultiMap(),
		multimap.NewMultiMap(),
		&sync.RWMutex{},
	}
}

// Put 将一对key-value写入map
func (b *BiMultiMap) Put(key, value interface{}) {
	b.RLock()
	if b.immutable {
		panic("Cannot modify immutable map")
	}
	b.RUnlock()
	b.forward.Put(key, value)
	b.inverse.Put(value, key)

}

// PutAll 将一个key对应的value列表写入map
func (b *BiMultiMap) PutAll(key interface{}, value []interface{}) {
	b.RLock()
	if b.immutable {
		panic("Cannot modify immutable map")
	}
	b.RUnlock()
	b.forward.PutAll(key, value)
	for _, v := range value {
		b.inverse.Put(v, key)
	}
}

// Get 返回给定键的所有值，以及该元素是否存在
func (b *BiMultiMap) Get(key interface{}) ([]interface{}, bool) {
	return b.forward.Get(key)
}

// GetInverse 返回map中给定值的所有键以及该元素是否存在
func (b *BiMultiMap) GetInverse(v interface{}) ([]interface{}, bool) {
	return b.inverse.Get(v)
}

// ContainsKey 如果map包含指定key，则返回true。
func (b *BiMultiMap) ContainsKey(key interface{}) bool {
	return b.forward.ContainsKey(key)
}

// ContainsValue 如果map包含至少一个指定值的键-值对，则返回true。
func (b *BiMultiMap) ContainsValue(value interface{}) bool {
	return b.inverse.ContainsKey(value)
}

// Remove 从map中删除一对key-value 如果不存在则返回
func (b *BiMultiMap) Remove(key interface{}, value interface{}) {
	b.RLock()
	if b.immutable {
		panic("Cannot modify immutable map")
	}
	b.RUnlock()

	if !b.ContainsKey(key) {
		return
	}
	b.forward.Remove(key, value)
	b.inverse.Remove(value, key)
}





