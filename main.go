package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/QingGo/ddnet_2022_chn_stat_backend/db"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	csvFilePath := os.Getenv("CSV_FILE_PATH")
	if csvFilePath == "" {
		csvFilePath = "../results_example.csv"
	}
	isProduct := os.Getenv("DEBUG")
	db.Init(csvFilePath)
	router := gin.Default()
	if isProduct == "" {
		config := cors.DefaultConfig()
		config.AllowOrigins = []string{"*"}
		router.Use(cors.New(config))
	}
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, db.Find(name))
	})
	router.Run()
}
