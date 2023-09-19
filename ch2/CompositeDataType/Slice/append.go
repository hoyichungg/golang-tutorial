package main

import (
	"fmt"
)

func appendInt(x []int, y int) []int {
	// var z []int: 宣告一個新的整數切片 z
	var z []int
	// zlen := len(x) + 1: 計算 x 切片加上新元素後的長度。
	zlen := len(x) + 1
	if zlen <= cap(x) {
		//if zlen <= cap(x) {...}:
		// 檢查新的長度是否小於或等於 x 的容量（cap 函數可以獲取切片的容量）。
		// 如果滿足，意味著原切片 x 仍有空間容納新的元素。所以，我們可以直接擴展 x 來創建新的切片 z 而不需要重新分配記憶體。
		z = x[:zlen]
	} else {
		//else {...}:
		// 如果原切片 x 的容量不足以容納新元素，我們需要創建一個新的具有更大容量的切片。
		// 切片的容量會翻倍，這是為了確保操作的均攤線性時間複雜性。
		// 使用 make 函數創建新切片並指定所需的長度和容量。
		// 之後使用 copy 函數將原切片 x 的內容複製到新切片 z 中。
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	// z[len(x)] = y: 將新的整數值 y 添加到新切片 z 的末尾。
	z[len(x)] = y
	// return z: 返回新的切片 z。
	return z
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

// 0  cap=1    [0]
// 1  cap=2    [0 1]
// 2  cap=4    [0 1 2]
// 3  cap=4    [0 1 2 3]
// 4  cap=8    [0 1 2 3 4]
// 5  cap=8    [0 1 2 3 4 5]
// 6  cap=8    [0 1 2 3 4 5 6]
// 7  cap=8    [0 1 2 3 4 5 6 7]
// 8  cap=16   [0 1 2 3 4 5 6 7 8]
// 9  cap=16   [0 1 2 3 4 5 6 7 8 9]
