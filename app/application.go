package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartServer() {
	mapUrl()
	err := router.Run(":9000")
	if err != nil {
		return
	}
}
