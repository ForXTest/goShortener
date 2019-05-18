package main

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"urlShotener/cmd/database"
	"urlShotener/cmd/handlers"
)

func main() {

	configData, err := GetConfig()

	if err != nil {
		log.Fatalf("Failed to get config: %v\n", err)
	}

	log.Printf("Config %+v", configData)

	MySqlConnect, err := database.NewMySQL(configData.MySQL)
	if err != nil {
		log.Fatalf("Failed connect to mysql: %v\n", err)
	}

	urlHandlers := &handlers.ShortUrlHandlers{MySqlConnect}

	gin.SetMode(configData.GinMode.Mode)

	router := gin.Default()
	router.POST("/short/", urlHandlers.SaveShortUrl)
	router.GET("/:short", urlHandlers.GetFullUrl)

	router.Run(fmt.Sprintf(":%d", configData.Server.Port));
}
