package api

import (
	"arch/api/handlers"
	"arch/repoi"

	"fmt"
)

func Api(UserREpo repoi.UserRepoI,TodoRepo repoi.TodoRepoI){
	
	opr:= -1

	h:= handlers.NewHandlers(UserREpo,TodoRepo)
	
	fmt.Println("welcome my todo app")
	
	for{

		fmt.Println("chose oprt")

		fmt.Println(
			`
			1.create user
			 2.update users
			 3.delete  user by name
			 4.get users
			 5.get user by name
			 
			 6.Create todo
			 7.update todo
			 8.delet todo
			 9.get todos
			 10.get todos by task
			 0.exit
			 `,
		)

		fmt.Scanln(&opr)

		switch opr{
		case 1:h.CreateUser()
		case 2:h.UpdateUser()
		case 3:h.DeleteUserByName()
		case 4:h.GetUsers()
		case 5:h.GetUserByName()
		case 6:h.CreateTodo()
		case 7:h.UpdateTodo()
		case 8:h.DeleteTodoByTask()
		case 9:h.GetTodos()
		case 10:h.GetTodoByTask()
		case 0:return
		}
	
	}
	
}
