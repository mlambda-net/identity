package entity

import "github.com/google/uuid"

type Identity struct {
  ID   uuid.UUID
  Name string
  LastName string
  Email    string
  Password string
  Roles    []*Role
}

func (i *Identity) RoleIds() []string {
  roles := make([]string,0)
  for _, role := range i.Roles {
    roles = append(roles, role.ID.String())
  }
  return roles
}


func NewIdentityFromRegister(name string, lastName string, email string) *Identity {
  return &Identity{
		Name:     name,
		LastName: lastName,
		Email:    email,
	}
}
