package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	v := reflect.ValueOf(3)
	fmt.Println(v)

	t := reflect.TypeOf(v.Int())
	fmt.Println(t)

	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))

	v2 := reflect.ValueOf(3)
	fmt.Println(v2.String())

	x := 2
	dr := reflect.ValueOf(&x).Elem()
	px := dr.Addr().Interface().(*int)
	*px = 100
	fmt.Println(x)


}
