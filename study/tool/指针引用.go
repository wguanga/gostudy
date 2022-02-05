package tool

import "fmt"

func DataType() {
	//指针
	a := 10
	b := &a
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("a: %T\n", a)
	fmt.Printf("b: %T\n", b)
	//数组定义
	array_1 := [2]int{1, 2}
	fmt.Printf("array_1: %T\n", array_1)
	fmt.Printf("array_1[1]: %T\n", array_1[1])
	//切片类型，理解为动态数组
	array_2 := []int{1, 2, 3}
	fmt.Printf("array_2: %T\n", array_2)
	//变量数据类型
	/* var a1 int8 =10
	var a2 int16=10
	var a3 int32=10
	var a4 int64=10
	var b1 uint8=10
	var b2 uint16=10
	var b3 uint32=10
	var b4 uint64=10
	var c1 float32=10
	var c2 float64=10 */

	//进制输出
	output := 10
	fmt.Printf("output: %d\n", output) //十进制输出
	fmt.Printf("output: %b\n", output) //二进制输出
	fmt.Printf("output: %o\n", output) //八进制输出
	fmt.Printf("output: %x\n", output) //十六进制输出
}

func Array_addr() {
	//指向数组的指针定义
	var array_addr []*int
	array_1 := []int{10, 20}
	array_addr = make([]*int, len(array_1))
	for i := 0; i < len(array_1); i++ {
		array_addr[i] = &array_1[i]
	}

	for i := 0; i < len(array_addr); i++ {
		fmt.Printf("*array_addr[%v]: %v\n", i, *array_addr[i])
	}
}
