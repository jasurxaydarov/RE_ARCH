package main

import (
	"arch/api"
	"arch/config"
	"arch/storage/postgres"
	"context"
	"fmt"
	"log"
)

func main(){

	fmt.Println("hello git")

	var ctx =context.Background()
		
	cfg :=config.NEwConfig()

	conn,err:=postgres.Conn(cfg)

	if err!=nil{
		log.Println("err on con to db ",err)
		return 
	}

	fmt.Println(conn)
	
	defer conn.Close(ctx)

	UserRepo:= postgres.NewUserRepo(conn)
	TodoRepo:=postgres.NewTodoRepo(conn)
	fmt.Println("errrrrrrrrrrrrr")
	api.Api(UserRepo,TodoRepo)

}