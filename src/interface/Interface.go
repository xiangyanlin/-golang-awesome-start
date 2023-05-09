package _interface

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
)

// DataWriter 定义一个数据写入器
type DataWriter interface {
	WriteData(data interface{}) error
}

// 定义文件结构，用于实现DataWriter
type file struct {
}

// WriteData 实现DataWriter接口的WriteData方法
func (d *file) WriteData(data interface{}) error {
	// 模拟写入数据
	fmt.Println("WriteData:", data)
	return nil
}

func writeData() {
	// 实例化file
	f := new(file)
	// 声明一个DataWriter的接口
	var writer DataWriter
	// 将接口赋值f，也就是*file类型
	writer = f
	// 使用DataWriter接口进行数据写入
	writer.WriteData("data")
}

// MyStringList 将[]string定义为MyStringList类型
type MyStringList []string

// Len 实现sort.Interface接口的获取元素数量方法
func (m MyStringList) Len() int {
	return len(m)
}

// Less 实现sort.Interface接口的比较元素方法
func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

// Swap 实现sort.Interface接口的交换元素方法
func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func mySort() {

	// 准备一个内容被打乱顺序的字符串切片
	names := MyStringList{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	// 使用sort包进行排序
	sort.Sort(names)
	// 遍历打印结果
	for _, v := range names {
		fmt.Printf("%s\n", v)
	}

}

// 声明一个设备结构
type device struct {
}

// 实现io.Writer的Write()方法
func (d *device) Write(p []byte) (n int, err error) {
	return 0, nil
}

// Close 实现io.Closer的Close()方法
func (d *device) Close() error {
	return nil
}

func operateDevice() {
	// 声明写入关闭器, 并赋予device的实例
	var wc io.WriteCloser = new(device)
	// 写入数据
	wc.Write(nil)
	// 关闭设备
	wc.Close()
	// 声明写入器, 并赋予device的新实例
	var writeOnly io.Writer = new(device)
	// 写入数据
	writeOnly.Write(nil)
}

func convert() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)      // 成功: f == os.Stdout
	c := w.(*bytes.Buffer) // 死机：接口保存*os.file，而不是*bytes.buffer
	fmt.Println(f, c)
}

// Flyer 定义飞行动物接口
type Flyer interface {
	Fly()
}

// Walker 定义行走动物接口
type Walker interface {
	Walk()
}

// 定义鸟类
type bird struct {
}

// Fly 实现飞行动物接口
func (b *bird) Fly() {
	fmt.Println("bird: fly")
}

// Walk 为鸟添加Walk()方法, 实现行走动物接口
func (b *bird) Walk() {
	fmt.Println("bird: walk")
}

// 定义猪
type pig struct {
}

// Walk 为猪添加Walk()方法, 实现行走动物接口
func (p *pig) Walk() {
	fmt.Println("pig: walk")
}

func animalAction() {
	// 创建动物的名字到实例的映射
	animals := map[string]interface{}{
		"bird": new(bird),
		"pig":  new(pig),
	}
	// 遍历映射
	for name, obj := range animals {
		// 判断对象是否为飞行动物
		f, isFlyer := obj.(Flyer)
		// 判断对象是否为行走动物
		w, isWalker := obj.(Walker)
		fmt.Printf("name: %s isFlyer: %v isWalker: %v\n", name, isFlyer, isWalker)
		// 如果是飞行动物则调用飞行动物接口
		if isFlyer {
			f.Fly()
		}
		// 如果是行走动物则调用行走动物接口
		if isWalker {
			w.Walk()
		}
	}
}

func emptyInterface() {
	var any interface{}
	any = 1
	fmt.Println(any)
	any = "hello"
	fmt.Println(any)
	any = false
	fmt.Println(any)
}

func typeSwitchDemo() {
	printType(1024)
	printType("pig")
	printType(true)
}
func printType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println(v, "is int")
	case string:
		fmt.Println(v, "is string")
	case bool:
		fmt.Println(v, "is bool")
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return -1, errors.New("math: square root of negative number")
	}
	return math.Sqrt(f), nil
}
func errorDemo() {

	result, err := Sqrt(-13)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
