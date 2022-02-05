package tool

import (
	"fmt"
)

//条件判断
func Judgement() {
	if flagg := false; !flagg {
		fmt.Printf("flagg: %v\n", flagg)
	}
	var name string
	var age int
	fmt.Printf("请输入name，age\n")
	/* fmt.Scan(&name, &age) */
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)

	a := 100
	switch {
	case a > 80:
		fmt.Println("优秀")
	case a > 70:
		fmt.Println("良好")
	case a > 60:
		fmt.Println("及格")
	case a > 0:
		fmt.Println("不及格")
	default:
		fmt.Println("default")
	}

	switch b := 100; b {
	case 10, 20, 30:
		fmt.Println("1")
	case 50, 100:
		fmt.Println("2")

	}

}

//循环
func Loop() {
	//for
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
	//for range
	array_1 := []int{10, 20, 30}
	fmt.Printf("array_1: %v\n", array_1)
	for i, v := range array_1 {
		fmt.Printf("index: %v  ", i)
		fmt.Printf("value: %v\n", v)
	} //i为索引，v为值
	for _, v := range array_1 {
		fmt.Printf("v: %v\n", v)
	} //只要值不要索引

}
