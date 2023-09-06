package main

import (
	"fmt"
	"os"
	"strings"
)
// 運行go run main.go arg1 arg2 arg3，那麼os.Args將是["main.go", "arg1", "arg2", "arg3"]
func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
