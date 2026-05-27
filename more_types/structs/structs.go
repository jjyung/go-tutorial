package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	// 定義 struct
	v := Vertex{1, 2}
	fmt.Println(v)
	// 存取欄位
	fmt.Println(v.X)
	// 修改欄位
	v.X = 4
	fmt.Println(v)
	
	// struct literal 的寫法
	v2 := Vertex{X: 3} // Y 欄位會被設為零值 0
	fmt.Println(v2)
	// struct literal 的另一種寫法，使用 & 會得到 struct 的指標
	v3 := &Vertex{X: 5, Y: 6}
	fmt.Println(v3)
	fmt.Println(v3.X) // Go 會自動 dereference 指標，所以可以直接存取欄位

	// 知道 struct 是值型別（assign 會複製）
	v4 := v
	v4.X = 10
	fmt.Println(v)  // 原本的 v 不受影響
	fmt.Println(v4) // v4 是 v 的複製品

	// 定義一個 `Person` struct，寫一個函式印出 `"我叫 OO，今年 XX 歲"`
	p := Person{Name: "Alice", Age: 30}
	printPerson(p)
}

// 定義一個 `Person` struct，寫一個函式印出 `"我叫 OO，今年 XX 歲"`
type Person struct {
	Name string
	Age  int
}

func printPerson(p Person) {
	fmt.Printf("我叫 %s，今年 %d 歲\n", p.Name, p.Age)
}
