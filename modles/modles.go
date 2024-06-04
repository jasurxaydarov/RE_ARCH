package modles

import "time"

type User struct{
	
	User_id			int
	User_name		string
	Gmail			string
	Age 			int
}
type Todo struct{

	Todo_id 		int
	Task			string
	Created_at		time.Time
	Is_completed	bool

}

type GetTodosResp struct{
	
	Todos 	[]Todo
	Count 	int
}