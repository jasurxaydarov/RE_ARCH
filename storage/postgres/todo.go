package postgres

import (
	"arch/modles"
	"arch/repoi"
	"context"

	"github.com/jackc/pgx/v5"
)

type TodoRepo struct{
	
	conn *pgx.Conn 
}

func NewTodoRepo( conn *pgx.Conn)repoi.TodoRepoI{
	
	return &TodoRepo{
		 conn: conn,
		}
}


func (t *TodoRepo)CreateTodo(ctx context.Context, user modles.User) error{

	return nil
}


func (t *TodoRepo)GetTodos(ctx context.Context, limit int, page int) (*[]modles.Todo, error){

	return nil,nil
}

func (t *TodoRepo)GetTodoBy(ctx context.Context, todo_id int) (*modles.Todo, error){

	return nil,nil
}

func (t *TodoRepo)UpdateTodo(ctx context.Context, todo modles.Todo) error{

	return nil
}

func (t *TodoRepo)DeleteTodoById(ctx context.Context, todo_id int) error{

	return nil
}