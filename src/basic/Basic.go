package main

import (
	"fmt"
)

func main() {
	//testFloat()
	//testBool()
	//testMultiStr()
	//testConvert()
	testPtr()
}

/**
 * 打印
 */
func testPrint() {
	println("Hello World Golang !!!")
	fmt.Println("Hello World Golang !!!")
}

// 声明全局变量
var a float32 = 3.14

/**
 * 整数
 */
func testInt() {
	var a int = 100
	var b int = 200
	b, a = a, b
	fmt.Println(a, b)
}

/**
 * 小数
 */
func testFloat() {
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)    // "true"!

}

/**
 * 负数
 */
func testComplex() {
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println(x * y)               // "(-5+10i)"
}

func testBool() {
	var n bool = true
	fmt.Println(n)
}

/**
 * 字符串拼接
 */
func testStr() {
	str := "Beginning of the string " +
		"second part of the string"
	fmt.Println(str)
}

/**
 * 多行文本
 */
func testMultiStr() {
	const str = `第一行
第二行
第三行
\r\n
`
	fmt.Println(str)
}

/**
 * 类型转换
 */
func testConvert() {
	a := 5.0
	b := int(a)
	fmt.Println(b)
}

func testPtr() {
	var cat int = 1
	var str string = "banana"
	var name = "zhangsan"
	fmt.Printf("%p %p", &cat, &str)
	fmt.Println(&name)
}
