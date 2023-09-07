package main

import "fmt"

// 一個函數的聲明由一個函數名字、參數列表（由函數的調用者提供參數變量的具體值）、一個可選的返迴值列表和包含函數定義的函數體組成。
// 如果函數沒有返迴值，那麽返迴值列表是省略的。
// 執行函數從函數的第一個語句開始，依次順序執行直到遇到return返迴語句，如果沒有返迴語句則是執行到函數末尾，然後返迴到函數調用者。

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   // "212°F = 100°C"
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
