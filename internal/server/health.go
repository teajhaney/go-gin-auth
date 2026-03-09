package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"time"
)


func health (c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"ok": true,
			"status": "go-auth is Healthy",
			"time": time.Now().UTC(),
		})
	}





