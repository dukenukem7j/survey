package main

import (
	"github.com/dukenukem7j/survey/controller"

	"github.com/gin-gonic/gin"
)

const port = "8080"

func main() {

	r := gin.Default()

	//gin.SetMode(gin.ReleaseMode)

	controller.RegisterRoutes(r)

	r.Run(":" + port)
}
