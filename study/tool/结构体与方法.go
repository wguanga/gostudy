package tool

import "fmt"

type struct_1 struct {
	value int
	next  *struct_1
}

func structure() {
	var a struct_1    //定义a结点
	var p_a *struct_1 //a结点的指针
	p_a = &a          //连接a与p_a
	a.value = 10
	var b struct_1    //定义b结点
	var p_b *struct_1 //b结点的指针
	b.value = 20
	p_b = &b     //连接b与p_b
	a.next = p_b //b为a的下一个
	c := a.next  //c=p_b
	d := *c      //d=b
	fmt.Printf("a: %v\n", a)
	fmt.Printf("*p_a: %v\n", *p_a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("*p_b: %v\n", *p_b)
	fmt.Printf("d: %v\n", d)

	fmt.Printf("a.value: %v\n", a.value)
	fmt.Printf("p_a.value: %v\n", p_a.value)
	fmt.Printf("(*p_a).value: %v\n", (*p_a).value)

	p_c := new(struct_1)
	fmt.Printf("p_c: %p\n", p_c)

}

func (str_1 struct_1) my_mythod(a int) {
	fmt.Printf("a: %v\n", a)
	str_1.value += a
}

func (p_str1 *struct_1) myp_mythod(a int) {
	p_str1.value += a
}

func Method() {
	var a struct_1
	a = struct_1{
		value: 11,
	}
	fmt.Printf("a.value: %v\n", a.value)

	a.my_mythod(10) //接收者为值变量
	fmt.Printf("a.value: %v\n", a.value)

	a.myp_mythod(10) //接收者为指针变量
	fmt.Printf("a.value: %v\n", a.value)

	b := &struct_1{
		value: 20,
	}
	fmt.Printf("b.value: %v\n", b.value)

	b.my_mythod(10)
	fmt.Printf("b.value: %v\n", b.value)

	b.myp_mythod(10)
	fmt.Printf("b.value: %v\n", b.value)
}

//通过结构体嵌套实现继承

type animal struct {
	name string
	age  int
}

func (a animal) eat() {
	fmt.Println("eat...")
}

type bird struct {
	animal
}

type eager struct {
	animal
}
