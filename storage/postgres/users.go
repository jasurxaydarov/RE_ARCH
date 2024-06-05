package postgres

import (
	"arch/modles"
	"arch/repoi"
	"context"
	"log"
	"github.com/jackc/pgx/v5"
)

type UserRepo struct {

	conn *pgx.Conn
}

func NewUserRepo(conn *pgx.Conn)repoi.UserRepoI{

	return &UserRepo{
			conn: conn,
	}


}

func(u *UserRepo) CreateUser(ctx context.Context, user modles.User) error{

	query:=
		`INSERT INTO
			Users(
				user_id,
				user_name,
				gmail,
				age,
				password
			)VALUES(
				$1,
				$2,
				$3,
				$4,
				$5
			)
			`
			log.Println(" CreateUser")
	_, err := u.conn.Exec(
		ctx,
		query,
		user.User_id,
		user.User_name,
		user.Gmail,
		user.Age,
		user.Password,
	)
	log.Println(" CreateUsereeeeeeeeeeeee")

	if err!=nil {
		
		log.Println("error on CreateUser", err)

		return err
	}

	return nil
}

func(u *UserRepo) GetUsers(ctx context.Context, limit int, page int) (*[]modles.User, error){

	query:=`
		SELECT 
			*
		FROM
			Users
		LIMIT 
			$1
		OFFSET
			$2
		`
	
		var (
			user 	modles.User
			users	[]modles.User
		)

	rows, err := u.conn.Query(ctx,query,limit,page)

	if err!=nil {
		
		log.Println("error on get users", err)

		return nil,err
	}

	for rows.Next() {

		rows.Scan(
			&user.User_id,
			&user.User_name,
			&user.Gmail,
			&user.Age,
			&user.Password,
		)

		users = append(users, user)

	}

	

	return &users,nil
}

func(u *UserRepo) GetUserByName(ctx context.Context, user_name string) (*modles.User, error){

	query:=`
		SELECT 
			*
		FROM
			users
		WHERE 
			user_name=$1
		`
	var user modles.User

	err:=u.conn.QueryRow(
		ctx,
		query,
		user_name,
		).Scan(
		&user.User_id,
		&user.User_name,
		&user.Gmail,
		&user.Age,
		&user.Password,
	)
	
	if err!=nil {
		
		log.Println("error on GetUserByUserNAme", err)

		return nil,err
	}

	return &user,nil
}

func(u *UserRepo) UpdateUser(ctx context.Context, user modles.User,user_name string) error{

	query:=`
		UPDATE 
			users
		SET
			user_name=$1,
			gmail=$2,
			age=$3,
			password=$4
		WHERE 
			user_name=$5
			
	`
	_,err :=u.conn.Exec(
		ctx,
		query,
		user.User_name,
		user.Gmail,
		user.Age,
		user.Password,
		user_name,
	)

	if err!=nil {
		
		log.Println("error on UpdateUser inner", err)

		return err
	}

	return nil
}

func(u *UserRepo) DeleteUserByName(ctx context.Context, user_name string) error{

	query:=`
		DELETE  FROM
			users
		WHERE
			user_name=$1
	`

	_,err:=u.conn.Exec(
		ctx,
		query,
		user_name,
	)

	if err!=nil {
		
		log.Println("error on DeleteUserByUserName ", err)

		return err
	}
	
	return nil
}



