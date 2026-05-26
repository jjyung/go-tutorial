# Go Tutorial

這是我用來熟悉 Go 語法與專案開發流程的練習 repo。

## 專案結構

```
go-tutorial/
├── hello/              # 入口程式，引用 greetings 模組
│   ├── go.mod
│   ├── go.sum
│   └── hello.go
├── greetings/          # 自製模組 — 產生隨機問候語
│   ├── go.mod
│   ├── greeting.go
│   └── greetings_test.go
├── quote/              # 示範引用外部模組 (rsc.io/quote)
│   ├── go.mod
│   ├── go.sum
│   └── quote.go
└── README.md
```

### hello

從「Hello, World!」出發，逐步擴充為引用本地自製 `greetings` 模組的入門程式。
示範 module replace、錯誤處理、slice/map 操作。

### greetings

自製的 Go module (`example.com/greetings`)，提供隨機問候語產生功能。
包含 `Hello(name)` 與 `Hellos(names)` 兩個函式，以及單元測試。

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

### 直接執行（開發用）

```bash
cd hello
go run .
```

### 編譯 + 安裝（產出 binary）

```bash
cd hello

# 1. 確認可以編譯
go build ./...

# 2. 安裝到 $GOPATH/bin/
go install

# 3. 確認 $GOPATH/bin/ 在 PATH 中（僅需設定一次）
#    將這行加入 ~/.zshrc（zsh）或 ~/.bashrc（bash）：
#       export PATH=$PATH:$(go env GOPATH)/bin
#    或當前 terminal 直接執行：
export PATH=$PATH:$(go env GOPATH)/bin

# 4. 從任何目錄執行
hello
```

> `go run .` 適合開發階段快速測試；`go install` 會產出 binary 到 `$GOPATH/bin/`，
> 把該目錄加入 PATH 後，就可以像一般指令一樣在任何地方執行。

## Go 版本

專案使用 **Go 1.26.3**。

## 學習目標

- [x] 建立 Go module (`go mod init`)
- [x] 撰寫基本 Go 程式 (package main, func main)
- [x] 使用 `fmt.Println` 輸出文字
- [x] 引用外部模組 (`go get`, go.mod, go.sum)
- [x] 自製本地 module + `replace` 指令
- [x] `go build` / `go install` / PATH 設定
- [x] 錯誤處理（回傳 error、log.Fatal）
- [x] slice、map 操作
- [x] 單元測試（`_test.go`）
- [ ] 更多語法練習… (ing)

## 注意點（你踩過的坑）

| 情境 | 正確做法 |
|------|----------|
| import 本地 module | 必須用 module path（e.g. `"example.com/greetings"`），不是目錄名 |
| `go.mod` 的 `require` | 讓 `go mod tidy` 自動產生，格式為 `module/path vX.Y.Z`（空格分隔，不用 `@`） |
| `go install` 後找不到指令 | 把 `$(go env GOPATH)/bin` 加到 `PATH` |
| `go mod tidy` 沒反應 | 檢查 import path 是否正確，否則它認不出要拉什麼模組 |

## 參考資源

- [Go 官方入門教學](https://go.dev/doc/tutorial/getting-started)
- [Go module 官方教學](https://go.dev/doc/tutorial/create-module)
