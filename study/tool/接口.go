package tool

import "fmt"

type usb interface {
	read()
	write()
}

type computer struct {
	name string
}

type phone struct {
	name string
}

func (com computer) read() {
	fmt.Printf("com.name: %v\n", com.name)
	fmt.Printf("computer read..\n")
}

func (com computer) write() {
	fmt.Printf("com.name: %v\n", com.name)
	fmt.Println("computer write..")
}

func (pho phone) read() {
	fmt.Printf("pho.name: %v\n", pho.name)
	fmt.Println("phone read")
}

func (pho phone) write() {
	fmt.Printf("pho.name: %v\n", pho.name)
	fmt.Println("phone write")
}

//接口嵌套
type flyer interface {
	fly()
}

type animer interface {
	eat()
	sleep()
}

type flyanimer interface {
	flyer
	animer
}

func Interfaceqiantao() {
	var com usb
	com = computer{
		name: "联想",
	}
	com.read()
	var comm computer
	comm = computer{
		name: "戴尔",
	}
	comm.write()
}
