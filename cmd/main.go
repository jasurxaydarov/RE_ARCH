package main

import (
	"arch/config"
	"arch/storage/postgres"
	"context"
	"fmt"
	"log"
)

func main(){

	fmt.Println("hello git")

	var (
		cfg config.Config
		ctx =context.Background()
		)

	conn,err:=postgres.Conn(cfg)

	if err!=nil{
		log.Println("err on con to db ",err)
		return 
	}

	fmt.Println(conn)
	
	defer conn.Close(ctx)
}