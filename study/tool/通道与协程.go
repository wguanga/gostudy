package tool

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var ch_1 = make(chan int)

func setchan() {
	defer wg.Done() //标记协程数减一
	rand.Seed(time.Now().Unix())
	a := rand.Intn(20)
	ch_1 <- a
	fmt.Printf("Set a: %v\n", a)
}

func getchan() {
	defer wg.Done() //标记协程数减一
	a := <-ch_1
	fmt.Printf("Get a: %v\n", a)
}

func Chanandxiecheng() {

	defer close(ch_1)

	wg.Add(1) //标记协程数加1
	go setchan()

	wg.Add(1) //标记协程数加1
	go getchan()

	wg.Wait() //如果协程数不等于0，则让权等待

}
