package model

type Login struct {
	Email    string `json:"login"`
	Password string `json:"password"`
}

type ChangePassword struct {
	Email string `json:"email"`
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}

type Register struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
}

type Update struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Email    string `json:"email"`
}

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Email    string `json:"email"`
}
