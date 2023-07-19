package service

import (
	"context"
	"gqlgen-playground/graph/gql"
	models "gqlgen-playground/infra/db/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*gql.User, error)
	ListUsersByID(ctx context.Context, IDs []string) ([]*gql.User, error)
}

type userService struct {
	exec boil.ContextExecutor
}

func convertUser(user *models.User) *gql.User {
	return &gql.User{
		ID:   user.ID,
		Name: user.Name,
	}
}

func (u *userService) GetUserByName(ctx context.Context, name string) (*gql.User, error) {
	// 1. SQLBoilerで生成されたORMコードを呼び出す
	user, err := models.Users( // from users
		qm.Select(models.UserTableColumns.ID, models.UserTableColumns.Name), // select id, name
		models.UserWhere.Name.EQ(name),                                      // where name = {引数nameの内容}
	).One(ctx, u.exec) // limit 1
	// 2. エラー処理
	if err != nil {
		return nil, err
	}
	// 3. 戻り値の*model.User型を作る
	return convertUser(user), nil
}

// サービス層内に実装された、IN句を用いた取得処理
func (u *userService) ListUsersByID(ctx context.Context, IDs []string) ([]*gql.User, error) {
	users, err := models.Users(
		qm.Select(models.UserTableColumns.ID, models.UserTableColumns.Name),
		models.UserWhere.ID.IN(IDs),
	).All(ctx, u.exec)
	if err != nil {
		return nil, err
	}
	return convertUserSlice(users), nil
}

func convertUserSlice(users models.UserSlice) []*gql.User {
	result := make([]*gql.User, 0, len(users))
	for _, user := range users {
		result = append(result, convertUser(user))
	}
	return result
}
