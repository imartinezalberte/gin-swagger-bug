package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imartinezalberte/hello-swagger/docs"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/swaggo/swag"
)

// config options for display: https://swagger.io/docs/open-source-tools/swagger-ui/usage/configuration/#display
var swagUiConfig = httpSwagger.UIConfig(map[string]string{
	"displayOperationId":    "true",
	"defaultModelRendering": `"model"`,
	"filter":                "true",
	"showExtensions":        "true",
	"syntaxHighlight":       `{"active":"true","theme":"obsidian"}`,
	"tagSorter":             `"alpha"`,
})

func main() {
	engine := gin.Default()

	g := engine.Group("/api/v1")
	{
		g.GET("/ping", ping)
	}

	engine.GET("/greeting", greeting)

	engine.GET("/swagger/health/*any", func(c *gin.Context) { swapSpecHandler(c, docs.SwaggerInfohealth) })
	engine.GET("/swagger/greeting/*any", func(c *gin.Context) { swapSpecHandler(c, docs.SwaggerInfogreeting) })

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

func swapSpecHandler(c *gin.Context, spec *swag.Spec) {
	name := spec.InstanceName()
	gin.WrapH(httpSwagger.Handler(
		httpSwagger.InstanceName(name),
		httpSwagger.URL(fmt.Sprintf("/swagger/%s/doc.json", name)),
		httpSwagger.DocExpansion("none"),
		swagUiConfig,
	))(c)
}
