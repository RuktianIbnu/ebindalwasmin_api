package controllers

import (
	"fmt"
	"net/http"
	"../models"
	//"../config"
	"encoding/json"
	//jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
)

var er error
var pasporData models.PasporData
var pasporDataDetail models.PasporDataDetail
var statusResPaspor models.StatusResPaspor

func CreateDataPaspor(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.FormValue("data"))
// read response data
data, _ := ioutil.ReadAll( r.Body )


// print request `Content-Type` header
// requestContentType := res.Request.Header.Get( "Content-Type" )
// fmt.Println( "Request content-type:", requestContentType )

// print response body
fmt.Printf( "%s\n", data )
	// fmt.Println(r.Body)
	// body := json.NewDecoder(r.Body)
	// err := body.Decode(&pasporData)
	// if err != nil {
    //     panic(err)
    // }
	// fmt.Println(err)
	statusRes.Status	= 200
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(statusRes.Status)

	//db := config.Connect()
	//var result = statusResPaspor
	//dataPaspor := json.NewDecoder(r.Body).Decode(&pasporData)
	
    // if r.Method == "POST" {

	// 	tokenString := r.Header.Get("Authorization")
	// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 		if jwt.GetSigningMethod("HS256") != token.Method {
	// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	// 		}

	// 		return []byte("secret"), nil
	// 	})
	// 	if token != nil && err == nil {
	// 		dataPaspor := json.NewDecoder(r.Body).Decode(&pasporData)
        
	// 		createData, err := db.Prepare("INSERT INTO data_paspor(id_jenis, id_user, id_kantor, tanggal, laki, perempuan, total, id_wilayah_kerja) VALUES(?,?,?,?,?,?,?,?)")
	// 		if err != nil {
	// 			statusResPaspor.Status	= 400
	// 			statusResPaspor.Msg		= "gagal tambah data"

	// 			result = statusResPaspor
	// 			panic(err.Error())
				
	// 		} else {
	// 			createData.Exec(dataPaspor)
	// 			statusRes.Status	= 200
	// 			statusRes.Msg		= "berhasil tambah data"
	// 			result = statusResPaspor
	// 		}
	// 	} else {
	// 		statusRes.Status	= 400
	// 			statusRes.Msg		= "not authorized"
	// 			statusRes.Token		= ""
	
	// 			result := statusRes
	// 		w.Header().Set("Content-Type", "application/json")
	// 		json.NewEncoder(w).Encode(result)
	// 	}
	// }
}
