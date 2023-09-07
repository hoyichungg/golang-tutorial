package main

import "fmt"

// 常量boilingF是在包一級范圍聲明語句聲明的
const boilingF = 212.0

func main() {
	// f和c兩個變量是在main函數內部聲明的聲明語句聲明的
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}
