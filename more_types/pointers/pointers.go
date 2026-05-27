package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	// --- zero value of a pointer is nil ---
	var z *int
	fmt.Println("z == nil?", z == nil) // true

	// Safely check for nil before dereferencing:
	if z != nil {
		fmt.Println(*z)
	} else {
		fmt.Println("z is nil, cannot dereference")
	}

	// --- Go 沒有指標算術（不像 C）---
	// 在 C 語言中，你可以這樣做：
	//   int arr[] = {10, 20, 30};
	//   int *p = arr;       // 指向 arr[0]
	//   printf("%d\n", *(p + 1));  // p+1 偏移到 arr[1]，印出 20
	//
	// Go 不允許指標加減運算，以下寫法無法編譯：
	//   arr := [3]int{10, 20, 30}
	//   p := &arr[0]
	//   fmt.Println(*(p + 1))  // ✗ compilation error: invalid operation
	fmt.Println("Go: pointer arithmetic is not allowed")
	//
	// --- 使用指標交換兩個變數的值 ---
	a, b := 1, 2
	fmt.Printf("Before swap: a=%d, b=%d\n", a, b)
	swap(&a, &b) // 傳遞 a 和 b 的地址
	fmt.Printf("After swap: a=%d, b=%d\n", a, b)
}

func swap(x, y *int) {
	*x, *y = *y, *x
}
