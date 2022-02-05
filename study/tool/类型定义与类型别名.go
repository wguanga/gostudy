package tool

import "fmt"

func Abotype() {
	//类型定义
	type myint int
	var a myint
	a = 10
	fmt.Printf("a: %v\n", a)

	//类型别名
	type myint2 = int
	var b myint2
	b = 10
	fmt.Printf("b: %v\n", b)
}
