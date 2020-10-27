package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"ebindalwasmin_api/router"

	"github.com/joho/godotenv"
)

func init() {
	if strings.ToLower(os.Getenv("GIN_MODE")) != "release" {
		if err := godotenv.Load(); err != nil {
			log.Fatalln(err)
		}
	}
}

func main() {
	if err := router.Routes().Run(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		log.Fatalln(err)
	}
}
