package sqlc

import (
	"context"
)

type CreateUserTxParams struct {
	CreateUserParams
	AfterCreate func(user User) error
}

type CreateUserTxResult struct {
	User User
}

func (store *SQLStore) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	var result CreateUserTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.User, err = q.CreateUser(ctx, arg.CreateUserParams)
		if err != nil {
			return err
		}
		// 这里的回调函数是将发送一封证明邮件给刚创建的用户的任务加入
		// 到消息队列，但是入队到事务提交可能有较大的延迟，因此需要在
		// 处理任务时不能马上就取出处理，而是要有个延迟，否则处理任务
		// 时需要在数据库获取用户的邮箱，会出错
		return arg.AfterCreate(result.User)
	})

	return result, err
}
