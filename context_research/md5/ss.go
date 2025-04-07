package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"reflect"
)

func main() {
	s, _ := ioutil.ReadFile("/Users/wei/jwtRS256.key")

	srcCode := md5.Sum([]byte(s))
	code := fmt.Sprintf("%x", srcCode)
	fmt.Println(string(code))
	x := 2
	d := reflect.ValueOf(&x)
	f := d.Elem()
	fmt.Println(f)
}
