package handlers

import (
	"arch/modles"
	"arch/repoi"

	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type Handlers struct{
	UserRepo	repoi.UserRepoI
	TodoRepo repoi.TodoRepoI
}

func NewHandlers(UserRepo repoi.UserRepoI,TodoRepo repoi.TodoRepoI)Handlers{

	return Handlers{
		UserRepo: UserRepo,
		TodoRepo :TodoRepo,
	}

}

func (h *Handlers)CreateUser(){
	
	var user  modles.User
	user.User_id=uuid.New()
	
	fmt.Println("enter user name ")
	fmt.Scanln(&user.User_name)

	fmt.Println("enter gmail ")
	fmt.Scanln(&user.Gmail)

	fmt.Println("enter age  ")
	fmt.Scanln(&user.Age)

	fmt.Println("enter password  ")
	fmt.Scanln(&user.Password)

	fmt.Println(h.UserRepo)
	err := h.UserRepo.CreateUser(context.Background(), user)

	if err!=nil{

		log.Println("error on create user",err)

		return 
	}

	fmt.Println("created")

}


func (h *Handlers)UpdateUser(){

	var user modles.User
	var user_name string

	fmt.Println("enter Updating user's name ")
	fmt.Scanln(&user_name)

	fmt.Println("enter Updating user's new name ")
	fmt.Scanln(&user.User_name)

	fmt.Println("enter Updating user's new gmail ")
	fmt.Scanln(&user.Gmail)

	fmt.Println("enter Updating user's new age ")
	fmt.Scanln(&user.Age)

	fmt.Println("enter Updating user's new password ")
	fmt.Scanln(&user.Password)


	err:= h.UserRepo.UpdateUser(context.Background(),user,user_name)

	if err != nil{
		
		log.Println("err on handlers UpdateUser ",err)
		return
	}

	fmt.Println("succusfully updated")

}

func (h *Handlers)DeleteUserByName(){
	
	var user_name string

	fmt.Println("enter deleting user name")
	fmt.Scanln(&user_name)

	err:= h.UserRepo.DeleteUserByName(context.Background(),user_name)

	if err != nil{
		
		log.Println("err on handlers  DeleteUserById",err)
		return
	}

	fmt.Println("succusfully deleted ")

}

func(h *Handlers)GetUsers(){

	var limit, page int
	limit=5

	fmt.Println("CHOSE PAGE ")
	fmt.Scanln(&page)
	
	users, err:= h.UserRepo.GetUsers(context.Background(),limit,page)

	if err != nil{
		
		log.Println("err on handlers GetUsers ",err)
		return
	}

	fmt.Println(users)
}

func (h *Handlers)GetUserByName(){
	var user_name string

	fmt.Println("enter getting user name")
	fmt.Scanln(&user_name)


	user,err:=h.UserRepo.GetUserByName(context.Background(),user_name)

	if err != nil{
		
		log.Println("err on handlers GetUserbyname ",err)
		return
	}

	fmt.Println(user)

}


