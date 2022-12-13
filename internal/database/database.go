package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

type ConnectionInfo struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func InitDatabaseClient(c *ConnectionInfo) error {
	dsn := c.Username + ":" + c.Password + "@tcp(" + c.Host + ":" + c.Port + ")/" + c.Database

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("\nFail to connect the database.\nPlease make sure the connection info is valid %#v", c)
		return err
	}

	return nil
}