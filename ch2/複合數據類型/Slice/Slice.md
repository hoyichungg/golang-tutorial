## Slice

Slice（切片）代表變長的序列，序列中每個元素都有相同的類型。
一個slice類型一般寫作[]T，其中T代表slice中元素的類型；slice的語法和數組很像，只是沒有固定長度而已。


數組和slice之間有着緊密的聯繫。
一個slice是一個輕量級的數據結構，提供了訪問數組子序列（或者全部）元素的功能，而且slice的底層確實引用一個數組對象。一個slice由三個部分構成：指針、長度和容量

指針指向第一個slice元素對應的底層數組元素的地址，要註意的是slice的第一個元素併不一定就是數組的第一個元素。長度對應slice中元素的數目；長度不能超過容量，容量一般是從slice的開始位置到底層數據的結尾位置。內置的len和cap函數分别返迴slice的長度和容量。

多個slice之間可以共享底層的數據，併且引用的數組部分區間可能重疊

```
months := [...]string{1: "January", /* ... */, 12: "December"}
```

![Slice](https://wizardforcel.gitbooks.io/gopl-zh/content/images/ch4-01.png)

```
Q2 := months[4:7]
summer := months[6:9]
fmt.Println(Q2)     // ["April" "May" "June"]
fmt.Println(summer) // ["June" "July" "August"]
```

兩個slice都包含了六月份，下面的代碼是一個包含相同月份的測試

```
for _, s := range summer {
    for _, q := range Q2 {
        if s == q {
            fmt.Printf("%s appears in both\n", s)
        }
    }
}
```

如果切片操作超出cap(s)的上限將導致一個panic異常，但是超出len(s)則是意味着擴展了slice，因爲新slice的長度會變大：

```
fmt.Println(summer[:20]) // panic: out of range

hendlessSummer := summer[:5] // extend a slice (within capacity)
fmt.Println(endlessSummer)  // "[June July August September October]"
```
字符串的切片操作和[]byte字節類型切片的切片操作是類似的。

它們都寫作x[m:n]，併且都是返迴一個原始字節繫列的子序列，底層都是共享之前的底層數組，因此切片操作對應常量時間複雜度。

x[m:n]切片操作對於字符串則生成一個新字符串，如果x是[]byte的話則生成一個新的[]byte。

slice唯一合法的比較操作是和nil比較，例如：

```
if summer == nil { /* ... */ }
```

一個零值的slice等於nil。一個nil值的slice併沒有底層數組。一個nil值的slice的長度和容量都是0，但是也有非nil值的slice的長度和容量也是0的，例如[]int{}或make([]int, 3)[3:]。與任意類型的nil值一樣，我們可以用[]int(nil)類型轉換表達式來生成一個對應類型slice的nil值。

```
var s []int    // len(s) == 0, s == nil
s = nil        // len(s) == 0, s == nil
s = []int(nil) // len(s) == 0, s == nil
s = []int{}    // len(s) == 0, s != nil
```
如果你需要測試一個slice是否是空的，使用len(s) == 0來判斷，而不應該用s == nil來判斷。

內置的make函數創建一個指定元素類型、長度和容量的slice。容量部分可以省略，在這種情況下，容量將等於長度。

```
make([]T, len)
make([]T, len, cap) // same as make([]T, cap)[:len]
```
在底層，make創建了一個匿名的數組變量，然後返迴一個slice；只有通過返迴的slice才能引用底層匿名的數組變量。

在第一種語句中，slice是整個數組的view。
在第二個語句中，slice隻引用了底層數組的前len個元素，但是容量將包含整個的數組。
額外的元素是留給未來的增長用的。

---

## append函數

內置的append函數用於向slice追加元素：

```
var runes []rune
for _, r := range "Hello, 世界" {
    runes = append(runes, r)
}
fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']"
```
在循環中使用append函數構建一個由九個rune字符構成的slice，當然對應這個特殊的問題我們可以通過Go語言內置的[]rune("Hello, 世界")轉換操作完成。


```
func appendInt(x []int, y int) []int {
    var z []int
    zlen := len(x) + 1
    if zlen <= cap(x) {
        // There is room to grow.  Extend the slice.
        z = x[:zlen]
    } else {
        // There is insufficient space.  Allocate a new array.
        // Grow by doubling, for amortized linear complexity.
        zcap := zlen
        if zcap < 2*len(x) {
            zcap = 2 * len(x)
        }
        z = make([]int, zlen, zcap)
        copy(z, x) // a built-in function; see text
    }
    z[len(x)] = y
    return z
}


func main() {
    var x, y []int
    for i := 0; i < 10; i++ {
        y = appendInt(x, i)
        fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
        x = y
    }
}
```

每一次容量的變化都會導致重新分配內存和copy操作：

```
0  cap=1    [0]
1  cap=2    [0 1]
2  cap=4    [0 1 2]
3  cap=4    [0 1 2 3]
4  cap=8    [0 1 2 3 4]
5  cap=8    [0 1 2 3 4 5]
6  cap=8    [0 1 2 3 4 5 6]
7  cap=8    [0 1 2 3 4 5 6 7]
8  cap=16   [0 1 2 3 4 5 6 7 8]
9  cap=16   [0 1 2 3 4 5 6 7 8 9]
```
讓我們仔細査看i=3次的迭代。當時x包含了[0 1 2]三個元素，但是容量是4，因此可以簡單將新的元素添加到末尾，不需要新的內存分配。然後新的y的長度和容量都是4，並且和x引用着相同的底層數組，如圖4.2所示。

![append](https://wizardforcel.gitbooks.io/gopl-zh/content/images/ch4-02.png)

在下一次迭代時i=4，現在沒有新的空餘的空間了，因此appendInt函數分配一個容量爲8的底層數組，將x的4個元素[0 1 2 3]複製到新空間的開頭，然後添加新的元素i，新元素的值是4。新的y的長度是5，容量是8；後面有3個空閒的位置，三次迭代都不需要分配新的空間。當前迭代中，y和x是對應不同底層數組的view。

![append](https://wizardforcel.gitbooks.io/gopl-zh/content/images/ch4-03.png)

內置的append函數可能使用比appendInt更複雜的內存擴展策略。

內置的append函數則可以追加多個元素，甚至追加一個slice。

```
var x []int
x = append(x, 1)
x = append(x, 2, 3)
x = append(x, 4, 5, 6)
x = append(x, x...) // append the slice x
fmt.Println(x)      // "[1 2 3 4 5 6 1 2 3 4 5 6]"
```
