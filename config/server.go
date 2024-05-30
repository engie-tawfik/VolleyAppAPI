package config

import (
	"github.com/gin-gonic/gin"
)

var Server *gin.Engine

func InitServer() {
	Server = gin.Default()
	Server.Use(gin.Recovery())
	go Server.Run("localhost:" + Port)
}
