package entity

type Identity struct {
	Id       int64
	Name     string
	LastName string
	Email    string
	Password string
}

func NewIdentityFromRegister(name string, lastName string, email string) *Identity {

	return &Identity{
		Name:     name,
		LastName: lastName,
		Email:    email,
	}
}
