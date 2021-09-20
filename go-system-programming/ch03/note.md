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
## 入出力に関するその他のインターフェース
- io.Coloser
  - `func Close() error`をもつ
  - ファイルを閉じる
- io.Seeker
  - `func Seek(offset int64, whence int) (int64, error)`をもつ
  - 読み書き位置を移動
- io.ReaderAt
  - `func ReadAt(p []byte, off int64) (n int, err error)`をもつ
  - 対象となるオブジェクトがランダムアクセスを行えるとき、好きな位置に自由にアクセスできる

## io.Readerの実装例
```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	for {
		buffer := make([]byte, 5)
		size, err := os.Stdin.Read(buffer)
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		fmt.Printf("size=%d input='%s'\n", size, string(buffer))
	}
}
```
```go
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	var buffer bytes.Buffer
	reader := bytes.NewBufferString("Example of io.TeeReader\n")
	teeReader := io.TeeReader(reader, &buffer)
	_, _ = ioutil.ReadAll(teeReader)
	fmt.Println(buffer.String())
}
```
