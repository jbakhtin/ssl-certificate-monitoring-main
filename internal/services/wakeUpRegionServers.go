package wakeUpRegionServers

import (
	"context"
	"gorm.io/gorm"
	"log"
	"net/http"
	"ssl-monitor-main/internal/database"
	"ssl-monitor-main/internal/models"
	"sync"
)

type Event struct {
	DB *gorm.DB
}

type Response struct {
	Success bool
	Message string
}

var wait_rutine sync.WaitGroup


func Run(c context.Context) (Response, error) {
	var regionServer models.StirShakenRegionServer
	rows, err := database.DB.Model(&regionServer).Where("enabled = ?", 1).Rows()

	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := database.DB.ScanRows(rows, &regionServer)

		if err != nil {
			log.Println(err)
			continue
		}

		wait_rutine.Add(1)
		go request(regionServer.Host)
	}

	wait_rutine.Wait()

	return Response{Success: true}, nil
}

func request(url string) {
	_, err := http.Get(url)

	if err == nil {
		log.Println(err)
	}

	wait_rutine.Done()
}
