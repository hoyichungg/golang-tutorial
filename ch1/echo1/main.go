package main

import (
	"fmt"
	"os" // os這個package提供了操作繫統無關（跨平台）的
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ { // os.Args這個變量是一個字符串（string）的slice
		s += sep + os.Args[i] // string類型，+號表示字符串的連接
		sep = " "
	}
	fmt.Println(s)
}

// ### Go的for循環有很多種形式

// for initialization; condition; post {
// zero or more statements
// }

// a traditional "while" loop
// for condition {
// ...
// }

// a traditional infinite loop
// for {
// ...
// }
