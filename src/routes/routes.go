package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, services map[string]string) {
	router.GET("/auth/*proxyPath", func(c *gin.Context) {
		proxyRequest(c, services["auth"])
	})
	router.GET("/items/*proxyPath", func(c *gin.Context) {
		proxyRequest(c, services["items"])
	})
	router.GET("/transactions/*proxyPath", func(c *gin.Context) {
		proxyRequest(c, services["transactions"])
	})
	// Add additional routes as needed
}

func proxyRequest(c *gin.Context, targetService string) {
	proxyPath := c.Param("proxyPath")
	targetURL := targetService + proxyPath

	resp, err := http.Get(targetURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Service unavailable"})
		return
	}
	defer resp.Body.Close()

	c.Status(resp.StatusCode)
	c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
}
