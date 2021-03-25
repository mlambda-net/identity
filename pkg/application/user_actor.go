package application

import (
	"errors"
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/mlambda-net/identity/pkg/application/message"
	"github.com/mlambda-net/identity/pkg/domain/services"
	"github.com/mlambda-net/identity/pkg/domain/spec"
	"github.com/mlambda-net/identity/pkg/domain/utils"
	"github.com/mlambda-net/net/pkg/core"
)

type userActor struct {
  service services.UserService
  query   services.QueryService
}

func NewUserProps(config *utils.Configuration) *actor.Props {
	service := services.NewUserService(config)
	query := services.NewQueryService(config)
	return actor.PropsFromProducer(func() actor.Actor { return &userActor{service: service, query: query} })
}

func (u *userActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *message.Filter:
		r, err := u.userQuery(msg)
		if err != nil {
			ctx.Respond(err)
		} else {
			ctx.Respond(r)
		}

	case *message.Update:
		email := ctx.MessageHeader().Get("sub")
		if email == msg.Email {
			u, err := u.update(msg)
			if err != nil {
				ctx.Respond(err)
			} else {
				ctx.Respond(u)
			}
		} else {
			ctx.Respond(errors.New("the email doesn't match with the actual user"))
		}

	case *message.Delete:
		id := ctx.MessageHeader().Get("id")
		if id == fmt.Sprintf("%d", msg.Id) {
			err := u.deleteUser(msg)
			if err != nil {
				ctx.Respond(err)
			} else {
				ctx.Respond(core.Unit())
			}
		} else {
			ctx.Respond(errors.New("the id doesn't match with the actual user"))
		}

	case *message.ChangePassword:
		email := ctx.MessageHeader().Get("sub")
		if email == msg.Email {
			err := u.changePassword(msg)
			if err != nil {
				ctx.Respond(err)
			} else {
				ctx.Respond(core.Unit())
			}
		} else {
			ctx.Respond(errors.New("the user email doesn't match with the actual user"))
		}

	case *message.Find:
		email := ctx.MessageHeader().Get("sub")
		if email == msg.Email {
			usr, err := u.Find(msg)
			if err != nil {
				ctx.Respond(err)
			} else {
				ctx.Respond(usr)
			}
		} else {
			ctx.Respond(errors.New("the user email doesn't match with the actual user"))
		}
	}
}

func (u *userActor) userQuery(msg *message.Filter) (*message.Results, *core.Error) {
	switch msg.By {
	case message.EMAIL:
		return u.filter(spec.ByEmail(msg.Email))
	case message.ID:
		return u.filter(spec.ById(msg.Id))
	case message.NAME:
		return u.filter(spec.ByName(msg.Name, msg.LastName))
	}
	return nil, &core.Error{
		Status:  500,
		Message: "not search type specify",
	}
}

func (u *userActor) filter(spec spec.Expression) (*message.Results, *core.Error) {

	users, err := u.query.Query(spec)
	if err != nil {
		return nil, &core.Error{
			Status:  500,
			Message: err.Error(),
		}
	}

	results := make([]*message.Result, 0)

	for _, user := range users {
		r := &message.Result{
			Id:       user.Id,
			Name:     user.Name,
			LastName: user.LastName,
			Email:    user.Email,
		}
		results = append(results, r)
	}

	return &message.Results{Results: results}, nil
}

func (u *userActor) deleteUser(msg *message.Delete) *core.Error {

	err := u.service.Delete(msg.Id)
	if err != nil {
		return &core.Error{
			Status:  500,
			Message: err.Error(),
		}
	}

	return nil

}

func (u *userActor) update(msg *message.Update) (*message.Response, *core.Error) {

	user, err := u.query.ByEmail(msg.Email)
	if err != nil {
		return nil, &core.Error{
			Status:  500,
			Message: err.Error(),
		}
	}

	user.Name = msg.Name
	user.LastName = msg.LastName

	err = u.service.Update(user)
	if err != nil {
		return nil, &core.Error{
			Status:  500,
			Message: err.Error(),
		}
	}
	return &message.Response{Id: user.Id}, nil
}

func (u *userActor) changePassword(msg *message.ChangePassword)  *core.Error {

  usr, err := u.query.ByEmail(msg.Email)
  if err != nil {
    return  &core.Error{
      Status:  500,
      Message: err.Error(),
    }
  }

  old, err := utils.Decrypt(usr.Password)

  if err != nil {
    return  &core.Error{
      Status:  500,
      Message: err.Error(),
    }
  }


  if old != msg.OldPassword  {
    return  &core.Error{
      Status:  500,
      Message: "the password mismatch",
    }
  }


 err = u.service.ChangePassword(msg.Email, msg.Password)
  if err != nil {
    return  &core.Error{
      Status:  500,
      Message: err.Error(),
    }
  }

  return nil

}

func (u *userActor) Find(msg *message.Find) (*message.Result, *core.Error) {
	usr, err := u.query.ByEmail(msg.Email)
	if err != nil {
		return nil, &core.Error{
			Status:  500,
			Message: "The user can not be retrieved",
		}
	} else {
		return &message.Result{
			Id:       usr.Id,
			Name:     usr.Name,
			LastName: usr.LastName,
			Email:    usr.Email,
		}, nil
	}
}

