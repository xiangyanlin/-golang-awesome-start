package main

import "fmt"

/**
 * 流程控制
 */
func main() {
	//testIf(12)
	//testLoop()
	//testLoopRange()
	//testSwitch("world")
	//testGoto()
	//testBreak()
	testContinue()
}

/**
 * If
 */
func testIf(val int32) {
	if val == 1 {
		// do something
		fmt.Println("process one")
	} else if val == 2 {
		// do something else
		fmt.Println("process two")
	} else {
		// catch-all or default
		fmt.Println("process other")
	}
}

/**
 * 循环
 */
func testLoop() {
	// 遍历, 决定处理第几行
	for y := 1; y <= 9; y++ {
		// 遍历, 决定这一行有多少列
		for x := 1; x <= y; x++ {
			fmt.Printf("%d*%d=%d ", x, y, x*y)
		}
		// 手动生成回车
		fmt.Println()
	}
}

/**
 *数组、切片、字符串返回索引和值。
 *map 返回键和值。
 *通道（channel）只返回通道内的值。
 */
func testLoopRange() {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}

func testSwitch(word string) {
	switch word {
	case "hello":
		fmt.Println(1)
	case "world":
		fmt.Println(2)
	default:
		fmt.Println(0)
	}
}

func testGoto() {

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				// 跳转到标签
				goto breakHere
			}
		}
	}
	// 手动返回, 避免执行进入标签
	return
	// 标签
breakHere:
	fmt.Println("done")
}

func testBreak() {
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				break OuterLoop
			case 3:
				fmt.Println(i, j)
				break OuterLoop
			}
		}
	}
}

func testContinue() {
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				continue OuterLoop
			}
		}
	}
}
