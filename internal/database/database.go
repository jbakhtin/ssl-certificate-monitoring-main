package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

type ConnectionInfo struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func InitDatabaseClient(c *ConnectionInfo) error {
	dsn := c.User + ":" + c.Password + "@tcp(" + c.Host + ":" + c.Port + ")/" + c.Name

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("\nFail to connect the database.\nPlease make sure the connection info is valid %#v", c)
		return err
	}

	return nil
}