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

- 読み込みの補助関数
  - ioutil.ReadAll
    - すべて読み込み
    ```go
    buf, err := ioutil.ReadAll(reader)
    ```
  - io.ReadFull
    - バッファの分だけ読み込めなかったらエラー
    ```go
    buffer := make([]byte, 4)
    size, err := io.ReadFull(reader, buffer)
  ```
- コピーの補助関数
  - io.Copy
    - io.Readerからio.Writerにそのままデータを渡す
    ```go
    writeSize, err := io.Copy(writer, reader)
    writeSize, err := io.CopyN(writer, reader, size)
    ```
