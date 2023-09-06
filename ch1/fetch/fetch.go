package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// net/http和io包，http.Get函數是創建HTTP請求的函數
func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		// resp的Body字段包括一個可讀的服務器響應流
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// io.ReadAll函數從response中讀取到全部內容 其結果保存在變量b中。
		b, err := io.ReadAll(resp.Body)
		// resp.Body.Close這一句會關閉resp的Body流，防止資源洩露，Printf函數會將結果b寫出到標準輸出流中。
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
