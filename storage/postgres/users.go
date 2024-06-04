package postgres

import (
	"arch/modles"
	"arch/repoi"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type UserRepo struct {

	conn pgx.Conn
}

func NewUserRepo(conn pgx.Conn)repoi.UserRepoI{

	return &UserRepo{
			conn: conn,
	}


}

func(u *UserRepo) CreateUser(ctx context.Context, user modles.User) error{

	query:=
		`INSERT INTO
			users(
				user_name,
				gmail,
				age
			)VALUES(
				$1,
				$2,
				$3
			)
			`
	_, err := u.conn.Exec(
		ctx,
		query,
		user.User_name,
		user.Gmail,
		user.Age,
	)

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
			users
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

		rows.Scan(&user.User_id,&user.User_name,&user.Gmail,&user.Age)

		users = append(users, user)

	}

	

	return &users,nil
}

func(u *UserRepo) GetUserById(ctx context.Context, user_id int) (*modles.User, error){

	query:=`
		SELECT 
			*
		FROM
			users
		WHERE 
			user_id=$1
		`
	var user modles.User

	err:=u.conn.QueryRow(ctx,query,user_id).Scan(
		&user.User_id,
		&user.User_name,
		&user.Gmail,
		&user.Age,
	)
	
	if err!=nil {
		
		log.Println("error on GetUserById", err)

		return nil,err
	}

	return &user,nil
}

func(u *UserRepo) UpdateUser(ctx context.Context, user modles.User) error{

	query:=`
		UPDATE 
			users
		SET
			user_name=$1,
			gmail=$2,
			age=$3
	`
	_,err :=u.conn.Exec(ctx,query,user.User_name,user.Gmail,user.Age)

	if err!=nil {
		
		log.Println("error on UpdateUser", err)

		return err
	}

	return nil
}

func(u *UserRepo) DeleteUserById(ctx context.Context, user_id int) error{

	query:=`
		DELETE 
			users
		WHERE
			user_id=$1
	`

	_,err:=u.conn.Exec(ctx,query,user_id)

	if err!=nil {
		
		log.Println("error on DeleteUserById ", err)

		return err
	}
	
	return nil
}