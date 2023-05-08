package main

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"
)

/**
 * 函数
 */
func main() {
	//testAnonymousFunction()
	//testClosure()
	//testVariableParam()
	//testDefer()
	//readValue("test")
	//testRecursion()
	//fmt.Println(div(1, 0))
	//testPanic()
	//testRecover()
	testTime()
}

/*
*匿名函数
 */
func testAnonymousFunction() {
	func(data int) {
		fmt.Println("hello", data)
	}(100)
}

/*
*闭包
 */
func testClosure() {
	// 创建一个累加器, 初始值为1
	accumulator := Accumulate(1)
	// 累加1并打印
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator)
	// 创建一个累加器, 初始值为1
	accumulator2 := Accumulate(10)
	// 累加1并打印
	fmt.Println(accumulator2())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator2)
}

// Accumulate 提供一个值, 每次调用函数会指定对值进行累加
func Accumulate(value int) func() int {
	// 返回一个闭包
	return func() int {
		// 累加
		value++
		// 返回一个累加值
		return value
	}
}

/*
*可变参数
 */
func testVariableParam() {
	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234
	MyPrintf(v1, v2, v3, v4)
}
func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

/*
* 延迟执行语句
 */
func testDefer() {
	fmt.Println("defer begin")
	// 将defer放入延迟调用栈
	defer fmt.Println(1)
	defer fmt.Println(2)
	// 最后一个放入, 位于栈顶, 最先调用
	defer fmt.Println(3)
	fmt.Println("defer end")
}

var (
	// 一个演示用的映射
	valueByKey = make(map[string]int)
	// 保证使用映射时的并发安全的互斥锁
	valueByKeyGuard sync.Mutex
)

// 根据键读取值
func readValue(key string) int {
	// 对共享资源加锁
	valueByKeyGuard.Lock()
	// 取值
	v := valueByKey[key]
	// 对共享资源解锁
	valueByKeyGuard.Unlock()
	// 返回值
	return v
}

func readValueUseDefer(key string) int {

	valueByKeyGuard.Lock()

	// defer后面的语句不会马上调用, 而是延迟到函数结束时调用
	defer valueByKeyGuard.Unlock()
	return valueByKey[key]
}

/*
*递归
 */
func testRecursion() {
	result := 0
	for i := 1; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
}

/*
斐波那契数列
*/
func fibonacci(n int) (res int) {
	if n <= 2 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

/*
处理运行时错误
*/

// 定义除数为0的错误
var errDivisionByZero = errors.New("division by zero")

func div(dividend, divisor int) (int, error) {
	// 判断除数为0的情况并返回
	if divisor == 0 {
		return 0, errDivisionByZero
	}
	// 正常计算，返回空错误
	return dividend / divisor, nil
}

func testHandlerError() {
}

/*
*
宕机
*/
func testPanic() {
	defer fmt.Println("宕机后要做的事情1")
	defer fmt.Println("宕机后要做的事情2")
	panic("宕机")
}

/*
*
宕机恢复
*/

// 崩溃时需要传递的上下文信息
type panicContext struct {
	function string // 所在函数
}

// ProtectRun 保护方式允许一个函数
func ProtectRun(entry func()) {
	// 延迟处理的函数
	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		err := recover()
		switch err.(type) {
		case runtime.Error: // 运行时错误
			fmt.Println("runtime error:", err)
		default: // 非运行时错误
			fmt.Println("error:", err)
		}
	}()
	entry()
}
func testRecover() {
	fmt.Println("运行前")
	// 允许一段手动触发的错误
	ProtectRun(func() {
		fmt.Println("手动宕机前")
		// 使用panic传递上下文
		panic(&panicContext{
			"手动触发panic",
		})
		fmt.Println("手动宕机后")
	})
	// 故意造成空指针访问错误
	ProtectRun(func() {
		fmt.Println("赋值宕机前")
		var a *int
		*a = 1
		fmt.Println("赋值宕机后")
	})
	fmt.Println("运行后")
}

/*
*
计算函数执行时间
*/
func testTime() {
	start := time.Now() // 获取当前时间
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum++
	}
	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
	elapsed1 := time.Now().Sub(start)
	fmt.Println("该函数执行完成耗时：", elapsed1)
}

func GetArea(weight int, height int) int {
	return weight * height
}
