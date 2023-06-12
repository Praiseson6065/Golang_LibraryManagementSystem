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
	BookId       int    `json:"Id"`
	BookName     string `json:"BookName"`
	ISBN         string `json:"ISBN"`
	Pages        int    `json:"Pages"`
	Publisher    string `json:"Publisher"`
	Author       string `json:"Author"`
	Taglines     string `json:"Taglines"`
	InsertedTime string `json :"InsertedTime"`
	Votes        int    `json:"votes"`
}
type SearchBook struct {
	SearchValue  string `json:"SearchValue"`
	SearchColumn string `json:"SearchColumn"`
}
