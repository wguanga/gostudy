package tool

import "fmt"

func init() {
	fmt.Println("init...")
}

func DeferandInit() {
	//defer语句
	fmt.Println("111")
	defer fmt.Println("222")
	defer fmt.Println("333")
	defer fmt.Println("444")

}
