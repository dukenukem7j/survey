package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func employeeHandler(c *gin.Context) {
	id := c.Param("id")

	c.HTML(http.StatusOK, "employee.html", map[string]interface{}{
		"Id": id,
	})
}
