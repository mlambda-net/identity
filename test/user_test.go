package test

import (
	"fmt"
	"github.com/mlambda-net/identity/pkg/application/message"
	"github.com/mlambda-net/identity/pkg/infrastructure/endpoint/server"
	"github.com/mlambda-net/net/pkg/net"
	"github.com/mlambda-net/net/pkg/security"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_Create_User(t *testing.T) {
	s := server.NewServer()
	s.Start()
	t.Run("create", create)
	t.Run("query", query)
	t.Run("update", update)
	t.Run("change password", changePassword)
	t.Run("delete", delete)
}

func changePassword(t *testing.T) {

	user := &message.Create{
		Name:     "Roy",
		LastName: "Gonzalez",
		Email:    "coyote@acme.com",
		Password: "123",
	}

	c := net.NewClient("localhost", "9001")
	tk := getToken(map[string]interface{}{})


	ul := c.Actor("user")
	auth := c.Actor("auth")
 	auth.Request(user)

	ul.Token(tk).Request(&message.ChangePassword{
		Email:       "coyote@acme.com",
		Password:    "132",
		OldPassword: "123",
	})



	rsp, err := auth.Request(&message.Authenticate{
		Email:    "coyote@acme.com",
		Password: "132",
	}).Unwrap()

	token := rsp.(*message.Token).Value
	assert.NotNil(t, token)
	assert.Nil(t, err)

}

func query(t *testing.T) {
	query := &message.Filter{
		Email: "yordivad@gmail.com",
		By:    message.EMAIL,
	}


	c := net.NewClient("localhost", "9001")
	r, err := c.Actor("user").Token(getToken(map[string]interface{}{})).Request(query).Unwrap()

	assert.Nil(t, err)
	assert.NotNil(t, r)

}

func create(t *testing.T) {
	user := &message.Create{
		Name:     "Roy",
		LastName: "Gonzalez",
		Email:    "yordivad@gmail.com",
		Password: "123",
	}

	c := net.NewClient("localhost", "9001")
	r, err := c.Actor("auth").Request(user).Unwrap()
	assert.Nil(t, err)
	assert.NotNil(t, r)
}

func delete(t *testing.T) {
	c := net.NewClient("localhost", "9001")
	users, err := filter(c)
	assert.Nil(t, err)

	for _, user := range users {
		claims :=map[string]interface{}{"id": user.Id}
		c.Actor("user").Token(getToken(claims)).Request(&message.Delete{Id: user.Id})
	}

	users, err = filter(c)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(users))
}

func filter(c net.Client) ([]*message.Result, error) {

	query := &message.Filter{
		Email: "yordivad@gmail.com",
		By:    message.EMAIL,
	}
	claims :=map[string]interface{}{}
	r, err := c.Actor("user").Token(getToken(claims)).Request(query).Unwrap()
	users := r.(*message.Results)
	return users.Results, err
}


func update(t *testing.T) {
	c := net.NewClient("localhost", "9001")
	users, err := filter(c)
	assert.Nil(t, err)
	assert.True(t, len(users) > 0)
	first := users[0]
	newUser := &message.Update{
		Email:       first.Email,
		Name:     "Name",
		LastName: "Last Name",
	}

	claims :=map[string]interface{}{"email": first.Email}
	c.Actor("user").Token(getToken(claims)).Send(newUser)
	users, err = filter(c)
	assert.Nil(t, err)
	user := users[0]
	assert.Equal(t, "Name", user.Name)
	assert.Equal(t, "Last Name", user.LastName)

}

func getToken(claims map[string]interface{})  string {
	claims["authorize"] = true
	tk := security.NewToken(os.Getenv("SECRET_KEY"))
	token, _ := tk.Create(claims)
	t := fmt.Sprintf("bearer: %s", token)
	return t
}
