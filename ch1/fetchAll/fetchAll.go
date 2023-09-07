package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()               // start := time.Now(): 紀錄程式開始時的時間。
	ch := make(chan string)           // ch := make(chan string): 建立一個 string 類型的通道，用於 goroutines 之間的通訊。
	for _, url := range os.Args[1:] { // for _, url := range os.Args[1:]: 循環遍歷命令列參數，忽略第一個參數（這是執行的程序的名稱）
		go fetch(url, ch) // start a goroutine   go fetch(url, ch): 為每一個 URL 啟動一個 goroutine，並且開始下載 URL 的內容
	}
	for range os.Args[1:] { // for range os.Args[1:]: 確保我們從通道中接收和印出與命令列參數相同數量的消息
		// ch 是一個 Go 語言中的 channel。可以將其視為一個通訊機制，它允許一個 goroutine 將數據傳遞給另一個 goroutine
		// 這是從 channel 中讀取（或接收）一個值的語法。當你使用 <-ch，你正在試圖從 ch 這個 channel 中讀取一個值。這個操作會阻塞，直到有一個 goroutine 往這個 channel 寫入一個值，或者這個 channel 被關閉。
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now() // start := time.Now(): 記錄獲取開始時的時間。
	resp, err := http.Get(url) // resp, err := http.Get(url): 發出 HTTP GET 請求以獲取 URL 的內容。 如果有錯誤，發送錯誤消息到通道，然後返回。
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body) //nbytes, err := io.Copy(io.Discard, resp.Body): 計算 HTTP 回應的大小。這裡讀取了回應的主體但沒有儲存，因為它被拷貝到 io.Discard，即被丟棄。
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
