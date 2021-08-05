/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/4/26 1:21 下午, by lishanlei, create
*/

/*
DESCRIPTION
MultiMap提供了一个多重映射、线程安全的Map结构
*/

package multimap

import "sync"

// Entry 多重映射map中的一个key-value结构.
type Entry struct {
	Key   interface{}
	Value interface{}
}

//  MultiMap 是一个具备多重映射的map结构
type MultiMap struct {
	m map[interface{}][]interface{}
	*sync.RWMutex
}

// NewMultiMap 返回一个初始化multiMap
func NewMultiMap() *MultiMap {
	return &MultiMap{map[interface{}][]interface{}{}, &sync.RWMutex{}}
}

// Get 返回给定键的所有值，以及该元素是否存在
func (m *MultiMap) Get(key interface{}) (value []interface{}, found bool) {
	m.RLock()
	defer m.RUnlock()
	value, found = m.m[key]
	return
}

// Put 将一对key-value写入map
func (m *MultiMap) Put(key interface{}, value interface{}) {
	m.Lock()
	defer m.Unlock()
	m.m[key] = append(m.m[key], value)
}

// PutAll 将一个key对应的value列表写入map
func (m *MultiMap) PutAll(key interface{}, values []interface{}) {
	for _, value := range values {
		m.Put(key, value)
	}
}

// Remove 从map中删除一对key-value 如果不存在则返回
func (m *MultiMap) Remove(key interface{}, value interface{}) {
	m.Lock()
	defer m.Unlock()
	values, found := m.m[key]
	if found {
		for i, v := range values {
			if v == value {
				m.m[key] = append(values[:i], values[i+1:]...)
			}
		}
	}
	if 0 == len(m.m[key]) {
		delete(m.m, key)
	}
}

// RemoveAll 从map中删除key对应的所有key-value
func (m *MultiMap) RemoveAll(key interface{}) {
	m.Lock()
	defer m.Unlock()
	delete(m.m, key)
}

// Contains 如果map包含至少一个键值对以及键值和值，则返回true
func (m *MultiMap) Contains(key interface{}, value interface{}) bool {
	m.RLock()
	defer m.RUnlock()
	values, found := m.m[key]

	for _, v := range values {
		if v == value {
			return found
		}
	}
	return false
}

// ContainsKey 如果map包含指定key，则返回true。
func (m *MultiMap) ContainsKey(key interface{}) (found bool) {
	m.RLock()
	defer m.RUnlock()
	_, found = m.m[key]
	return
}

// ContainsValue 如果map包含至少一个指定值的键-值对，则返回true。
func (m *MultiMap) ContainsValue(value interface{}) bool {
	m.RLock()
	defer m.RUnlock()
	for _, values := range m.m {
		for _, v := range values {
			if v == value {
				return true
			}
		}
	}
	return false
}

// Entries 返回此Multimap中包含的所有键值对的集合。
func (m *MultiMap) Entries() []*Entry {
	count := 0
	entries := make([]*Entry, m.Size())
	m.RLock()
	defer m.RUnlock()
	for key, values := range m.m {
		for _, value := range values {
			entries[count] = &Entry{Key: key, Value: value}
			count++
		}
	}
	return entries
}

// Keys 返回一个集合，该集合包含此map中每个键值对的键。
// 不会折叠重复的键
func (m *MultiMap) Keys() []interface{} {
	keys := make([]interface{}, m.Size())
	count := 0
	m.RLock()
	defer m.RUnlock()
	for key, values := range m.m {
		for range values {
			keys[count] = key
			count++
		}
	}
	return keys
}

// KeySet 返回一个集合，该集合包含此map中包含的所有不同键。
func (m *MultiMap) KeySet() []interface{} {
	keys := make([]interface{}, len(m.m))
	count := 0
	m.RLock()
	defer m.RUnlock()
	for key := range m.m {
		keys[count] = key
		count++
	}
	return keys
}

// Values 返回map中包含的每个键值对的所有值。
// 不会折叠重复的值
func (m *MultiMap) Values() []interface{} {
	values := make([]interface{}, m.Size())
	count := 0
	m.RLock()
	defer m.RUnlock()
	for _, vs := range m.m {
		for _, v := range vs {
			values[count] = v
			count++
		}
	}
	return values
}

// Clear 删除map
func (m *MultiMap) Clear() {
	m.Lock()
	defer m.Unlock()
	m.m = make(map[interface{}][]interface{})
}

// Empty 判断map是否为空
func (m *MultiMap) Empty() bool {
	return m.Size() == 0
}

// Size 返回map中的键/值对的数量
func (m *MultiMap) Size() int {
	size := 0
	m.RLock()
	defer m.RUnlock()
	for _, values := range m.m {
		size += len(values)
	}
	return size
}
