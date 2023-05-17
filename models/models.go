package models

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}
type RegisterUser struct {
	Name     string
	Email    string
	Password string
}
type UserData struct {
	ID    int
	Name  string
	Email string
	Exp   int
}
