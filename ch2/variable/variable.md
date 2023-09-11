## 變量 

var聲明語句可以創建一個特定類型的變量，然後給變量附加一個名字，併且設置變量的初始值。變量聲明的一般語法如下：

例如：

```
var 變量名字 類型 = 表達式
```
**類型**或**表達式**兩個部分可以省略其中的一個。

推導類型: 如果你省略了變量的類型信息，Go 會根據初始化表達式自動推導其類型。

例如：

```
var x = 10  // x 被推導為 int 類型
```
零值初始化: 如果你省略了初始化表達式，Go 會使用對應類型的零值來初始化變量。

這些零值包括：

數值類型的零值是 0
布爾類型的零值是 false
字符串類型的零值是空字符串 ""
slice、map、chan和函數的零值是 nil
數組和結構體的零值是其每個元素或字段的零值

例如：

```
var y int    // y 初始化為 0
var z string // z 初始化為 ""
```
這種零值初始化機制確保每一個在 Go 中聲明的變量都會有一個明確的初始值，因此不存在未初始化的變量。這不僅簡化了代碼的編寫，還確保了程序在各種邊界情況下都能正確、安全地運行。

也可以在一個聲明語句中同時聲明一組變量，或用一組初始化表達式聲明併初始化一組變量。如果省略每個變量的類型，將可以聲明多個類型不同的變量（類型由初始化表達式推導）

例如：

```
var i, j, k int // int, int, int
var b, f, s = true, 2.3, "four" // bool, float64, string 
```

一組變量也可以通過調用一個函數，由函數返迴的多個返迴值初始化

例如：

```
var f, err = os.Open(name) // os.Open returns a file and an error
```

---

## 簡短變量聲明

Go 語言的函數內部，提供了一種簡便的方式來聲明並初始化局部變量，這種方式稱為簡短變量聲明

例如：

```
名字 := 表達式
```

這裡的 := 是簡短變量聲明的關鍵。使用此語法，Go 會自動推導出 名字 的類型，基於 表達式 的類型。

例如：

```
anim := gif.GIF{LoopCount: nframes}
freq := rand.Float64() * 3.0
t := 0.0
```

因爲簡潔和靈活的特點，簡短變量聲明被廣泛用於大部分的局部變量的聲明和初始化。

var形式的聲明語句往往是用於需要顯式指定變量類型地方，或者因爲變量稍後會被重新賦值而初始值無關緊要的地方。

例如：

```
i := 100 // an int
var boiling float64 = 100 // a float64
var names []string
var err error
var p Point
```

和var形式聲明變語句一樣，簡短變量聲明語句也可以用來聲明和初始化一組變量

例如：

```
i, j := 0, 1
```

右邊各個的表達式值賦值給左邊對應位置的各個變量

例如：

```
i, j = j, i // 交換 i 和 j 的值
```
和普通var形式的變量聲明語句一樣，簡短變量聲明語句也可以用函數的返迴值來聲明和初始化變量，像下面的os.Open函數調用將返迴兩個值

例如：

```
f, err := os.Open(name)
if err != nil {
    return err
}
// ...use f...
f.Close()
```

---

## 指針

1. 變量與內存:
    
   - 每個變量都有一塊對應的內存空間來保存其值。
      
   - 變量名（如 x）提供了訪問這些值的方式。但有些變量可以透過表達式（如 x[i] 或 x.f）來存取。

   - 這些表達式大部分時間都用於讀取值，但在賦值時它們可以用來設定新值。

2. 指針的概念:

    - 指針保存了另一個變量的內存地址。

    - 每個變量都有一個唯一的內存地址，但不是每個值都有。

    - 通過指針，可以間接地讀取或更新變量的值，即使我們不知道變量的名稱。

3. 指針的操作:

   - 使用 var x int 聲明一個整數變量 x 之後，&x 會獲取該變量的內存地址，結果是一個指向 int 的指針，類型為 *int。
   
   - 如果有一個指針 p 保存了 x 的地址，我們可以說 "p 指向 x" 或 "p 保存了 x 的地址"。

   - *p 可以用來訪問或修改 p 指向的變量的值。


```
x := 1
p := &x         // p, of type *int, points to x
fmt.Println(*p) // "1"
*p = 2          // equivalent to x = 2
fmt.Println(x)  // "2"
```

變量有時候被稱爲可尋址的值。卽使變量由表達式臨時生成，那麽表達式也必鬚能接受&取地址操作。

任何類型的指針的零值都是nil。如果p != nil測試爲眞，那麽p是指向某個有效變量。指針之間也是可以進行相等測試的，隻有當它們指向同一個變量或全部是nil時才相等。


```
var x, y int
fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"
```
Go語言中，返迴函數中局部變量的地址也是安全的。例如下面的代碼，調用f函數時創建局部變量v，在局部變量地址被返迴之後依然有效，因爲指針p依然引用這個變量。

```
var p = f()

func f() *int {
    v := 1
    return &v
}
```
每次調用f函數都將返迴不同的結果：

```
fmt.Println(f() == f()) // "false"
```
通過指針來更新變量的值，然後返迴更新後的值，可用在一個表達式中

```
func incr(p *int) int {
    *p++ // 非常重要：隻是增加p指向的變量的值，併不改變p指針！！！
    return *p
}

v := 1
incr(&v)              // side effect: v is now 2
fmt.Println(incr(&v)) // "3" (and v is 3)
```

---

## new函數

另一個創建變量的方法是調用用內建的new函數。表達式new(T)將創建一個T類型的匿名變量，初始化爲T類型的零值，然後返迴變量地址，返迴的指針類型爲*T。

```
p := new(int) // p, *int  類型, 指向匿名的 int 變量
fmt.Println(*p) // "0"
*p = 2 // 設置 int 匿名變量的值爲 2
fmt.Println(*p) // "2"
```
new函數類似是一種語法醣，而不是一個新的基礎概念。
下面的兩個newInt函數有着相同的行爲：

```
func newInt() *int {                func newInt() *int {
    return new(int)                     var dummy int
}                                       return &dummy
                                    }
```

每次調用new函數都是返迴一個新的變量的地址，因此下面兩個地址是不同的：

```
p := new(int)
q := new(int)
fmt.Println(p == q) // "false"
```
---

## 變量的生命週期

變量的生命週期指的是在程序運行期間變量有效存在的時間間隔。
對於在包一級聲明的變量來説，它們的生命週期和整個程序的運行週期是一致的

```
for t := 0.0; t < cycles*2*math.Pi; t += res {
    x := math.Sin(t)
    y := math.Sin(t*freq + phase)
    img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
        blackIndex)
}
```

---


## 賦值

使用賦值語句可以更新一個變量的值，最簡單的賦值語句是將要被賦值的變量放在=的左邊，新值的表達式放在=的右邊。

```
x = 1                       // 命令變量的賦值
*p = true                   // 通過指針間接賦值
person.name = "bob"         // 結構體字段賦值
count[x] = count[x] * scale // 數組、slice或map的元素賦值
```

特定的二元算術運算符和賦值語句的複合操作有一個簡潔形式

```
count[x] *= scale

v := 1 
v++    // 等價方式 v = v + 1；v 變成 2 
v--    // 等價方式 v = v - 1；v 變成 1
```

---

## 元組賦值

元組賦值是另一種形式的賦值語句，它允許同時更新多個變量的值
在賦值之前，賦值語句右邊的所有表達式將會先進行求值，然後再統一更新左邊對應變量的值。

```
x, y = y, x

a[i], a[j] = a[j], a[i]
```

計算兩個整數值的的最大公約數（GCD）

```
func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x%y
    }
    return x
}
```

計算斐波納契數列（Fibonacci）的第N個數

```
func fib(n int) int {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        x, y = y, x+y
    }
    return x
}
```
元組賦值也可以使一繫列瑣碎賦值更加緊湊（譯註: 特别是在for循環的初始化部分）

```
i, j, k = 2, 3, 5
```
表達式太複雜的話，應該盡量避免過度使用元組賦值

有些表達式會産生多個值，比如調用一個有多個返迴值的函數。當這樣一個函數調用出現在元組賦值右邊的表達式中時（譯註：右邊不能再有其它表達式），左邊變量的數目必須和右邊一致。

```
f, err = os.Open("foo.txt") // function call returns two values
```
如果map査找（§4.3）、類型斷言（§7.10）或通道接收（§8.4.2）出現在賦值語句的右邊，它們都可能會産生兩個結果，有一個額外的布爾結果表示操作是否成功：

```
v, ok = m[key]             // map lookup
v, ok = x.(T)              // type assertion
v, ok = <-ch               // channel receive
```
map査找（§4.3）、類型斷言（§7.10）或通道接收（§8.4.2）出現在賦值語句的右邊時，併不一定是産生兩個結果，也可能隻産生一個結果。對於值産生一個結果的情形，map査找失敗時會返迴零值，類型斷言失敗時會發送運行時panic異常，通道接收失敗時會返迴零值（阻塞不算是失敗）

```
v = m[key]                // map査找，失敗時返迴零值
v = x.(T)                 // type斷言，失敗時panic異常
v = <-ch                  // 管道接收，失敗時返迴零值（阻塞不算是失敗）

_, ok = m[key]            // map返迴2個值
_, ok = mm[""], false     // map返迴1個值
_ = mm[""]                // map返迴1個值
```
和變量聲明一樣，我們可以用下劃線空白標識符_來丟棄不需要的值。

```
_, err = io.Copy(dst, src) // 丟棄字節數
_, ok = x.(T)              // 隻檢測類型，忽略具體值
```

---
##  可賦值性

賦值語句是顯式的賦值形式，但是程序中還有很多地方會發生隱式的賦值行爲

```
medals := []string{"gold", "silver", "bronze"}
```
隱式地對slice的每個元素進行賦值操作，類似這樣寫的行爲：

```
medals[0] = "gold" 
medals[1] = "silver" 
medals[2] = "bronze"
```

---

## 類型

1. **變量**或**表達式**的類型定義了其在內存中的存儲特徵，例如大小和表示方式。
2. 這些類型還確定了支持的操作符和相關的方法集。
3. 在程式中，相同的內部結構的變量可能表示完全不同的概念。如：
4. int：可以是迭代索引、時間戳、文件描述符、月份等。
5. float64：可以是速度或溫度。
6. 字符串：可以是密碼或顔色名稱。
7. 通過類型聲明語句，可以創建具有相同底層結構但名稱不同的新類型。
8. 這種新命名的類型允許分隔不同的概念，即使底層結構相同，它們也是不兼容的。

```
type 類型名字 底層類型
```
爲了説明類型聲明，我們將不同溫度單位分别定義爲不同的類型：

```
package tempconv

import "fmt"

type Celsius float64    // 攝氏溫度
type Fahrenheit float64 // 華氏溫度

const (
    AbsoluteZeroC Celsius = -273.15 // 絶對零度
    FreezingC     Celsius = 0       // 結冰點溫度
    BoilingC      Celsius = 100     // 沸水溫度
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
```
底層數據類型決定了內部結構和表達方式，也決定是否可以像底層類型一樣對內置運算符的支持。

這意味着，Celsius和Fahrenheit類型的算術運算行爲和底層的float64類型是一樣的，正如我們所期望的那樣。

```
fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
boilingF := CToF(BoilingC)
fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
fmt.Printf("%g\n", boilingF-FreezingC)       // compile error: type mismatch
```
比較運算符==和<也可以用來比較一個命名類型的變量和另一個有相同類型的變量，或有着相同底層類型的未命名類型的值之間做比較。
但是如果兩個值有着不同的類型，則不能直接進行比較：

```
var c Celsius
var f Fahrenheit
fmt.Println(c == 0)          // "true"
fmt.Println(f >= 0)          // "true"
fmt.Println(c == f)          // compile error: type mismatch
fmt.Println(c == Celsius(f)) // "true"!
```

註意最後那個語句。盡管看起來想函數調用，但是Celsius(f)是類型轉換操作，它併不會改變值，僅僅是改變值的類型而已。測試爲眞的原因是因爲c和g都是零值。

一個命名的類型可以提供書寫方便，特别是可以避免一遍又一遍地書寫複雜類型（譯註：例如用匿名的結構體定義變量）。雖然對於像float64這種簡單的底層類型沒有簡潔很多，但是如果是複雜的類型將會簡潔很多，特别是我們卽將討論的結構體類型。

命名類型還可以爲該類型的值定義新的行爲。這些行爲表示爲一組關聯到該類型的函數集合，我們稱爲類型的方法集。我們將在第六章中討論方法的細節，這里值説寫簡單用法。

下面的聲明語句，Celsius類型的參數c出現在了函數名的前面，表示聲明的是Celsius類型的一個叫名叫String的方法，該方法返迴該類型對象c帶着°C溫度單位的字符串：

```
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
```
