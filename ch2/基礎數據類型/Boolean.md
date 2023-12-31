## 布爾型

一個布爾類型的值隻有兩種：true和false。
if和for語句的條件部分都是布爾類型的值，就像用x來表示x==true。
布爾值可以和&&（AND）和||（OR）操作符結合，併且可能會有短路行爲

```
s != "" && s[0] == 'x'
```

其中s[0]操作如果應用於空字符串將會導致panic異常。


因爲&&的優先級比||高（助記：&&對應邏輯乘法，||對應邏輯加法，乘法比加法優先級要高），下面形式的布爾表達式是不需要加小括弧的：

```
if 'a' <= c && c <= 'z' ||
    'A' <= c && c <= 'Z' ||
    '0' <= c && c <= '9' {
    // ...ASCII letter or digit...
}
```

布爾值併不會隱式轉換爲數字值0或1，反之亦然。必鬚使用一個顯式的if語句輔助轉換：

```
i := 0
if b {
    i = 1
    }
```

如果需要經常做類似的轉換, 包裝成一個函數會更方便:

```
func btoi(b bool) int {
    if b {
        return 1
    }
    return 0
}
```

數字到布爾型的逆轉換則非常簡單, 不過爲了保持對稱, 我們也可以包裝一個函數:

```
// itob reports whether i is non-zero.
func itob(i int) bool { return i != 0 }
```
[Slice]https://wizardforcel.gitbooks.io/gopl-zh/content/images/ch4-01.png'
