package structural

import (
	"bufio"
	"bytes"
	"fmt"
)

// 数据I/O对象及操作--------------------

//-------------Reader操作-------------------------------------------------

// Read() 法的功能是读取数据，并存放到字节切片 p 中。
func operateReaderRead() {
	data := []byte("C语言中文网")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var buf [128]byte
	n, err := r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)
}

// ReadByte()  方法的功能是读取并返回一个字节，如果没有字节可读，则返回错误信息。
func operateReaderReadByte() {
	data := []byte("Go语言入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	c, err := r.ReadByte()
	fmt.Println(string(c), err)
}

// ReadBytes() 方法的功能是读取数据直到遇到第一个分隔符“delim”，并返回读取的字节序列（包括“delim”）。
func operateReaderReadBytes() {
	data := []byte("C语言中文网, Go语言入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var delim byte = ','
	line, err := r.ReadBytes(delim)
	fmt.Println(string(line), err)
}

// ReadLine() ReadLine() 是一个低级的用于读取一行数据的方法，大多数调用者应该使用 ReadBytes('\n') 或者 ReadString('\n')。
// ReadLine 返回一行，不包括结尾的回车字符，如果一行太长（超过缓冲区长度），参数 isPrefix 会设置为 true 并且只返回前面的数据，剩余的数据会在以后的调用中返回。
func operateReaderReadLine() {
	data := []byte("Golang is a beautiful language. \r\n I like it!")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	line, prefix, err := r.ReadLine()
	fmt.Println(string(line), prefix, err)
}

// ReadRune()  方法的功能是读取一个 UTF-8 编码的字符，并返回其 Unicode 编码和字节数。
// 如果编码错误，ReadRune 只读取一个字节并返回 unicode.ReplacementChar(U+FFFD) 和长度 1。
func operateReaderReadRune() {
	data := []byte("C语言中文网")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	ch, size, err := r.ReadRune()
	fmt.Println(string(ch), size, err)
}

// ReadSlice() 方法的功能是读取数据直到分隔符“delim”处，并返回读取数据的字节切片，下次读取数据时返回的切片会失效。
// 如果 ReadSlice 在查找到“delim”之前遇到错误，它返回读取的所有数据和那个错误（通常是 io.EOF）。
func operateReaderReadSlice() {
	data := []byte("C语言中文网, Go语言入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var delim byte = ','

	line, err := r.ReadSlice(delim)
	fmt.Println(string(line), err)

	line, err = r.ReadSlice(delim)
	fmt.Println(string(line), err)

	line, err = r.ReadSlice(delim)
	fmt.Println(string(line), err)
}

// ReadString()  方法的功能是读取数据直到分隔符“delim”第一次出现，并返回一个包含“delim”的字符串。
// 如果 ReadString 在读取到“delim”前遇到错误，它返回已读字符串和那个错误（通常是 io.EOF）。
// 只有当返回的字符串不以“delim”结尾时，ReadString 才返回非空 err。
func operateReaderReadString() {
	data := []byte("C语言中文网, Go语言入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var delim byte = ','

	line, err := r.ReadString(delim)
	fmt.Println(line, err)
}

// Buffered() 方法的功能是返回可从缓冲区读出数据的字节数
func operateReaderBuffered() {
	data := []byte("Go语言入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var buf [14]byte

	n, err := r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)
	rn := r.Buffered()
	fmt.Println(rn)

	n, err = r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)
	rn = r.Buffered()
	fmt.Println(rn)
}

// Peek() Peek() 方法的功能是读取指定字节数的数据，这些被读取的数据不会从缓冲区中清除。
// 在下次读取之后，本次返回的字节切片会失效。
// 如果 Peek 返回的字节数不足 n 字节，则会同时返回一个错误说明原因，如果 n 比缓冲区要大，则错误为 ErrBufferFull。
func operateReaderPeek() {
	data := []byte("Go语言入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)

	bl, err := r.Peek(8)
	fmt.Println(string(bl), err)

	bl, err = r.Peek(14)
	fmt.Println(string(bl), err)

	bl, err = r.Peek(20)
	fmt.Println(string(bl), err)
}

//--------------- Writer ---------------------------------------------------

// Available() 方法的功能是返回缓冲区中未使用的字节数
func operateWriterAvailable() {
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	p := []byte("C语言中文网")
	fmt.Println("写入前未使用的缓冲区为：", w.Available())
	w.Write(p)
	fmt.Printf("写入%q后，未使用的缓冲区为：%d\n", string(p), w.Available())
}

// Buffered()  方法的功能是返回已写入当前缓冲区中的字节数
func operateWriterBuffered() {
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	p := []byte("C语言中文网")
	fmt.Println("写入前未使用的缓冲区为：", w.Buffered())
	w.Write(p)
	fmt.Printf("写入%q后，未使用的缓冲区为：%d\n", string(p), w.Buffered())
	w.Flush()
	fmt.Println("执行 Flush 方法后，写入的字节数为：", w.Buffered())
}
