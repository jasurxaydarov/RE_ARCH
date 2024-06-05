package postgres

import (
	"arch/modles"
	"arch/repoi"
	"context"
	"log"

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


func (t *TodoRepo)CreateTodo(ctx context.Context,todo modles.Todo,) error{

	query:=`
		INSERT INTO
			todo(
				todo_id,
				task,
				user_name
			)VALUES(
				$1,
				$2,
				$3
				
				
			)
	`

	_,err:=t.conn.Exec(
		ctx,query,
		todo.Todo_id,
		todo.Task,
		todo.User_name,
	)

	if err != nil{

		log.Println("err on CreateTodo ",err)
		return err
	}

	return nil
}


func (t *TodoRepo)GetTodos(ctx context.Context, limit int, page int) (*[]modles.Todo, error){

	query:=`
		SELECT 
			*
		FROM
			todo
		LIMIT 
			$1
		OFFSET
			$2
	`
	var (

		todo 	modles.Todo
		Todos	[]modles.Todo

	)

	rows,err:=t.conn.Query(ctx,query,limit,page)

	
	
	if err != nil{

		log.Println("err on GetTodos ",err)
		return nil,err
	}

	for rows.Next(){

		rows.Scan(
			&todo.Todo_id,
			&todo.Task,
			&todo.Created_at,
			&todo.Is_completed,
			&todo.User_name,
		)

		Todos=append(Todos,todo)
	
	}
	
	
	return &Todos,nil
}

func (t *TodoRepo)GetTodoByTask(ctx context.Context, task string) (*modles.Todo, error){

	query:=`
		SELECT 
			*
		FROM
			todo
		WHERE 
			task=$1
	`
	var Todo modles.Todo

	err:=t.conn.QueryRow(
		ctx,
		query,
		task,
		).Scan(
			&Todo.Todo_id,
			&Todo.Task,
			&Todo.Created_at,
			&Todo.Is_completed,
			&Todo.User_name,
		)

	if err != nil{

		log.Println("err on GetTodosByTask ",err)
		return nil,err
	}

	return &Todo,nil
}

func (t *TodoRepo)UpdateTodo(ctx context.Context, todo modles.Todo,task string) error{
	
	query:=`
		UPDATE 
			todo
		SET
			task=$1,
			is_completed=$2,
			user_name=$3
		WHERE 
			task=$4
			
	`

	_,err:=t.conn.Exec(
		ctx,query,
		todo.Task,
		todo.Is_completed,
		todo.User_name,
		task,
	)
	

	if err != nil{

		log.Println("err on UpdateTodo ",err)
		return err
	}

	return nil
}

func (t *TodoRepo)DeleteTodoByTask(ctx context.Context, task string) error{
	query:=`
		DELETE FROM
			todo
		WHERE
			task=$1
	`

	_,err:=t.conn.Exec(
		ctx,
		query,
		task,
	)

	if err!=nil {
		
		log.Println("error on DeleteUserBytask ", err)

		return err
	}
	
	return nil

	
}