package tool

import (
	"fmt"
)

func func_1() {
	fmt.Println("无参无返回")
}

func func_2(a int, b int) (c int, d int) {
	return a + b, a - b
}

func func_3(args ...int) {
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
}

//定义函数类型

type commfun func(int, int) int //也可以定义在函数内

func sum(a int, b int) int {
	return a + b
}

func dec(a int, b int) int {
	return a - b
}

func pr(name string) {
	fmt.Printf("hello, %v !\n", name)
}

//函数作为参数和返回值
func stringjoin(ss string, f func(string)) func(string) {
	f(ss)
	return pr
}

//闭包
func closure_addd() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func Funcmain() {
	//定义匿名函数
	max := func(a int, b int) int {
		if a >= b {
			return a
		} else {
			return b
		}
	}
	abmax := max(1, 2)
	fmt.Printf("abmax: %v\n", abmax)

	func_1()
	c, d := func_2(4, 6)
	_, f := func_2(4, 6) //可以用下划线丢弃返回值

	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)
	fmt.Printf("f: %v\n", f)

	//使用函数类型和函数变量
	var f1, f2 commfun
	f1 = sum
	fmt.Printf("f1(1, 2): %v\n", f1(1, 2))
	f2 = dec
	fmt.Printf("f2(1, 2): %v\n", f2(1, 2))

	//使用闭包
	cf := closure_addd()
	fmt.Printf("cf(10): %v\n", cf(10))
	fmt.Printf("cf(10): %v\n", cf(10))

}
