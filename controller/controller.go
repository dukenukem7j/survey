package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes : Registers all rooutes available
func RegisterRoutes(r *gin.Engine) {

	r.LoadHTMLGlob("./templates/**/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/employees/:id", employeeHandler)

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// the jwt middleware
	authMiddleware := authMiddleware()

	r.POST("/login", authMiddleware.LoginHandler)

	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	}

	admin := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))

	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin-overview.html", nil)
	})

	admin.POST("/employees/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "add" {

		}
	})

	r.Static("/public", "./public")

}
