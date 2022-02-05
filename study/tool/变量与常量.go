//变量学习
package tool

import (
	"fmt"
)

func Variable() {
	//单独变量定义
	var x string = "abc"
	fmt.Printf("x: %v\n", x)

	//批量声明变量
	/* 方法一
	var (
		name string
		age  int
		male bool
	)
	name = "光"
	age = 22
	male = true */

	//方法二
	var (
		name string = "guang"
		age  int    = 22
		male bool   = true
	)
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("male: %v\n", male)

	//变量定义方法三,类型推断
	y := false //短变量声明，只能用于函数内部
	var z = "321"
	fmt.Printf("a: %v\n", y)
	fmt.Printf("z: %v\n", z)

	//简洁批量初始化变量
	var name_1, age_1, male_1 = "emit", 22, true
	fmt.Printf("name_1: %v\n", name_1)
	fmt.Printf("age_1: %v\n", age_1)
	fmt.Printf("male_1: %v\n", male_1)
}

//常量学习
func Constant() {
	//常量定义符const
	//常量定义时必须初始化赋值
	const name string = "guang"
	const age = 22 //类型推断，省略类型
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)

	//iota常量：初始为0，每调用一次iota则加1，遇到const关键字重置为0
	const a1, a2, a3 = iota, iota, iota
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)

	const (
		b1 = iota
		b2 = iota
		b3 = iota
	)
	fmt.Printf("b1: %v\n", b1)
	fmt.Printf("b2: %v\n", b2)
	fmt.Printf("b3: %v\n", b3)
	fmt.Printf("b3: %v\n", b3)
	//使用_跳过一个增加的值
	const (
		c1 = iota //0
		_         //1
		c2 = iota //2
	)
	fmt.Printf("c1: %v\n", c1)
	fmt.Printf("c2: %v\n", c2)
}
