package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context){
		for{
			select {
			case <-ctx.Done():
				fmt.Println("function 1 Done...")
				return
			default:
				fmt.Println("function 1 continue monitor...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	go func(ctx context.Context){
		for{
			select {
			case <-ctx.Done():
				fmt.Println("function 2 Done...")
				return
			default:
				fmt.Println("function 2 continue monitor...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("Ok, time to stop working")
	cancel()

	time.Sleep(2 * time.Second)



}
