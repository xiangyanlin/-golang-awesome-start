package main

import "fmt"

func main() {
	//testArray()
	//testMultiArray()
	//testSplice()
	testAppendSplice()
}

/**
 * 数组
 */
func testArray() {
	var a [3]int             // 定义三个整数的数组
	fmt.Println(a[0])        // 打印第一个元素
	fmt.Println(a[len(a)-1]) // 打印最后一个元素
	// 打印索引和元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	// 仅打印元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

}

/**
 * 多维数组
 */
func testMultiArray() {
	var array [2][2]int
	// 设置每个元素的整型值
	array[0][0] = 10
	array[0][1] = 20
	array[1][0] = 30
	array[1][1] = 40

	for _, elementArray := range array {
		for _, element := range elementArray {
			fmt.Println(element)
		}
	}
}

/**
 * 切片
 */
func testSplice() {
	var a = [3]int{1, 2, 3}
	fmt.Println(a, a[1:2])

	// 声明字符串切片
	var strList []string
	// 声明整型切片
	var numList []int
	// 声明一个空切片
	var numListEmpty = []int{}
	// 输出3个切片
	fmt.Println(strList, numList, numListEmpty)
	// 输出3个切片大小
	fmt.Println(len(strList), len(numList), len(numListEmpty))
	// 切片判定空的结果
	fmt.Println(strList == nil)
	fmt.Println(numList == nil)
	fmt.Println(numListEmpty == nil)
}

/**
 * 切片添加元素
 */
func testAppendSplice() {
	var a []int
	a = append(a, 1)                 // 追加1个元素
	a = append(a, 1, 2, 3)           // 追加多个元素, 手写解包方式
	a = append(a, []int{1, 2, 3}...) // 追加一个切片, 切片需要解包
	// 仅打印元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
}
