# go-multibimap

基于go开发，具备多重映射、双向映射、线程安全的Map结构实现

## 快速开始

仅使用双向映射的结构
```
import "go-multibimap"
biMap := multibimap.NewBiMap()
biMap.Put(1, "1111")
biMap.Put(2, "2222")
val, ok := biMap.Get(1) // "1111", true
val, ok = biMap.GetInverse("1111") // 1, true

biMap.Remove(1)
biMap.RemoveInverse("2222")
```
仅使用多重映射的结构
```
import "go-multibimap"
usPresidents := []struct {
    firstName  string
    middleName string
    lastName   string
    termStart  int
    termEnd    int
}{
    {"George", "", "Washington", 1789, 1797},
    {"John", "", "Bush", 1797, 1801},
    {"Thomas", "", "Jefferson", 1801, 1809},
    {"James", "", "Madison", 1809, 1817},
    {"James", "", "Monroe", 1817, 1825},
    {"John", "Quincy", "Adams", 1825, 1829},
    {"John", "", "Tyler", 1841, 1845},
    {"James", "", "Polk", 1845, 1849},
    {"Grover", "", "Cleveland", 1885, 1889},
    {"Benjamin", "", "Harrison", 1889, 1893},
    {"Grover", "", "Cleveland", 1893, 1897},
    {"George", "Herbert Walker", "Bush", 1989, 1993},
    {"George", "Walker", "Bush", 2001, 2009},
    {"Barack", "Hussein", "Obama", 2009, 2017},
}

m := multibimap.NewMultiMap()

for _, president := range usPresidents {
    m.Put(president.firstName, president.lastName)
}

for _, firstName := range m.KeySet() {
    lastNames, _ := m.Get(firstName)
    fmt.Printf("%v: %v\n", firstName, lastNames)
}

// Print
// John: [Bush Adams Tyler]
// Thomas: [Jefferson]
// James: [Madison Monroe Polk]
// Grover: [Cleveland Cleveland]
// Benjamin: [Harrison]
// Barack: [Obama]
// George: [Washington Bush Bush]

```
使用具备多重映射、双向映射的结构
```
import "go-multibimap"
usPresidents := []struct {
    firstName  string
    middleName string
    lastName   string
    termStart  int
    termEnd    int
}{
    {"George", "", "Washington", 1789, 1797},
    {"John", "", "Bush", 1797, 1801},
    {"Thomas", "", "Jefferson", 1801, 1809},
    {"James", "", "Madison", 1809, 1817},
    {"James", "", "Monroe", 1817, 1825},
    {"John", "Quincy", "Adams", 1825, 1829},
    {"John", "", "Tyler", 1841, 1845},
    {"James", "", "Polk", 1845, 1849},
    {"Grover", "", "Cleveland", 1885, 1889},
    {"Benjamin", "", "Harrison", 1889, 1893},
    {"Grover", "", "Cleveland", 1893, 1897},
    {"George", "Herbert Walker", "Bush", 1989, 1993},
    {"George", "Walker", "Bush", 2001, 2009},
    {"Barack", "Hussein", "Obama", 2009, 2017},
}
m := multibimap.NewBiMultiMap()
for _, president := range usPresidents {
    m.Put(president.firstName, president.lastName)
}

fmt.Println(m.Get("George"))
fmt.Println(m.GetInverse("Bush"))

// Print
// [Washington Bush Bush] true
// [John George George] true
```
