package user_model

type User struct {
	Id       uint32 `json:"id"`
	Username string `json:"username" validate:"required,max=16"`
	Password string `json:"password" validate:"required,max=32"`
}
