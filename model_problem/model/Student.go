package model

import "fmt"

type Person struct {
	name string
	SS string
}

type Student struct {
	Person
	name string
}

func NewStu() Student{
	s := Student{
		Person: Person{SS: "c", name: "l"},
		name: "chen",

	}
	fmt.Println(s.name)
	fmt.Println(s.Person.name)
	fmt.Println(s.SS)
	return s

}