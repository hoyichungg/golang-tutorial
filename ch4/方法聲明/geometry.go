package main

import "math"

type Point struct{ X, Y float64 }

// 這行定義了一個名為 Point 的結構體，它有兩個浮點數字段 X 和 Y，代表一個二維平面上的點。

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 這是一個獨立的函數，名為 Distance。它接受兩個 Point 作為參數，並返回兩點之間的距離（浮點數）。
// 距離的計算使用了 math.Hypot 函數，這個函數計算直角三角形的斜邊長度，也就是兩點之間的距離。

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 這是一個方法，而不是一個獨立的函數。它與 Point 結構體相關聯。
// 這個方法的名稱也是 Distance，但它是 Point 的一個方法。這意味著你可以使用一個 Point 對象來調用它，例如：p.Distance(q)。
// 這個方法的功能與上面的獨立函數相同，都是計算兩點之間的距離。
