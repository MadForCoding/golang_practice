package user

type User struct{
	Na string
}

type GetName interface{
	Name() string
}

type GetObject interface {
	Ob() *User
}

func (u *User) Name() string{
	return u.Na
}

func (u *User) Ob() *User{
	return u
}


