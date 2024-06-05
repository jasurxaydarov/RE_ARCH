package repoi

import (
	"arch/modles"
	"context"

)

type UserRepoI interface{
	CreateUser(ctx context.Context, user modles.User) error
	DeleteUserByName(ctx context.Context, user_name string) error
 	GetUserByName(ctx context.Context, user_name string) (*modles.User, error)
	GetUsers(ctx context.Context, limit int, page int) (*[]modles.User, error)
	UpdateUser(ctx context.Context, user modles.User, user_name string) error


}

type TodoRepoI interface{
	CreateTodo(ctx context.Context,todo modles.Todo)error
	GetTodos(ctx context.Context,limit int,page int)(*[]modles.Todo,error)
	GetTodoByTask(ctx context.Context,task string)(*modles.Todo,error)
	UpdateTodo(ctx context.Context,todo modles.Todo,task string)error
	DeleteTodoByTask(ctx context.Context,task string)error
		
}

