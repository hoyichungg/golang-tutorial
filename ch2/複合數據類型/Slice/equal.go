package main

// slice和數組的字面值語法很類似，它們都是用花括弧包含一繫列的初始化元素，但是對於slice併沒有指明序列的長度。
// 這會隱式地創建一個合適大小的數組，然後slice的指針指向底層的數組。
// 就像數組字面值一樣，slice的字面值也可以按順序指定初始化值序列，或者是通過索引和元素值指定，或者的兩種風格的混合語法初始化。

func equal(x, y []string) bool { //它接收兩個參數 x 和 y，兩者都是字符串切片
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
