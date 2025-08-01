package service

import "context"

type (
	//interface is here
	IUserLogin interface {
		Login(ctx context.Context) error
	}

	IUserInfo interface {
		GetInfoUserId()
	}
)
