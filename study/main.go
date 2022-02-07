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

	/* //二叉树操作
	bt := new(algorithm.Btree)
	arry := make([]int, 0)
	arry = append(arry, 5, 2, 1, 6, 3)
	for _, v := range arry {
		bt.Breeinsert(v)
	}
	arry_2 := make([]int, 0)
	arry_2 = append(arry_2, 1, 3, 2, 6, 5)
	//var arry_3 []int
	//p_arr := &arry_3
	//algorithm.Postorder_traversal(bt.Root, p_arr)
	//fmt.Printf("p_arr: %v\n", p_arr)

	fmt.Printf("arry_2: %v\n", arry_2)

	fmt.Printf("algorithm.Jugselbtree(bt.Root, arry_2): %v\n", algorithm.Jugselbtree(bt.Root, arry_2))
	fmt.Printf("bt: %v\n", bt.Root) */

	/* //快排
	arry := []int{5, 6, 8, 2, 4, 9, 3}
	algorithm.Kuaipai(arry, 0, len(arry)-1)
	fmt.Printf("arry: %v\n", arry) */

	arry_1 := new([]int)
	arry_2 := make([]int, 0)
	arry_3 := []int{}
	fmt.Printf("arry_1: %T\n", arry_1)
	fmt.Printf("arry_2: %T\n", arry_2)
	fmt.Printf("arry_3: %T\n", arry_3)
}
