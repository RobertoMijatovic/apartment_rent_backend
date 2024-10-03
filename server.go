package main

import (
	"apartmentRent/helpers"
	"apartmentRent/routes"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load("apartmentRent.env")
	if err != nil {
		panic(helpers.ErrEnvLoad)
	}

	connString := os.Getenv("CON_STRING")
	if connString == "" {
		panic(helpers.ErrEnvLoad)
	}

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		panic(helpers.ErrDbConn)
	}

	routes := routes.InitRoutes(db)

	routes.Run("127.0.0.1:1312")
}
