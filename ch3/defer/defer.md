## Defer

在 Go 語言中，defer 關鍵字用於確保一個函數呼叫在程式執行結束後、返回之前被執行。 通常，defer 用於處理資源的清理操作，例如關閉檔案、解鎖互斥鎖、關閉資料庫連線等。

以下是關於 Go 中 defer 的一些關鍵點和範例：

1. **基本用法**：你可以使用 defer 關鍵字 followed by a function call 來延遲該函數的執行。

```
package main

import "fmt"

func main() {
    defer fmt.Println("World")
    fmt.Println("Hello")
}


// Hello
// World
```

2. **執行順序**：如果有多個 defer 語句，它們會以 LIFO (後進先出) 的順序執行，即最後一個 defer 語句最先執行。

```
package main

import "fmt"

func main() {
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")
}


// 3
// 2
// 1
```
3. **常見用途**：defer 經常用於資源清理，例如關閉檔案。

```
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("filename.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    // 處理文件內容...
}
```

在上述範例中，無論函數如何返回（正常結束或因錯誤而提前返回），file.Close() 都會被調用，確保檔案資源被正確釋放。

4. **參數評估**：defer 語句中的函數參數在 defer 語句所在的函數執行時即被評估，而不是在延遲函數執行時。

```
package main

import "fmt"

func main() {
    i := 1
    defer fmt.Println(i)
    i++
    fmt.Println(i)
}

// 2
// 1
```