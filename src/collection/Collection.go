package main

import (
	"container/list"
	"fmt"
	"sync"
)

func main() {
	//testArray()
	//testMultiArray()
	//testSplice()
	//testAppendSplice()
	//testCopySplice()
	//testDeleteSplice()
	//testMultiSplice()
	//testMap()
	//testSyncMap()
	testList()
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

/**
 * 切片复制
 */
func testCopySplice() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	//copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	for _, v := range slice1 {
		fmt.Printf("%d ", v)
	}
	fmt.Println("----------")
	for _, v := range slice2 {
		fmt.Printf("%d ", v)
	}
}

/**
 * 切片删除
 */
func testDeleteSplice() {
	var a = []int{1, 2, 3}
	a = append(a[:1], a[1+1:]...) // 删除中间1个元素
	//a = append(a[:i], a[i+N:]...)  // 删除中间N个元素
	//a = a[:i+copy(a[i:], a[i+1:])] // 删除中间1个元素
	//a = a[:i+copy(a[i:], a[i+N:])] // 删除中间N个元素

	for _, v := range a {
		fmt.Printf("%d ", v)
	}
}

/**
 * 多维切片
 */
func testMultiSplice() {
	// 声明一个二维整型切片并赋值
	slice := [][]int{{10}, {100, 200}}
	fmt.Println(slice)
}

/**
 * map
 */
func testMap() {
	scene := make(map[string]int)
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	delete(scene, "china")
	for k, v := range scene {
		fmt.Println(k, v)
	}
}

/**
 * sync map
 */
func testSyncMap() {
	var scene sync.Map
	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	// 根据键删除对应的键值对
	scene.Delete("london")
	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}

func testList() {
	var valueList list.List
	valueList.PushBack("fist")
	valueList.PushFront(67)

	for i := valueList.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
