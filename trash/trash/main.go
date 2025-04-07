package main

import (
	"fmt"
	"sync"
)

func main() {
	var s = []string{
		"kk",
	}
	f := Appp(&s)
	fmt.Println(s)
	fmt.Println(f)
	fmt.Println("-------------")
	ss := &Stu{}
	ss.m.Store("22", "bbb")
	ss.m.Range(func(key, value interface{}) bool {
		fmt.Println(key,value)
		return true
	})
}

func Appp(slice *[]string) []string {
	*slice = append(*slice, "chen")
	return nil
}

type Stu struct {
	m sync.Map
}