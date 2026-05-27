# Go Tour — 課後練習清單

> 📖 **Tour 連結**：https://go.dev/tour/welcome/1
>
> 從 A Tour of Go 挑出你還沒練過的主題，依序攻克。
> 每個主題完成後打勾 ✅，有卡住的地方寫筆記。

---

## 1. Pointers（指標）

- [x] 理解 `&`（取址）與 `*`（取值）
- [x] 知道指標的 zero value 是 `nil`
- [x] 知道 Go 沒有指標算術（不像 C）
- [x] **練習**：寫一個函式 `swap(a, b *int)` 交換兩個數值

[pointers](./more_types/pointers/pointers.go)

---

## 2. Structs（結構體）

- [x] 定義 struct、存取欄位
- [x] struct literal 的寫法
- [x] 知道 struct 是值型別（assign 會複製）
- [x] **練習**：定義一個 `Person` struct，寫一個函式印出 `"我叫 OO，今年 XX 歲"`

[structs](./more_types/structs/structs.go)

---

## 3. Methods（方法）

- [ ] 對自訂型別定義 method（不只在 struct）
- [ ] 值接收者 vs 指標接收者的差別
- [ ] 知道指標 receiver 可以修改接收者的值
- [ ] **思考**：什麼時候該用指標 receiver？

```

```

---

## 4. Interfaces（介面）

- [ ] 定義 interface
- [ ] 理解「隱式實作」（不用寫 `implements`）
- [ ] 知道 `interface{}` 是任意型別（Go 1.18 後可用 `any`）
- [ ] **練習**：定義一個 `Speaker` interface（有 `Say() string`），讓 `Person` 和 `Dog` 都實作它

```

```

---

## 5. Type Assertions & Type Switches

- [ ] `value.(Type)` 語法
- [ ] `value.(type)` 在 switch 中的用法
- [ ] **練習**：寫一個函式接收 `any`，根據型別印出不同訊息

```

```

---

## 6. Stringer Interface

- [ ] 實作 `String() string` 來自訂 struct 的列印格式
- [ ] **練習**：幫 `Person` 加上 `String()`，讓 `fmt.Println(p)` 印出有意義的內容

```

```

---

## 7. Defer

- [ ] 理解 `defer` 的執行時機（函式結束前，LIFO 順序）
- [ ] 知道 deferred 函式的參數在宣告時就求值了
- [ ] **練習**：寫一個函式，用 defer 印出 "函式結束"，觀察執行順序

```

```

---

## 8. Panic & Recover

- [ ] 知道 `panic` 會停止正常流程
- [ ] 知道 `recover` 只能在 deferred 函式中生效
- [ ] **練習**：寫一個防呆函式，呼叫某個會 panic 的函式並用 recover 接住

```

```

---

## 9. Goroutines（併發基礎）

- [ ] `go f()` 啟動一個 goroutine
- [ ] 理解 goroutine 是非同步的，main 函式結束就全部結束
- [ ] **練習**：用 `go` 啟動兩個函式，觀察輸出順序

```

```

---

## 10. Channels

- [ ] 建立 channel：`make(chan int)`
- [ ] 發送 `ch <- v`、接收 `<-ch`
- [ ] 知道 unbuffered channel 會阻塞直到另一方準備好
- [ ] buffered channel：`make(chan int, 2)`
- [ ] `for v := range ch` 搭配 `close(ch)`
- [ ] **練習**：寫一個 producer-consumer 模式，producer 發送 5 個數，consumer 印出

```

```

---

## 11. Select

- [ ] `select` 可以同時等多個 channel
- [ ] `default` 分支做非阻塞操作
- [ ] **練習**：兩個 goroutine 各發送資料，用 select 接收先到的那個

```

```

---

## 12. sync.Mutex（互斥鎖）

- [ ] 知道什麼時候需要鎖（多個 goroutine 讀寫共享變數）
- [ ] `mu.Lock()` / `mu.Unlock()`
- [ ] **練習**：用 goroutine 同時對一個 counter +1 一百次，用 Mutex 保護正確性

```

```

---

## Bonus（選修）

- [ ] **iota** — 列舉常數
- [ ] **embedding** — Go 的組合式繼承
- [ ] **errors.Is / errors.As** — Go 1.13 的錯誤包裝
- [ ] **go generate** — 自動程式碼產生
- [ ] **testing / bench** — 基準測試（benchmark）

---

## 隨手筆記

```
在此記錄卡關的地方或學到的心得：
```

