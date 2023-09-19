package main

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}


a := [...]int{0, 1, 2, 3, 4, 5}
reverse(a[:])
fmt.Println(a) // "[5 4 3 2 1 0]"

//一種將slice元素循環向左鏇轉n個元素的方法是三次調用reverse反轉函數，第一次是反轉開頭的n個元素，然後是反轉剩下的元素，最後是反轉整個slice的元素。
//（如果是向右循環鏇轉，則將第三個函數調用移到第一個調用位置就可以了。）

s := []int{0, 1, 2, 3, 4, 5}
// Rotate s left by two positions.
reverse(s[:2])
reverse(s[2:])
reverse(s)
fmt.Println(s) // "[2 3 4 5 0 1]"

// reverse(s[:2]):

// s[:2]：這個操作從字符串 s 中取前兩個字符。例如，如果 s = "abcdef", 則 s[:2] 會返回 "ab"。

// reverse：這個函數將反轉它所接收的字符串或切片。所以，如果 s = "abcdef"，那麼 reverse(s[:2]) 會返回 "ba"。

// reverse(s[2:]):

// s[2:]：這個操作從字符串 s 中取第三個字符到最後。例如，如果 s = "abcdef", 則 s[2:] 會返回 "cdef".

// reverse：同樣地，這個函數會反轉它所接收的字符串或切片。所以，如果 s = "abcdef"，那麼 reverse(s[2:]) 會返回 "fedc"。