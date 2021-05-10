package interMsg

type Tb interface {
	A() string
	b() string
}

type BB struct {
	Tb
}

func (b *BB) A() string{
	return "A"
}

func (b *BB) b() string{
	return "b"
}