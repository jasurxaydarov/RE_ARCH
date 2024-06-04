package repoi

import (
	"arch/modles"
	"context"
)

type UserRepoI interface{
	CreateUser(ctx context.Context,user modles.User)error
	GetUsers(ctx context.Context,limit int,page int)(*[]modles.User,error)
	GetUserById(ctx context.Context,user_id int)(*modles.User,error)
	UpdateUser(ctx context.Context,users modles.User)error
	DeleteUserById(ctx context.Context,user_id int)error

}

type TodoRepoI interface{
	CreateTodo(ctx context.Context,user modles.User)error
	GetTodos(ctx context.Context,limit int,page int)(*[]modles.Todo,error)
	GetTodoBy(ctx context.Context,todo_id int)(*modles.Todo,error)
	UpdateTodo(ctx context.Context,todo modles.Todo)error
	DeleteTodoById(ctx context.Context,todo_id int)error
		
}