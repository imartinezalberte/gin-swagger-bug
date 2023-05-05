package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/imartinezalberte/hello-swagger/docs"
	swaggerfiles "github.com/swaggo/files"
  ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	engine := gin.Default()

	g := engine.Group("/api/v1")
	{
		g.GET("/ping", ping)
	}

	engine.GET("/greeting", greeting)

	engine.GET("/swagger/health/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.InstanceName("health")))
	engine.GET("/swagger/greeting/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.InstanceName("greeting")))

	engine.Run(":8080")
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags health
// @Produce json
// @Success 200 {json} {}
// @Router /ping [get]
func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}


// @BasePath /

// GreetingExample godoc
// @Summary greeting example
// @Schemes
// @Description do greeting
// @Tags greeting
// @Produce json
// @Success 200 {json} {}
// @Router /greeting [get]
func greeting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello world!",
	})
}
