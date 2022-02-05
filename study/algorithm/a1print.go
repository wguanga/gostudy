package algorithm

import (
	"fmt"
	"sync"
)

func A1print() {
	//交替打印
	wg := sync.WaitGroup{}

	defer wg.Wait()
	ch1 := make(chan byte, 1)
	ch2 := make(chan byte, 1)
	wg.Add(1)
	go func() {

		for i := 0; i < 10; i++ {
			ch1 <- 'a'
			fmt.Printf("%v\n", i)
			<-ch2
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {

		for i := 'a'; i < 'a'+10; i++ {
			<-ch1
			fmt.Printf("%c\n", i)
			ch2 <- 'b'
		}
		wg.Done()
	}()
}
