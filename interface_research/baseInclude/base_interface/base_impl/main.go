package base_impl

import (
	"fmt"
	"golang_practice/interface_research/baseInclude/base_interface"
)

type Base struct {
	AInterface base_interface.A
}

func (b *Base) TryBase() {
	fmt.Println(b.AInterface.AAAConfig())
}