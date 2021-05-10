package main

import (
	"context"
	"fmt"
	"time"
)

var key string = "chen"

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	child := context.WithValue(ctx, key, "wei")

	go watch(child)

	time.Sleep(10 * time.Second)
	fmt.Println("Ok, stop to work")

	cancel()

	time.Sleep(2 * time.Second)
	fmt.Println("STOP")


}

func watch(ctx context.Context){
	for{
		select{
		case <-ctx.Done():
			fmt.Println(ctx.Value(key).(string), " Time to stop...")
			return
		default:
			fmt.Println(ctx.Value(key))
			time.Sleep(2 * time.Second)
		}
	}
}
