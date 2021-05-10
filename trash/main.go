package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var ab = make(chan struct{}, 1)
	fmt.Println(len(ab))
	ab <- struct{}{}
	fmt.Println(len(ab))
	fmt.Println("111111")
	var a int32 = 0
	for i := 0; i < 100; i++{
		go func(){
			atomic.AddInt32(&a, 1)
			fmt.Println(atomic.LoadInt32(&a))
			//fmt.Println(a)
			//atomic.AddInt32(&a, -1)
		}()
	}

	time.Sleep(time.Second)
	fmt.Println(atomic.LoadInt32(&a))


}
