package models

type User struct {
	ID        string  `form:"id" json:"id"`
	Name  	  string
	Email     string
	Id_kantor string
	Nama_kantor string
}

type Auth struct {
	Email     string
	Password  string
}

type StatusRes struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data []User
}