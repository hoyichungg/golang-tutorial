package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // map 的键（key）是字符串，而值（value）是整数，表示该行出現的次数
	files := os.Args[1:]           // os.Args 是一个字符串切片，包含了命令行参数。  os.Args[1:] 包含所有其他的参数。
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		// 讀取文件
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// os.File 是一個結構體，它提供了許多與文件操作相關的方法（例如讀、寫、關閉等）
// * 是一個指針標識符。 f *os.File 不是一個實際的文件對象，而是一個指向文件對象的指針
// 文件通常涉及到系統資源，通過指針傳遞可以避免不必要的值複製，這在效率上是有利的
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
