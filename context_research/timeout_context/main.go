package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 500 * time.Millisecond)
	defer cancel()

	go handle(ctx, 520 * time.Millisecond)

	select{
	case <- ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration){
	select{
		case <- ctx.Done():
			fmt.Println("handle", ctx.Err())
		case <- time.After(duration):
			fmt.Println("process request with", duration)

	}
}
