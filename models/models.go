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
	Usertype string
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
type Logdb struct {
	LogId      int
	UserId     int
	UserType   string
	Operation  string
	InsertTime string
	UserName   string
}

type Book struct {
	BookId    int
	BookName  string
	ISBN      string
	Pages     int
	Publisher string
	Author    string
	Quantity  int
	Taglines  []string
}
