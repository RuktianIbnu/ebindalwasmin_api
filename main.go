package main

import (
	//"database/sql"
	"encoding/json"
	"fmt"
	//"log"
	"net/http"
	"./models"
	"./config"
	"./controllers"

	_ "github.com/go-sql-driver/mysql"
	sessions "github.com/kataras/go-sessions"
	//"golang.org/x/crypto/bcrypt"
	// "os"
)

func main() {
	db := config.Connect()
	
	// config.Connect()
	// routes.NewRouter()

	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/logout", logout)
	defer db.Close()
	

	fmt.Println("Server running on port :8001")
	http.ListenAndServe(":8001", nil)
}

func logout(w http.ResponseWriter, r *http.Request) {
	var statusRes models.StatusRes
	session := sessions.Start(w, r)
	session.Clear()
	sessions.Destroy(w, r)

	statusRes.Status = 200
	statusRes.Msg = "berhasil logout"
	res := statusRes

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}