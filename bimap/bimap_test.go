/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/4/29 11:27 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package bimap

import (
	"reflect"
	"sync"
	"testing"
)

func TestNewBiMap(t *testing.T) {
	actual := NewBiMap()
	expected := &BiMap{
		RWMutex:   &sync.RWMutex{},
		immutable: false,
		forward:   make(map[interface{}]interface{}),
		inverse:   make(map[interface{}]interface{}),
	}

	if !reflect.DeepEqual(actual.inverse, expected.inverse) || !reflect.DeepEqual(actual.forward, expected.forward) || actual.immutable != expected.immutable {
		t.Errorf("NewBiMap() expected %v, got %v", actual, expected)
	}

}

func TestBiMap_Put(t *testing.T) {

	var tests = []struct{
		testName string
		input map[string]string
	}{
		{
			testName: "case0",
			input: map[string]string{
				"1": "a",
				"2": "b",
				"3": "c",
				"4": "d",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			m := NewBiMap()
			for k, v := range test.input {
				m.Put(k, v)
			}

			if size := m.Size(); size != len(test.input) {
				t.Errorf("Put test case:%s, expected %v, got %v\n", test.testName, 7, size)
			}

			fwdExpected := make(map[interface{}]interface{})
			invExpected := make(map[interface{}]interface{})

			for k, v := range test.input {
				fwdExpected[k] = v
				invExpected[v] = k
			}

			biExpected := &BiMap{
				RWMutex:   &sync.RWMutex{},
				immutable: false,
				forward:   fwdExpected,
				inverse:   invExpected,
			}

			if !reflect.DeepEqual(m.forward, fwdExpected) || !reflect.DeepEqual(m.inverse, invExpected) {
				t.Errorf("Put test case:%s, expected %v, got %v", test.testName, m, biExpected)
			}
		})
	}

}

func TestBiMap_ContainsKey(t *testing.T) {
	var tests = []struct{
		testName string
		input map[string]string
		output map[string]bool
	}{
		{
			testName: "case0",
			input: map[string]string{
				"1": "a",
				"2": "b",
				"3": "c",
				"4": "d",
			},
			output: map[string]bool{
				"1": true,
				"2": true,
				"4": true,
				"5": false,
 			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			m := NewBiMap()
			for k, v := range test.input {
				m.Put(k, v)
			}


			for k, expected := range test.output {
				if actual := m.ContainsKey(k); actual != expected {
					t.Errorf("ContainsKey test case:%s, key=%s, expected %v, got %v", test.testName, k, expected, actual)
				}
			}

		})

	}
}

func TestBiMap_ContainsValue(t *testing.T) {
	var tests = []struct{
		testName string
		input map[string]string
		output map[string]bool
	}{
		{
			testName: "case0",
			input: map[string]string{
				"1": "a",
				"2": "b",
				"3": "c",
				"4": "d",
			},
			output: map[string]bool{
				"a": true,
				"b": true,
				"c": true,
				"e": false,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			m := NewBiMap()
			for k, v := range test.input {
				m.Put(k, v)
			}

			for v, expected := range test.output {
				if actual := m.ContainsValue(v); actual != expected {
					t.Errorf("ContainsValue test case:%s, key=%s, expected %v, got %v", test.testName, v, expected, actual)
				}
			}

		})
	}
}

func TestBiMap_Get(t *testing.T) {
	var tests = []struct{
		testName string
		input map[string]string
		output map[string][]interface{}
	}{
		{
			testName: "case0",
			input: map[string]string{
				"1": "a",
				"2": "b",
				"3": "c",
				"4": "d",
			},
			output: map[string][]interface{}{
				"1": []interface{}{"a", true},
				"2": []interface{}{"b", true},
				"6": []interface{}{nil, false},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			m := NewBiMap()
			for k, v := range test.input {
				m.Put(k, v)
			}

			for key, expected := range test.output {
				if actualV, ok := m.Get(key); actualV != expected[0] || ok != expected[1] {
					t.Errorf("Get test case:%s, expected %v, got %v", test.testName, expected, []interface{}{actualV, ok})
				}
			}

		})
	}
}


func BenchmarkBiMap_Put100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := NewBiMap()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkPut100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := make(map[interface{}]interface{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m[n] = struct{}{}
		}
	}
}

func BenchmarkBiMap_Put1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := NewBiMap()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}
func BenchmarkPut1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := make(map[interface{}]interface{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m[n] = struct{}{}
		}
	}
}

func BenchmarkBiMap_Put10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := NewBiMap()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}
func BenchmarkPut10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := make(map[interface{}]interface{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m[n] = struct{}{}
		}
	}
}

func BenchmarkBiMap_Put100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := NewBiMap()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkPut100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := make(map[interface{}]interface{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m[n] = struct{}{}
		}
	}
}

func benchmarkPut(b *testing.B, m *BiMap, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Put(n, struct{}{})
		}
	}
}


