package main

import (
	"fmt"
	"math"
)


// 定義一個 `Vertex` struct，並為它定義一個 `Abs` 方法，計算從原點到該點的距離（即向量的長度）
type Vertex struct {
	X, Y float64
}

// 定義 `Abs` 方法，接收者是 `Vertex` 的值（非指標）
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 定義一個 `Scale` 方法，接收者是 `Vertex` 的指標，將該點的座標乘以一個比例因子
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// --- 什麼時候該用指標 receiver？ ---
//
// 1. 需要修改接收者（見上方 *Vertex.Scale）
// 2. 避免拷貝大型 struct
// 3. 保持一致：同型別 method 最好統一用指標 receiver

// 以下用 ValueReveiver / PointerReceiver 對比「值 vs 指標」的行為差異
type User struct {
	Name string
	Age  int
}

// 值 receiver — Go 會複製整個 User 再傳進去，裡面怎麼改都不影響原來的
func (u User) ValueBirthday() {
	u.Age++ // 改的是副本，外面看不到
	fmt.Printf("   (value receiver) inside: Age=%d\n", u.Age)
}

// 指標 receiver — 直接操作原來的 User，不需要複製
func (u *User) PtrBirthday() {
	u.Age++ // 改的是原始資料
	fmt.Printf("   (pointer receiver) inside: Age=%d\n", u.Age)
}

// 指標 receiver 才能真的修改接收者
type Counter struct{ n int }

func (c *Counter) Inc() { c.n++ } // ✓ 改得到原始的 c
// func (c Counter) Inc() { c.n++ } // ✗ 改的是副本，外面看不到

// 定義一個 `MyFloat` 類型，並為它定義一個 `Abs` 方法，計算絕對值
type MyFloat float64

// 定義 `Abs` 方法，接收者是 `MyFloat` 的值
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())

	Scale(&v, 10)
	fmt.Println(Abs(v))

	p := &Vertex{4, 3}
	p.Scale(3)
	Scale(p, 8)

	fmt.Println(v, p)

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// --- 什麼時候該用指標 receiver？ ---
	u := User{Name: "Alice", Age: 30}

	fmt.Println("--- 值 receiver 的行為 ---")
	fmt.Printf("  before: Age=%d\n", u.Age)
	u.ValueBirthday() // 裡面 Age++
	fmt.Printf("  after:  Age=%d (沒變！)\n", u.Age)

	fmt.Println("--- 指標 receiver 的行為 ---")
	fmt.Printf("  before: Age=%d\n", u.Age)
	u.PtrBirthday() // 裡面 Age++
	fmt.Printf("  after:  Age=%d (增加了！)\n", u.Age)

	c2 := Counter{n: 0}
	c2.Inc()
	c2.Inc()
	fmt.Println("Counter after 2 Inc() calls:", c2.n)
}
