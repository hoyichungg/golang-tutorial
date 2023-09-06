package main

import (
	"bufio"
	"fmt"
	"os"
)

// ### NOTE:
// Maps:
// counts := make(map[string]int)

// Slices:
// mySlice := make([]int, 5)      // 創建一個長度為5的int類型的slice，所有元素初始化為0
// anotherSlice := make([]int, 5, 10) // 創建一個長度為5、容量為10的int類型的slice

//Channels
// ch := make(chan int) // 創建了一個新的 int 類型的未緩衝的 channel

func main() {
	counts := make(map[string]int) // map是Go語言內置的key/value型數據結構，這個數據結構能夠提供常數時間的存儲、獲取、測試操作
	//scanner對象可以從程序的標準輸入中讀取內容。
	//對input.Scanner的每一次調用都會調入一個新行，併且會自動將其行末的換行符去掉；其結果可以用input.Text()得到。
	//Scan方法在讀到了新行的時候會返迴true，而在沒有新行被讀入時，會返迴false。
	input := bufio.NewScanner(os.Stdin) // 使用bufio.NewScanner創建了一個新的scanner，它用於讀取從os.Stdin
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
