package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1) //  遞歸案例
}

func main() {
	fmt.Println(factorial(5)) // 輸出: 120
}

//堆疊溢出 (Stack Overflow)：如果遞歸沒有正確地進行，例如基礎案例沒有正確定義，則可能會導致無窮的遞歸呼叫，最終造成堆疊溢出。
//效能問題：過深或不必要的遞歸可能會導致效能問題。在某些情況下，使用迴圈可能更高效。
//尾遞歸 (Tail Recursion)：這是遞歸的一種特殊形式，其中函數的最後一個操作是遞歸呼叫。某些編譯器和語言可以優化尾遞歸，使其與迴圈一樣高效。
