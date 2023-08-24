package main

import (
	"github.com/gin-gonic/gin"

	"emailchaser.com/backend-go/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	app := gin.Default()
	app.Run()
}
