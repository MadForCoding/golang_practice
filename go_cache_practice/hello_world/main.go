package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	user2 "golang_practice/interface_research/user"
	"time"
)


func main() {
	c := cache.New(5 * time.Minute, 10 * time.Minute)

	c.Set("foo", "bar", cache.DefaultExpiration)

	foo, found := c.Get("foo")
	if found{
		fmt.Println(foo)
	}

	user := &user2.User{Na: "chen"}

	fmt.Println(checkType(user))

}

func check(us *user2.User) string{
	var n user2.GetName = us
	if in, ok := n.(user2.GetObject); ok {
		fmt.Println(in)
	}

	fmt.Println(n.Name())
	var o user2.GetObject = us
	fmt.Println(o.Ob())

	return ""
}

func checkType(us interface{}) string{
	switch v := us.(type){
	case user2.GetName:
		return fmt.Sprintln(v.Name())
	case user2.GetObject:
		return fmt.Sprintln(v.Ob())
	}
	return ""
}
