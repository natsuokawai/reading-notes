# 低レベルアクセスのための抽象化レイヤー2: io.Reader

## io.Readerは入力の抽象化
- io.Readerインターフェース
  ```go
  type Reader interface {
    func Read(p []byte) (n int, err error)
  }
  ```
## io.Readerの実装例   
```go
buffer := make([]byte, 1024)
size, err := r.Read(buffer)
```

読み込み用のバッファを用意→読み込みという処理を毎回書かなくても済むような関数も用意されている
- ioutil.ReadAll
  - すべて読み込み
  ```go
  buf, err := ioutil.ReadAll(reader)
  ```
