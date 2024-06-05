package modles

import (
	"time"

	"github.com/google/uuid"
)

type User struct{
	
	User_id			uuid.UUID
	User_name		string
	Gmail			string
	Age 			int
	Password        string
}
type Todo struct{

	Todo_id 		uuid.UUID
	Task			string
	Created_at		time.Time
	Is_completed	bool
	User_name		string

}

type GetTodosResp struct{
	
	Todos 	[]Todo
	Count 	int
}

type Userr struct{
	
	User_id			int
	User_name		string
	Gmail			string
	Age 			int
	Password        string
}