package main

import (
	"fmt"
)

type student struct {
	Name string
	Age  int
}

func main() {
	/* var a = make(chan int, 1)
	fmt.Printf("a: %v\n", a)
	s := "abcss"
	t := tool.LengthOfLongestSubstring(s)
	fmt.Printf("t: %v\n", t) */

	m := make(map[string]*student)
	fmt.Printf("m: %T\n", m)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
		fmt.Printf("m[stu.Name]: %T\n", m[stu.Name])
	}

	//类型断言
	/* var in interface{}
	in = "1111"
	flag, ok1 := in.(string)
	_, ok2 := in.(int)
	fmt.Printf("flag: %v\n", flag)
	fmt.Printf("ok1: %v   ok1: %v\n", ok1, ok2)
	*/

	//两个栈实现队列
	/* 	cq := algorithm.Constructor()
	   	for i := 0; i < 10; i++ {
	   		cq.AppendTail(i)
	   	}
	   	for i := 0; i < 11; i++ {
	   		fmt.Printf("cq.DeleterHead(): %v\n", cq.DeleterHead())
	   	} */

	//错误处理

	/* condi_1 := 0
	//var fff func()
	fff := func(a int, b int) (float32, error) {
		if b == 0 {
			return 0, errors.New("除数为0")
		} else {
			return float32(a) / float32(b), nil
		}
	}

	if _, err := fff(10, condi_1); err != nil {
		panic(err)
	} */

	/* if condi_1 == 0 {
		panic(errors.New("除数为0"))
	} */

	//tool.Dbrelationshiptest()
}
