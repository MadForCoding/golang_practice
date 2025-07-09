package main

import "fmt"

func main() {
	number := "1"
	for i := 0; i < 1000; i++ {
		number = multiplyStringByDigit(number, 2)
	}
	fmt.Println(number)
}

func multiplyStringByDigit(text string, base int) string {
	if len(text) == 0 || base == 0 {
		return "0"
	}
	res := []byte{}
	carry := 0
	for i := len(text) - 1; i >= 0; i-- {
		digitInt := int(text[i] - '0')
		product := digitInt*base + carry
		res = append([]byte{byte(product%10 + '0')}, res...)
		carry = product / 10
	}
	if carry > 0 {
		res = append([]byte{byte(carry + '0')}, res...)
	}
	return string(res)
}
