package interEntity

import "golang_practice/interface_research/inter1/interMsg"

type P struct {
	interMsg.Tb
}

func (p *P) A() string{
	return "A"
}


