package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"github.com/aws/aws-lambda-go/lambda"
)

const DB_HOST = "cid-rep-dev-v2.cluster-cxkxcvxokdny.us-east-2.rds.amazonaws.com"
const DB_PORT = "3306"
const DB_USER = "calleridrepdev"
const DB_NAME = "calleridrep_db"
const DB_PASSWORD = "STAZP!^KUEx4jE4jpmUC2Exbp6"

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
	db, err := sql.Open("mysql", DB_USER + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME )

	if err != nil {
		log.Println(err)
	}

	database = db

	defer db.Close()

	lambda.Start(HandleLambdaEvent)
}
