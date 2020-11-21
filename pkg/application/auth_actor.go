package application

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/mlambda-net/identity/pkg/application/message"
	"github.com/mlambda-net/identity/pkg/domain/entity"
	"github.com/mlambda-net/identity/pkg/domain/services"
	"github.com/mlambda-net/identity/pkg/domain/utils"
	"github.com/mlambda-net/net/pkg/core"
)

type authActor struct {
	service services.UserService
}

func NewAuthProps(config *utils.Configuration) *actor.Props {
	service := services.NewUserService(config)
	return actor.PropsFromProducer(func() actor.Actor { return &authActor{service: service} })
}

func (a *authActor) Receive(ctx actor.Context) {

	switch msg := ctx.Message().(type) {

	case *message.Create:
		r, err := a.createUser(msg)
		if err != nil {
			ctx.Respond(err)
		} else {
			ctx.Respond(r)
		}

	case *message.Authenticate:
		token, err := a.authenticate(msg)
		if err != nil {
			ctx.Respond(err)
		} else {
			ctx.Respond(token)
		}
	}
}

func (a *authActor) createUser(msg *message.Create) (*message.Response, *core.Error) {

	if msg.Password == "" {
		return nil, &core.Error{
			Status:  500,
			Message: "the password is empty",
		}
	}

	newPass, err := utils.Encrypt(msg.Password)
	if err != nil {
		return nil, &core.Error{
			Status:  500,
			Message: err.Error(),
		}
	}

	user := entity.NewIdentityFromRegister(msg.Name, msg.LastName, msg.Email)
	user.Password = newPass

	r, err := a.service.Create(user)
	if err != nil {
		return nil, &core.Error{
			Status:  500,
			Message: err.Error(),
		}
	}

	return &message.Response{Id: r.Id}, nil
}

func (a *authActor) authenticate(msg *message.Authenticate) (*message.Token, *core.Error) {

	token, err := a.service.Authenticate(msg.Email, msg.Password)
	if err != nil {
		return nil, &core.Error{
			Status:  500,
			Message: err.Error(),
		}
	} else {
		return &message.Token{Value: token}, nil
	}

}
