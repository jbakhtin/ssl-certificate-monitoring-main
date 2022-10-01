package main

import (
	"database/sql"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

var db_host = os.Getenv("DB_HOST")
var db_port = os.Getenv("DB_PORT")
var db_user = os.Getenv("DB_USER")
var db_name = os.Getenv("DB_NAME")
var db_password = os.Getenv("DB_PASSWORD")

var database *sql.DB

type MyEvent struct {
	Name string `json:"What is your name?"`
	Age int     `json:"How old are you?"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

type ServerRegion struct{
	Id int
	Host string
}

func request(url string) {
	_, err := http.Get(url)

	if err == nil {
		log.Println(err)
	}
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	rows, err := database.Query("select id, host from stir_shaken_region_servers where enabled = 1")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		p := ServerRegion{}
		err := rows.Scan(&p.Id, &p.Host)
		if err != nil {
			fmt.Println(err)
			continue
		}

		go request(p.Host)
	}

	return MyResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}

func main() {
	db, err := sql.Open("mysql", db_user + ":" + db_password + "@tcp(" + db_host + ":" + db_port + ")/" + db_name )

	if err != nil {
		log.Println(err)
	} else {
		log.Println("DB connection successful")
	}

	database = db

	defer db.Close()

	lambda.Start(HandleLambdaEvent)
}
