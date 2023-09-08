package impl

import (
	"context"
	"github.com/luyasr/hello-world/apps/user"
)

var (
	_ user.Service = (*Impl)(nil)
)

type Impl struct {
}

func NewImpl() *Impl {
	return &Impl{}
}

func (i *Impl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.User, error) {
	return &user.User{Id: 1,
		CreateUserRequest: req,
	}, nil
}

func (i *Impl) GetUser(ctx context.Context, id int) (*user.User, error) {
	return &user.User{Id: id,
		CreateUserRequest: &user.CreateUserRequest{Username: "tom", Password: "12345"},
	}, nil
}
