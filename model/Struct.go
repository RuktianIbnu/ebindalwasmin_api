package model

// // User ...
// type User struct {
// 	ID         string `form:"id" json:"id"`
// 	Name       string
// 	Email      string
// 	IDKantor   string
// 	NamaKantor string
// }

// Auth ...
type Auth struct {
	Email    string
	Password string
}

// StatusRes ...
type StatusRes struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Token  string `json:"token"`
	Data   []User
}
