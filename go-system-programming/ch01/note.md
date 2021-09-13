# chapter1

## Goのデバッガdelveを使う

### 前提
- dlv（CLIツール）のインストール
  - `go install github.com/go-delve/delve/cmd/dlv@latest`
- delve-vimのインストール（vimの場合）
  - `Plug 'sebdah/vim-delve', { 'for': ['go'] }'`
    - https://github.com/sebdah/vim-delve

### 手順
vimを開き以下のコードをmain.goとして保存
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
```

上記プログラムのfmt.Printlnの行にカーソルがある状態で`:DlvAddBreakpoint`を実行してブレークポイントを設置
`DlvDebug`でデバッガを起動

以下、実行ログ
（`s`で現在の行の関数呼び出しにステップイン、nextで次の行にフォーカスを移す）

```go
Type 'help' for list of commands.
Breakpoint 1 set at 0x10aba34 for main.main() ./main.go:6
> main.main() ./main.go:6 (hits goroutine(1):1 total:1) (PC: 0x10aba34)
     1: package main
     2:
     3: import "fmt"
     4:
     5: func main() {
=>   6:         fmt.Println("Hello, world!")
     7: }
(dlv) s
> fmt.Println() /usr/local/go/src/fmt/print.go:273 (PC: 0x10a63ca)
   268: }
   269:
   270: // Println formats using the default formats for its operands and writes to standard output.
   271: // Spaces are always added between operands and a newline is appended.
   272: // It returns the number of bytes written and any write error encountered.
=> 273: func Println(a ...interface{}) (n int, err error) {
   274:         return Fprintln(os.Stdout, a...)
   275: }
   276:
   277: // Sprintln formats using the default formats for its operands and returns the resulting string.
   278: // Spaces are always added between operands and a newline is appended.
(dlv) s
> fmt.Println() /usr/local/go/src/fmt/print.go:274 (PC: 0x10a63ff)
   269:
   270: // Println formats using the default formats for its operands and writes to standard output.
   271: // Spaces are always added between operands and a newline is appended.
   272: // It returns the number of bytes written and any write error encountered.
   273: func Println(a ...interface{}) (n int, err error) {
=> 274:         return Fprintln(os.Stdout, a...)
   275: }
   276:
   277: // Sprintln formats using the default formats for its operands and returns the resulting string.
   278: // Spaces are always added between operands and a newline is appended.
   279: func Sprintln(a ...interface{}) string {
(dlv) s
> fmt.Fprintln() /usr/local/go/src/fmt/print.go:262 (PC: 0x10a626a)
   257: // after the last operand.
   258:
   259: // Fprintln formats using the default formats for its operands and writes to w.
   260: // Spaces are always added between operands and a newline is appended.
   261: // It returns the number of bytes written and any write error encountered.
=> 262: func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
   263:         p := newPrinter()
   264:         p.doPrintln(a)
   265:         n, err = w.Write(p.buf)
   266:         p.free()
   267:         return
(dlv) next
> fmt.Fprintln() /usr/local/go/src/fmt/print.go:263 (PC: 0x10a62af)
   258:
   259: // Fprintln formats using the default formats for its operands and writes to w.
   260: // Spaces are always added between operands and a newline is appended.
   261: // It returns the number of bytes written and any write error encountered.
   262: func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
=> 263:         p := newPrinter()
   264:         p.doPrintln(a)
   265:         n, err = w.Write(p.buf)
   266:         p.free()
   267:         return
   268: }
(dlv) next
> fmt.Fprintln() /usr/local/go/src/fmt/print.go:264 (PC: 0x10a62b9)
   259: // Fprintln formats using the default formats for its operands and writes to w.
   260: // Spaces are always added between operands and a newline is appended.
   261: // It returns the number of bytes written and any write error encountered.
   262: func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
   263:         p := newPrinter()
=> 264:         p.doPrintln(a)
   265:         n, err = w.Write(p.buf)
   266:         p.free()
   267:         return
   268: }
   269:
(dlv) next
> fmt.Fprintln() /usr/local/go/src/fmt/print.go:265 (PC: 0x10a62d6)
   260: // Spaces are always added between operands and a newline is appended.
   261: // It returns the number of bytes written and any write error encountered.
   262: func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
   263:         p := newPrinter()
   264:         p.doPrintln(a)
=> 265:         n, err = w.Write(p.buf)
   266:         p.free()
   267:         return
   268: }
   269:
   270: // Println formats using the default formats for its operands and writes to standard output.
(dlv) s
> os.(*File).Write() /usr/local/go/src/os/file.go:172 (PC: 0x109fc2a)
   167: }
   168:
   169: // Write writes len(b) bytes to the File.
   170: // It returns the number of bytes written and an error, if any.
   171: // Write returns a non-nil error when n != len(b).
=> 172: func (f *File) Write(b []byte) (n int, err error) {
   173:         if err := f.checkValid("write"); err != nil {
   174:                 return 0, err
   175:         }
   176:         n, e := f.write(b)
   177:         if n < 0 {
(dlv) next
> os.(*File).Write() /usr/local/go/src/os/file.go:173 (PC: 0x109fc67)
   168:
   169: // Write writes len(b) bytes to the File.
   170: // It returns the number of bytes written and an error, if any.
   171: // Write returns a non-nil error when n != len(b).
   172: func (f *File) Write(b []byte) (n int, err error) {
=> 173:         if err := f.checkValid("write"); err != nil {
   174:                 return 0, err
   175:         }
   176:         n, e := f.write(b)
   177:         if n < 0 {
   178:                 n = 0
(dlv) next
> os.(*File).Write() /usr/local/go/src/os/file.go:176 (PC: 0x109fcc2)
   171: // Write returns a non-nil error when n != len(b).
   172: func (f *File) Write(b []byte) (n int, err error) {
   173:         if err := f.checkValid("write"); err != nil {
   174:                 return 0, err
   175:         }
=> 176:         n, e := f.write(b)
   177:         if n < 0 {
   178:                 n = 0
   179:         }
   180:         if n != len(b) {
   181:                 err = io.ErrShortWrite
(dlv) s
> os.(*File).write() /usr/local/go/src/os/file_posix.go:48 (PC: 0x10a004a)
    43:         return n, err
    44: }
    45:
    46: // write writes len(b) bytes to the File.
    47: // It returns the number of bytes written and an error, if any.
=>  48: func (f *File) write(b []byte) (n int, err error) {
    49:         n, err = f.pfd.Write(b)
    50:         runtime.KeepAlive(f)
    51:         return n, err
    52: }
    53:
(dlv) next
> os.(*File).write() /usr/local/go/src/os/file_posix.go:49 (PC: 0x10a0084)
    44: }
    45:
    46: // write writes len(b) bytes to the File.
    47: // It returns the number of bytes written and an error, if any.
    48: func (f *File) write(b []byte) (n int, err error) {
=>  49:         n, err = f.pfd.Write(b)
    50:         runtime.KeepAlive(f)
    51:         return n, err
    52: }
    53:
    54: // pwrite writes len(b) bytes to the File starting at byte offset off.
(dlv) s
> internal/poll.(*FD).Write() /usr/local/go/src/internal/poll/fd_unix.go:261 (PC: 0x109edb2)
   256:                 return n, oobn, sysflags, sa, err
   257:         }
   258: }
   259:
   260: // Write implements io.Writer.
=> 261: func (fd *FD) Write(p []byte) (int, error) {
   262:         if err := fd.writeLock(); err != nil {
   263:                 return 0, err
   264:         }
   265:         defer fd.writeUnlock()
   266:         if err := fd.pd.prepareWrite(fd.isFile); err != nil {
(dlv) next
> internal/poll.(*FD).Write() /usr/local/go/src/internal/poll/fd_unix.go:262 (PC: 0x109edfb)
   257:         }
   258: }
   259:
   260: // Write implements io.Writer.
   261: func (fd *FD) Write(p []byte) (int, error) {
=> 262:         if err := fd.writeLock(); err != nil {
   263:                 return 0, err
   264:         }
   265:         defer fd.writeUnlock()
   266:         if err := fd.pd.prepareWrite(fd.isFile); err != nil {
   267:                 return 0, err
(dlv) next
> internal/poll.(*FD).Write() /usr/local/go/src/internal/poll/fd_unix.go:265 (PC: 0x109ee73)
   260: // Write implements io.Writer.
   261: func (fd *FD) Write(p []byte) (int, error) {
   262:         if err := fd.writeLock(); err != nil {
   263:                 return 0, err
   264:         }
=> 265:         defer fd.writeUnlock()
   266:         if err := fd.pd.prepareWrite(fd.isFile); err != nil {
   267:                 return 0, err
   268:         }
   269:         var nn int
   270:         for {
(dlv) next
> internal/poll.(*FD).Write() /usr/local/go/src/internal/poll/fd_unix.go:266 (PC: 0x109eee2)
   261: func (fd *FD) Write(p []byte) (int, error) {
   262:         if err := fd.writeLock(); err != nil {
   263:                 return 0, err
   264:         }
   265:         defer fd.writeUnlock()
=> 266:         if err := fd.pd.prepareWrite(fd.isFile); err != nil {
   267:                 return 0, err
   268:         }
   269:         var nn int
   270:         for {
   271:                 max := len(p)
(dlv) next
> internal/poll.(*FD).Write() /usr/local/go/src/internal/poll/fd_unix.go:269 (PC: 0x109ef92)
   264:         }
   265:         defer fd.writeUnlock()
   266:         if err := fd.pd.prepareWrite(fd.isFile); err != nil {
   267:                 return 0, err
   268:         }
=> 269:         var nn int
   270:         for {
   271:                 max := len(p)
   272:                 if fd.IsStream && max-nn > maxRW {
   273:                         max = nn + maxRW
   274:                 }
(dlv) next
> internal/poll.(*FD).Write() /usr/local/go/src/internal/poll/fd_unix.go:270 (PC: 0x109ef9b)
   265:         defer fd.writeUnlock()
   266:         if err := fd.pd.prepareWrite(fd.isFile); err != nil {
   267:                 return 0, err
   268:         }
   269:         var nn int
=> 270:         for {
   271:                 max := len(p)
   272:                 if fd.IsStream && max-nn > maxRW {
   273:                         max = nn + maxRW
   274:                 }
   275:                 n, err := ignoringEINTRIO(syscall.Write, fd.Sysfd, p[nn:max])
(dlv) next
> internal/poll.(*FD).Write() /usr/local/go/src/internal/poll/fd_unix.go:271 (PC: 0x109efca)
   266:         if err := fd.pd.prepareWrite(fd.isFile); err != nil {
   267:                 return 0, err
   268:         }
   269:         var nn int
   270:         for {
=> 271:                 max := len(p)
   272:                 if fd.IsStream && max-nn > maxRW {
   273:                         max = nn + maxRW
   274:                 }
   275:                 n, err := ignoringEINTRIO(syscall.Write, fd.Sysfd, p[nn:max])
   276:                 if n > 0 {
(dlv) next
> internal/poll.(*FD).Write() /usr/local/go/src/internal/poll/fd_unix.go:272 (PC: 0x109efd9)
   267:                 return 0, err
   268:         }
   269:         var nn int
   270:         for {
   271:                 max := len(p)
=> 272:                 if fd.IsStream && max-nn > maxRW {
   273:                         max = nn + maxRW
   274:                 }
   275:                 n, err := ignoringEINTRIO(syscall.Write, fd.Sysfd, p[nn:max])
   276:                 if n > 0 {
   277:                         nn += n
(dlv) next
> internal/poll.(*FD).Write() /usr/local/go/src/internal/poll/fd_unix.go:275 (PC: 0x109f02b)
   270:         for {
   271:                 max := len(p)
   272:                 if fd.IsStream && max-nn > maxRW {
   273:                         max = nn + maxRW
   274:                 }
=> 275:                 n, err := ignoringEINTRIO(syscall.Write, fd.Sysfd, p[nn:max])
   276:                 if n > 0 {
   277:                         nn += n
   278:                 }
   279:                 if nn == len(p) {
   280:                         return nn, err
(dlv) exit
```
