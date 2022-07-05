package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthChecker struct{}

func MapHealthCheckRoutes(r *gin.Engine) {
	r.GET("/ping", ping)
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
