package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	//easyTry()
	reflectValue()
}

func easyTry() {
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) // "*os.File"
}

func reflectValue() {
	x := 2
	// Elem() 必须是 Ptr 或者是 Interface
	d := reflect.ValueOf(x).Elem()   // d refers to the variable x
	px := d.Addr().Interface().(*int) // px := &x
	*px = 3                           // x = 3
	fmt.Println(x)
}