package userclient

type UserInfo struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type GetUsersResponse struct {
	Status bool       `json:"status"`
	Code   int        `json:"code"`
	Msg    string     `json:"msg"`
	Data   []UserInfo `json:"data"`
}
