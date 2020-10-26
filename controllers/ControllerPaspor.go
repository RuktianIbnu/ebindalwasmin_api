package controllers

import (
	"ebindalwasmin_api/models"
	"fmt"
	"net/http"

	"encoding/json"

	jwt "github.com/dgrijalva/jwt-go"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"

	//"io/ioutil"
	"log"
)

var er error
var pasporData models.PasporData
var pasporDataDetail models.PasporDataDetail
var statusResPaspor models.StatusResPaspor
var result = statusResPaspor

// CreateDataPaspor ...
func CreateDataPaspor(w http.ResponseWriter, r *http.Request) {

	//db := config.Connect()

	if r.Method == "POST" {

		tokenString := r.Header.Get("Authorization")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte("secret"), nil
		})
		if token != nil && err == nil {
			var pasporDatas []models.PasporData
			err := json.NewDecoder(r.Body).Decode(&pasporDatas)
			log.Println(pasporDatas)
			if err != nil {
				statusResPaspor.Status = 500
				statusResPaspor.Msg = "data kosong"

				result = statusResPaspor
			} else {
				//createData, err := db.Prepare("INSERT INTO data_paspor(id_jenis, id_user, id_kantor, tanggal, laki, perempuan, total, id_wilayah_kerja) VALUES(?,?,?,?,?,?,?,?)")
				if err != nil {
					statusResPaspor.Status = 400
					statusResPaspor.Msg = "gagal tambah data"

					result = statusResPaspor
					panic(err.Error())

				} else {
					statusRes.Status = 200
					statusRes.Msg = "berhasil tambah data"
					result = statusResPaspor
				}
				w.Header().Set("Content-Type", "application/json")
				//w.Write(b)
			}
		} else {
			statusRes.Status = 400
			statusRes.Msg = "not authorized"
			statusRes.Token = ""

			result = statusResPaspor
		}

	}
}
