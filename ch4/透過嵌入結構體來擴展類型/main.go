package main

import "fmt"

// 基本結構體
type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

// 我們首先定義了一個 Person 結構體和一個 Greet 方法。

// 通過嵌入 Person 來擴展的結構體
type Student struct {
	Person
	Major string
}

// 然後，我們定義了一個 Student 結構體，它嵌入了 Person。這意味著 Student 繼承了 Person 的所有字段和方法。

func (s Student) Study() {
	fmt.Println(s.Name, "is studying", s.Major)
}

// 我們還為 Student 添加了一個新方法 Study。
// 透過這種方式，Student 既有自己的 Study 方法，也繼承了 Person 的 Greet 方法。

func main() {
	s := Student{
		Person: Person{Name: "Alice"},
		Major:  "Computer Science",
	}

	s.Greet() // 輸出: Hello, my name is Alice
	s.Study() // 輸出: Alice is studying Computer Science
}

// 需要注意的是，Go 語言中的嵌入並不等同於傳統的繼承。它更像是一種組合，但語法上提供了一種使其看起來像繼承的方式。
