## Map

哈希表是一種巧妙併且實用的數據結構。

它是一個無序的key/value對的集合，其中所有的key都是不同的，然後通過給定的key可以在常數時間複雜度內檢索、更新或刪除對應的value。

在Go語言中，一個map就是一個哈希表的引用，map類型可以寫爲map[K]V，其中K和V分别對應key和value。map中所有的key都有相同的類型，所以的value也有着相同的類型，但是key和value之間可以是不同的數據類型。

內置的make函數可以創建一個map：

```
ages := make(map[string]int) // mapping from strings to ints
```
我們也可以用map字面值的語法創建map，同時還可以指定一些最初的key/value：

```
ages := map[string]int{
    "alice":   31,
    "charlie": 34,
}
```
這相當於

```
ages := make(map[string]int)
ages["alice"] = 31
ages["charlie"] = 34
```

另一種創建空的map的表達式是map[string]int{}。

Map中的元素通過key對應的下標語法訪問：

```
ages["alice"] = 32
fmt.Println(ages["alice"]) // "32"
```
使用內置的delete函數可以刪除元素：

```
delete(ages, "alice") // remove element ages["alice"]
```

所有這些操作是安全的，卽使這些元素不在map中也沒有關繫；
如果一個査找失敗將返迴value類型對應的零值，例如，卽使map中不存在“bob”下面的代碼也可以正常工作，因爲ages["bob"]失敗時將返迴0。

```
ages["bob"] = ages["bob"] + 1 // happy birthday!
```
x += y和x++等簡短賦值語法也可以用在map上，所以上面的代碼可以改寫成

```
ages["bob"] += 1

ages["bob"]++
```
map中的元素併不是一個變量，因此我們不能對map的元素進行取址操作：

```
_ = &ages["bob"] // compile error: cannot take address of map element
```
禁止對map元素取址的原因是map可能隨着元素數量的增長而重新分配更大的內存空間，從而可能導致之前的地址無效。

要想遍歷map中全部的key/value對的話，可以使用range風格的for循環實現，和之前的slice遍歷語法類似。下面的迭代語句將在每次迭代時設置name和age變量，它們對應下一個鍵/值對：

```
for name, age := range ages {
    fmt.Printf("%s\t%d\n", name, age)
}
```
Map的迭代順序是不確定的，併且不同的哈希函數實現可能導致不同的遍歷順序
可以使用sort包的Strings函數對字符串slice進行排序。下面是常見的處理方式：

```
import "sort"

var names []string
for name := range ages {
    names = append(names, name)
}
sort.Strings(names)
for _, name := range names {
    fmt.Printf("%s\t%d\n", name, ages[name])
}
```
因爲我們一開始就知道names的最終大小，因此給slice分配一個合適的大小將會更有效。下面的代碼創建了一個空的slice，但是slice的容量剛好可以放下map中全部的key：

```
names := make([]string, 0, len(ages))
```
在第二個循環中，我們只關心names中的名字，所以我們使用“_”空白標識符來忽略第一個循環變量，也就是迭代slice時的索引。

map類型的零值是nil，也就是沒有引用任何哈希表。

```
var ages map[string]int
fmt.Println(ages == nil)    // "true"
fmt.Println(len(ages) == 0) // "true"
``` 
map上的大部分操作，包括査找、刪除、len和range循環都可以安全工作在nil值的map上，它們的行爲和一個空的map類似。但是向一個nil值的map存入元素將導致一個panic異常：

```
ages["carol"] = 21 // panic: assignment to entry in nil map
```
通過key作爲索引下標來訪問map將産生一個value。

如果key在map中是存在的，那麽將得到與key對應的value；如果key不存在，那麽將得到value對應類型的零值，正如我們前面看到的ages["bob"]那樣。這個規則很實用，但是有時候可能需要知道對應的元素是否眞的是在map之中。

例如，如果元素類型是一個數字，你可以需要區分一個已經存在的0，和不存在而返迴零值的0，可以像下面這樣測試：

```
age, ok := ages["bob"]
if !ok { /* "bob" is not a key in this map; age == 0. */ }
```
會經常看到將這兩個結合起來使用，像這樣：

```
if age, ok := ages["bob"]; !ok { /* ... */ }
```
在這種場景下，map的下標語法將産生兩個值；第二個是一個布爾值，用於報告元素是否眞的存在。
布爾變量一般命名爲ok，特别適合馬上用於if條件判斷部分。

和slice一樣，map之間也不能進行相等比較；唯一的例外是和nil進行比較。
要判斷兩個map是否包含相同的key和value

```
func equal(x, y map[string]int) bool {
    if len(x) != len(y) {
        return false
    }
    for k, xv := range x {
        if yv, ok := y[k]; !ok || yv != xv {
            return false
        }
    }
    return true
}
```
註意我們是如何用!ok來區分元素缺失和元素不同的。
我們不能簡單地用xv != y[k]判斷，那樣會導致在判斷下面兩個map時産生錯誤的結果

```
// True if equal is written incorrectly.
equal(map[string]int{"A": 0}, map[string]int{"B": 42})
```

Go語言中併沒有提供一個set類型，但是map中的key也是不相同的，可以用map實現類似set的功能。
下面的dedup程序讀取多行輸入，但是只打印第一次出現的行。（它是1.3節中出現的dup程序的變體。）
dedup程序通過map來表示所有的輸入行所對應的set集合，以確保已經在集合存在的行不會被重複打印。

```
func main() {
    seen := make(map[string]bool) // a set of strings
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        line := input.Text()
        if !seen[line] {
            seen[line] = true
            fmt.Println(line)
        }
    }

    if err := input.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
        os.Exit(1)
    }
}
```
