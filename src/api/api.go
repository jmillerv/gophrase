package api

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Serve() {
	router := gin.Default()

	router.POST("/entropy", postEntropy)
	router.GET("options", getOptions)
	router.GET("/passphrase", getPassphrase)

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}