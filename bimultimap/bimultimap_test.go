/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/4/29 1:52 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package bimultimap

import "testing"

func BenchmarkBiMultiMap_Put100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := NewBiMultiMap()
	b.StartTimer()
	benchmarkPut(b, m, size)
}


func BenchmarkBiMultiMap_Put1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := NewBiMultiMap()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkBiMultiMap_Put10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := NewBiMultiMap()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkBiMultiMap_Put100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := NewBiMultiMap()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkBiMultiMap_PutAll100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := NewBiMultiMap()
	b.StartTimer()
	benchmarkPut(b, m, size)
}


func BenchmarkBiMultiMap_PutAll1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := NewBiMultiMap()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPutAll(b, m, size)
}

func BenchmarkBiMultiMap_PutAll10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := NewBiMultiMap()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPutAll(b, m, size)
}

func BenchmarkBiMultiMap_PutAll100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := NewBiMultiMap()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPutAll(b, m, size)
}

func BenchmarkBiMultiMap_Get100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := NewBiMultiMap()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkBiMultiMap_Get1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := NewBiMultiMap()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}


func BenchmarkBiMultiMap_Get10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := NewBiMultiMap()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkBiMultiMap_Get100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := NewBiMultiMap()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}


func benchmarkGet(b *testing.B, m *BiMultiMap, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Get(n)
		}
	}
}

func benchmarkPut(b *testing.B, m *BiMultiMap, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Put(n, struct{}{})
		}
	}
}

func benchmarkPutAll(b *testing.B, m *BiMultiMap, size int) {
	v := make([]interface{}, 0)
	v = append(v, struct{}{})
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.PutAll(n, v)
		}
	}
}



