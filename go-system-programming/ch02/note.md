# 低レベルアクセスのための抽象化レイヤー1: io.Writer

## io.Writerは出力を抽象化したもの
- OSにおける抽象化
  - OSでは、システムコールをファイルディスクリプタ（捜査対象のファイルを識別するための番号）に対して呼び出す
  - ファイルディスクリプタが指すものはファイルに限らず、標準出力やソケットなども含まれる
  - ファイルディスクプタはOSがカーネルのレイヤーで用意している抽象化の仕組み
- Goにおける抽象化
  - WindowsやMacOSなど、OSごとにAPIが異なってくるが、Goではio.Writerというインターフェースでそれらを抽象化している
  ```go
    type Writer interface {
      Write(p []byte) (n int, err error)
    }
  ```
## io.Writerの実装例
```go
package main

import (
	"bytes"
	"fmt"
)
```
```go
func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("byte test\n"))
	fmt.Println(buffer.String())
}
```
```go
package main

import "os"

func main() {
	file, _ := os.Create("test.txt")
	file.Write([]byte("hello from go!\n"))
	file.Close()
}
```
```go
package main

import (
	"io"
	"net"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp", "ascii.jp:80")
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n")
	io.Copy(os.Stdout, conn)
}
```
```go
package main

import "os"

func main() {
	os.Stdout.Write([]byte("write to stdout from go\n"))
}
```
