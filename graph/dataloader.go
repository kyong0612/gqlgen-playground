package graph

import (
	"context"
	"errors"
	"gqlgen-playground/graph/gql"
	"gqlgen-playground/service"

	dataloader "github.com/graph-gophers/dataloader/v7"
)

type Loaders struct {
	UserLoader dataloader.Interface[string, *gql.User]
}

func NewLoaders(Srv service.Services) *Loaders {
	userBatcher := &userBatcher{Srv: Srv}

	return &Loaders{
		// dataloader.Loader[string, *gql.User]構造体型をセットするために、
		// dataloader.NewBatchedLoader関数を呼び出す
		UserLoader: dataloader.NewBatchedLoader[string, *gql.User](userBatcher.BatchGetUsers),
	}
}

type userBatcher struct {
	Srv service.Services
}

func (u *userBatcher) BatchGetUsers(ctx context.Context, IDs []string) []*dataloader.Result[*gql.User] {
	// 引数と戻り値のスライスlenは等しくする
	results := make([]*dataloader.Result[*gql.User], len(IDs))
	for i := range results {
		results[i] = &dataloader.Result[*gql.User]{
			Error: errors.New("not found"),
		}
	}

	// 検索条件であるIDが、引数でもらったIDsスライスの何番目のインデックスに格納されていたのか検索できるようにmap化する
	indexs := make(map[string]int, len(IDs))
	for i, ID := range IDs {
		indexs[ID] = i
	}

	// サービス層のメソッドを使い、指定されたIDを持つユーザーを全て取得する
	// (ListUsersByIDメソッド内では、IN句を用いたselect文が実行されている)
	users, err := u.Srv.ListUsersByID(ctx, IDs)

	// 取得結果を、戻り値resultの中の適切な場所に格納する
	for _, user := range users {
		var rsl *dataloader.Result[*gql.User]
		if err != nil {
			rsl = &dataloader.Result[*gql.User]{
				Error: err,
			}
		} else {
			rsl = &dataloader.Result[*gql.User]{
				Data: user,
			}
		}
		results[indexs[user.ID]] = rsl
	}
	return results
}
