package child

import "golang_practice/interface_research/baseInclude/base_interface/base_impl"

type Child struct {
	base_impl.Base
}

func (c *Child) AAAConfig() bool {
	return true
}
