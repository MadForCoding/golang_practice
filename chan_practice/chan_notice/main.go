package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool,1)

	go func(){
		for{
			select {
			case <- stop:
				fmt.Println("Time to stop")
				return
			default:
				fmt.Println("Continue monitor signal")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("Ok, Let's stop to work")
	stop <- true

	time.Sleep(1 * time.Second)
}
