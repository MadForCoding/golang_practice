package main

import (
	"context"
	"fmt"
	"time"
)

type s struct {
	cancel context.CancelFunc
}

func main() {
	ctx, cancel := tt(0)
	if ctx == nil {
		fmt.Println("OK")
		return
	}
	defer cancel()

}

func tt(number int) (context.Context, context.CancelFunc) {
	if number == 0 {
		return nil, nil
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		return ctx, cancel
	}
}
