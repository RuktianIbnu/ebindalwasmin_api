package controllers

import (
	//"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"../models"
	"../config"

	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	sessions "github.com/kataras/go-sessions"
	"golang.org/x/crypto/bcrypt"
	// "os"
)

var err error
var Users models.User
var Auths models.Auth
var statusRes models.StatusRes

func Login(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	var arr_user []models.User
	session := sessions.Start(w, r)
	if len(session.GetString("username")) != 0 && checkErr(w, r, err) {
		//check session if avaliabel
		statusRes.Status	= 200
		statusRes.Msg		= "berhasil login session avliabe"
		result := statusRes
		json.NewEncoder(w).Encode(result)
	}
	if r.Method != "POST" {
		statusRes.Status	= 400
		statusRes.Msg		= "Method Must be post"
		result := statusRes
		json.NewEncoder(w).Encode(result)
	}
	email := r.FormValue("email")
	password := r.FormValue("password")

	Users := QueryUser(email)
	

	//deskripsi dan compare password
	var password_tes = bcrypt.CompareHashAndPassword([]byte(Auths.Password), []byte(password))
	
	if password_tes == nil {
		//login success
		arr_user = append(arr_user, Users)
		session := sessions.Start(w, r)
		session.Set("email", Users.Email)
		session.Set("name", Users.Name)

		//key := r.Header.Get("key")
		sign := jwt.New(jwt.GetSigningMethod("HS256"))
		token, err := sign.SignedString([]byte("secret"))

		createData, er := db.Prepare("UPDATE users SET remember_token=? WHERE email=?")
			if err != nil {
				statusRes.Status	= 400
				statusRes.Msg		= "gagal generate token"
				statusRes.Token		= token
				result := statusRes
				
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(result)
			} else {
				if er != nil {
					statusRes.Status	= 400
					statusRes.Msg		= "gagal update token"
					statusRes.Token		= ""
					result := statusRes
					
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(result)
				}else{
					createData.Exec(token,email)
					statusRes.Status	= 200
					statusRes.Msg		= "berhasil login"
					statusRes.Token		= token
					statusRes.Data 		=  arr_user
					
					result := statusRes
					
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(result)
				}
			}
	} else {
		//login failed
		statusRes.Status	= 400
		statusRes.Msg		= "gagal login"
		statusRes.Data		= arr_user

		result := statusRes
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func QueryUser(email string) models.User {
	var arr_user []models.User
	db := config.Connect()

	rows, err := db.Query(`
		SELECT a.id, a.email, 
		a.name, a.password, a.id_kantor, b.nama_kantor
		FROM users as a
		left join kantor as b on a.id_kantor = b.id_kantor WHERE email=?`, email)
		if err != nil {
			log.Print(err)
		}
		for rows.Next() {
			if err := rows.Scan(&Users.ID, &Users.Email, &Users.Name, &Auths.Password, &Users.Id_kantor, &Users.Nama_kantor); err != nil {
				log.Fatal(err.Error())
			} else {
				arr_user = append(arr_user, Users)
			}
		}
		
	return Users
}

func checkErr(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {

		fmt.Println(r.Host + r.URL.Path)

		http.Redirect(w, r, r.Host+r.URL.Path, 301)
		return false
	}

	return true
}