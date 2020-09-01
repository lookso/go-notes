package resource

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UserInfoResponse struct {
	Id   string  `json:"id"`
	Name string `json:"name"`
	Sex  int    `json:"sex"`
	Age  int    `json:"age"`
}
