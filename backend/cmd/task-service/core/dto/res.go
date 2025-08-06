package dto

type Response struct {
	Status bool   `json:"status"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Data   any    `json:"data"`
}

type Task struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type Project struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	OwnerID     uint   `json:"owner_id"`
	StartedAt   string `json:"started_at"`
	FinishedAt  string `json:"finished_at"`
	DeadlineAt  string `json:"deadline_at"`
	Members     []uint `json:"members" gorm:"-"`
}

type UserInfo struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type ProjectResponse struct {
	ID          uint       `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	Owner       UserInfo   `json:"owner"`   // แสดงแทน owner_id
	Members     []UserInfo `json:"members"` // แสดงแทน member_ids
	StartedAt   string     `json:"started_at"`
	FinishedAt  string     `json:"finished_at"`
	DeadlineAt  string     `json:"deadline_at"`
}
