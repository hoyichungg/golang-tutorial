JSON是對JavaScript中各種類型的值——字符串、數字、布爾值和對象——Unicode本文編碼。它可以用有效可讀的方式表示第三章的基礎數據類型和本章的數組、slice、結構體和map等聚合數據類型。

基本的JSON類型有數字（十進製或科學記數法）、布爾值（true或false）、字符串，其中字符串是以雙引號包含的Unicode字符序列，支持和Go語言類似的反斜槓轉義特性，不過JSON使用的是\Uhhhh轉義數字來表示一個UTF-16編碼（譯註：UTF-16和UTF-8一樣是一種變長的編碼，有些Unicode碼點較大的字符需要用4個字節表示；而且UTF-16還有大端和小端的問題），而不是Go語言的rune類型。

這些基礎類型可以通過JSON的數組和對象類型進行遞歸組合。一個JSON數組是一個有序的值序列，寫在一個方括號中併以逗號分隔；一個JSON數組可以用於編碼Go語言的數組和slice。一個JSON對象是一個字符串到值的映射，寫成以繫列的name:value對形式，用花括號包含併以逗號分隔；JSON的對象類型可以用於編碼Go語言的map類型（key類型是字符串）和結構體。例如：

```
boolean         true
number          -273.15
string          "She said \"Hello, BF\""
array           ["gold", "silver", "bronze"]
object          {"year": 1980,
                 "event": "archery",
                 "medals": ["gold", "silver", "bronze"]}
```


