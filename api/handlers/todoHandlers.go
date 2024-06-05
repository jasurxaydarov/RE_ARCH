package handlers

import (
	"arch/modles"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
)


func(h *Handlers)CreateTodo(){
	
	
	var todo modles.Todo
	todo.Todo_id=uuid.New()

	fmt.Println("enter task")
	fmt.Scanln(&todo.Task)

	fmt.Println("enter users name")
	fmt.Scanln(&todo.User_name)

	err:=h.TodoRepo.CreateTodo(context.Background(),todo)

	if err != nil{
		log.Println("error on CreateTodo")
		return
	}

	fmt.Println("succesfully created ")

}

func(h *Handlers)UpdateTodo(){

	var (
		todo modles.Todo
		task string
	)

	fmt.Println("enter updating todo's task")
	fmt.Scanln(&task)

	fmt.Println("enter  new todo's task")
	fmt.Scanln(&todo.Task)

	fmt.Println("enter  new todo's is completed")
	fmt.Scanln(&todo.Is_completed)

	fmt.Println("enter  new todo's user_name")
	fmt.Scanln(&todo.User_name)

	err:=h.TodoRepo.UpdateTodo(context.Background(),todo,task)

	if err != nil{
		log.Println("error on UpdateTodo")
		return
	}

	fmt.Println("succesfully updated ")

}

func(h *Handlers)DeleteTodoByTask(){

	var task string

	fmt.Println("enter task")
	fmt.Scanln(&task)

	err:=h.TodoRepo.DeleteTodoByTask(context.Background(),task)

	if err != nil{
		log.Println("error on DeleteTodoByTask")
		return
	}

	fmt.Println("succesfully deleted  ")

}

func(h *Handlers)GetTodos(){

	var limit,page int=5,0

	fmt.Println("enter page")
	fmt.Scanln(&page)

	todos,err:=h.TodoRepo.GetTodos(context.Background(), limit, page)

	if err != nil{
		log.Println("error on GetTodos")
		return
	}

	fmt.Println(todos)

}

func(h *Handlers)GetTodoByTask(){

	var task string

	fmt.Println("enter task")
	fmt.Scanln(&task)

	todo,err:=h.TodoRepo.GetTodoByTask(context.Background(),task)

	if err != nil{
		log.Println("error on GetTodoByTask")
		return
	}

	fmt.Println(todo)

	
}