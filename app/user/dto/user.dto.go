package user

type CreateUserDTO struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
