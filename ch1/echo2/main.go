package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	//os.Args: 是一個字符串切片，其中包含命令行傳遞給程序的參數。
	//os.Args[0]是程序的名稱，os.Args[1:]是除程序名之外的所有參數。
	//range: Go的for ... range循環可以用於遍歷各種集合，如數組、切片、字符串或映射。當用於遍歷切片或數組時，它返回兩個值：索引和索引處的值。
	//_: 在Go中，下劃線_是一個特殊的標識符，稱為"空白標識符"。當你不需要使用某個值時，可以用_來忽略它。代碼中，我們並不關心os.Args的索引，只關心值，因此使用_來忽略索引。
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
