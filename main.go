package main

import (
	"database/sql"
	"go-clean-code/controller"
	"go-clean-code/handlers/command"
	"go-clean-code/handlers/query"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// Connect To MySql
	dbDriver := os.Getenv("DB_DRIVER")
	dbAddress := os.Getenv("DB_ADDRESS")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USER_NAME")
	dbPass := os.Getenv("DB_PASSWORD")

	dbase, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbAddress+")"+"/"+dbDatabase+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	defer func() {
		if err := dbase.Close(); err != nil {
			panic(err.Error())
		}
	}()

	r := gin.New()
	r.Use(gin.Logger())

	api := r.Group("/api/v1")

	command, err := command.NewUserCommand(dbase)
	if err != nil {
		panic(err.Error())
	}

	query, err := query.NewUserQuery(dbase)
	if err != nil {
		panic(err.Error())
	}

	handler := controller.NewUserController(command, query)
	handler.MapUserEndpoint(api)

	r.Run(":8080")
}
