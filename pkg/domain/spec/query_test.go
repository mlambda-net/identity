package spec

import (
	"github.com/mlambda-net/identity/pkg/domain/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_FilterByName(t *testing.T) {

	user := getUser()

	spec := Specify(ByName(user.Name, user.LastName))

	assert.Equal(t, spec.Query(), "(name = 'Roy' AND last_name = 'Gonzalez')")

}

func getUser() entity.Identity {
	return entity.Identity{
		Id:       1,
		Name:     "Roy",
		LastName: "Gonzalez",
		Email:    "yordivad@gmai.com",
	}
}
