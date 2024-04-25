package main

import (
	"fmt"
	mathClass "my-go/myMath"
	"strconv"
	"unsafe"
)

func main() {

	baseGrammarTest() //基础语法测试
	varTest()         // 变量测试
	constTest()       //常量测试
}

func baseGrammarTest() {
	var ka, ba int
	ka = 12
	ba = 13
	println(ka)
	if ka > ba {
		println(strconv.Itoa(ka) + ">" + strconv.Itoa(ba))
	} else {
		println(strconv.Itoa(ka) + "<" + strconv.Itoa(ba))
	}
	//result := add(2, 3)
	println("Hello, World!")
	println(mathClass.Add(1, 1))
	println(mathClass.Sub(1, 1))
	// %d 表示整型数字，%s 表示字符串
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	target_url := fmt.Sprintf(url, stockcode, enddate)
	println(target_url)

	var a string = "Roobt"
	println(a)
	var b, c int = 1, 2
	println(b, c)

	// 声明一个变量并初始化
	var aa = "RUNOOB"
	println(aa)
	// 没有初始化就为零值
	var bb int
	println(bb)
	// bool 零值为 false
	var cc bool
	println(cc)

	//以下几种类型为 <nil>
	//var aq *int
	//var aq []int
	//var aq map[string] int
	//var aq chan int
	//var aq func(string) int
	//var aq error // error 是接口
	//占位符
	fmt.Printf("%-10s %10s\n", "Left", "Right")
	fmt.Printf("%08d %8d\n", 12345, -67890)
	fmt.Printf("%.2f %.2e\n", 3.14159, 2.71828)

	//声明语句的写法只能在函数体中出现 , 全局声明用:
	/*var (
		a int
		b bool
	)*/
	f := "Runoob" // var f string = "Runoob"
	println(f)
	intVal := 1
	println(intVal)

	//测试内存地址值是否发生改变 0xc00000a160 0xc00000a168
	//实际上是在内存中又拷贝了一份j = 7
	i := 7
	j := i
	println(&i, &j)
}

var kb, pa int

func varTest() {
	kb, pa = 111, 333
	//抛弃赋值_为抛弃变量
	var cpp string
	_, cpp = 5, "Roop"
	println(cpp)

	//抛弃赋值_函数调用
	yu, _ := mathClass.Sum("56", 3)
	println(yu)
}

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a) //占用bytes字节
)

func constTest() {
	const NAME = "sam"
	const AGE, SEX = 18, "女"
	const HIGHT, WEIGHT = 1.72, 48.3
	const A, B = 18.5, 24
	//体脂率 BMI = 体重 / 身高²
	var bmi float32
	bmi = WEIGHT / (HIGHT * HIGHT)
	println(fmt.Sprintf("姓名: %s, 年龄: %d, 性别: %s, 体重指数: %.2f", NAME, AGE, SEX, bmi))

	if bmi >= A && bmi <= B {
		println("你身材真好!")
	} else if bmi < A {
		println("你该多吃点东西了!")
	} else {
		println("你吃的有点多了!")
	}
	// 常量枚举
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)
	println(Unknown, Female, Male)
	println(a, b, c)

	//特殊常量 第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1
	//iota 相当于索引号
	const (
		aa = iota
		bb = iota
		cc = iota
	)
	println(aa, bb, cc)

}
