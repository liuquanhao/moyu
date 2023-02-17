package user_model

type User struct {
	Id    uint32 `json:"id"`
	User  string `json:"user" validate:"required,max=16"`
	Token string `json:"token" validate:"required,max=32"`
}
