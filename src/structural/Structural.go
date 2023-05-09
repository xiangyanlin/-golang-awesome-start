package structural

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

type Color struct {
	R, G, B byte
}

type Point struct {
	X int
	Y int
}

func instantiatePoint() {
	var p Point
	p.X = 10
	p.Y = 20
	fmt.Println(p)
}

type Player struct {
	Name        string
	HealthPoint int
	MagicPoint  int
}

func instantiatePlayer() {
	tank := new(Player)
	tank.Name = "Canon"
	tank.HealthPoint = 300
	fmt.Println(tank)
}

type Command struct {
	Name    string // 指令名称
	Var     *int   // 指令绑定的变量
	Comment string // 指令的注释
}

func addressOfCommand() {
	var version int = 1
	cmd := &Command{}
	cmd.Name = "version"
	cmd.Var = &version
	cmd.Comment = "show version"
	fmt.Println(cmd)
}

type People struct {
	name  string
	child *People
}

func initPeople() {
	relation := &People{
		name: "爷爷",
		child: &People{
			name: "爸爸",
			child: &People{
				name: "我",
			},
		},
	}
	fmt.Println(relation)
}

type Address struct {
	Province    string
	City        string
	ZipCode     int
	PhoneNumber string
}

func InitAddress() {
	addr := Address{
		"四川",
		"成都",
		610000,
		"0",
	}
	fmt.Println(addr)
}

func initAnonymous() {
	// 实例化一个匿名结构体
	msg := &struct { // 定义部分
		id   int
		data string
	}{ // 值初始化部分
		1024,
		"hello",
	}
	printMsgType(msg)
}

// 打印消息类型, 传入匿名结构体
func printMsgType(msg *struct {
	id   int
	data string
}) {
	// 使用动词%T打印msg的类型
	fmt.Printf("%T\n", msg)
}

type Cat struct {
	Color string
	Name  string
}

// NewCatByName 模拟改造方法
func NewCatByName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}
func NewCatByColor(color string) *Cat {
	return &Cat{
		Color: color,
	}
}

type BlackCat struct {
	Cat // 嵌入Cat, 类似于派生
}

// NewCat “构造基类”
func NewCat(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

// NewBlackCat “构造子类”
func NewBlackCat(color string) *BlackCat {
	cat := &BlackCat{}
	cat.Color = color
	return cat
}

// 类型内嵌
type innerS struct {
	in1 int
	in2 int
}
type outerS struct {
	b      int
	c      float32
	int    // anonymous field
	innerS //anonymous field
}

func initInner() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10
	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)
	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is:", outer2)
}

// A 内嵌结构体
type A struct {
	ax, ay int
}
type B struct {
	A
	bx, by float32
}

func initAB() {
	b := B{A{1, 2}, 3.0, 4.0}
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)
}

// Wheel 车轮
type Wheel struct {
	Size int
}

// Engine 引擎
type Engine struct {
	Power int    // 功率
	Type  string // 类型
}

// Car 车
type Car struct {
	Wheel
	Engine
}

func initCar() {
	c := Car{
		// 初始化轮子
		Wheel: Wheel{
			Size: 18,
		},
		// 初始化引擎
		Engine: Engine{
			Type:  "1.4T",
			Power: 143,
		},
	}
	fmt.Printf("%+v\n", c)
}

//垃圾回收和SetFinalizer

type Road int

func findRoad(r *Road) {
	log.Println("road:", *r)
}
func entry() {
	var rd Road = Road(999)
	r := &rd
	runtime.SetFinalizer(r, findRoad)
}
func roadGC() {
	entry()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		runtime.GC()
	}
}

//链表操作
