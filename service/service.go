package service

import "github.com/volatiletech/sqlboiler/v4/boil"

type Services interface {
	UserService
}

type services struct {
	*userService
	// issueテーブルを扱うissueServiceなど、他のサービス構造体ができたらフィールドを追加していく
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService: &userService{exec: exec},
	}
}
