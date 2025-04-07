package tool

import "fmt"

type generateChar struct {
	Position int
}

var GenerateChar = &generateChar{
	Position: 65,
}

func (r *generateChar) GetNextCharacter() string {
	if r.Position >= 65+26 {
		panic("excess")
	}
	c := fmt.Sprintf("%c", r.Position)
	r.Position++
	return c
}

func (r *generateChar) Reset() {
	r.Position = 65
}
