package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"ssl-monitor-main/internal/configs"
	"ssl-monitor-main/internal/database"
	wakeUpRegionServers "ssl-monitor-main/internal/servises"
)

var db *gorm.DB

func main() {
	if err := configs.Init(); err != nil {
		log.Panicln(err)
	}

	err := database.InitDatabaseClient(&database.ConnectionInfo{
		Username:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database:     os.Getenv("DB_DATABASE"),
	})

	if err != nil {
		log.Panicln(err)
	}

	ctx := context.TODO()
	lambda.StartWithOptions(wakeUpRegionServers.Run, lambda.WithContext(ctx))
}