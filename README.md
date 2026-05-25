# Go Tutorial

這是我用來熟悉 Go 語法與專案開發流程的練習 repo。

## 專案結構

```
go-tutorial/
├── hello/          # 最基本的 Go 程式 — "Hello, World!"
│   ├── go.mod
│   └── hello.go
├── quote/          # 示範引用外部模組 (rsc.io/quote)
│   ├── go.mod
│   ├── go.sum
│   └── quote.go
└── README.md
```

### hello

最簡單的 Go 入門程式，無外部依賴，僅使用標準函式庫 `fmt` 輸出 `Hello, World!`。

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### quote

示範如何引用外部模組 [`rsc.io/quote`](https://pkg.go.dev/rsc.io/quote)，呼叫 `quote.Go()` 取得一段經典的 Go 語言諺語。

```go
package main

import "fmt"
import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
}
```

## 執行方式

### 執行 hello

```bash
cd hello
go run .
```

### 執行 quote

```bash
cd quote
go run .
```

## Go 版本

專案使用 **Go 1.26.3**。

## 學習目標

- [x] 建立 Go module (`go mod init`)
- [x] 撰寫基本 Go 程式 (package main, func main)
- [x] 使用 `fmt.Println` 輸出文字
- [x] 引用外部模組 (`go get`, go.mod, go.sum)
- [ ] 更多語法練習… (ing)

## 參考資源

- [Go 官方入門教學](https://go.dev/doc/tutorial/getting-started)
