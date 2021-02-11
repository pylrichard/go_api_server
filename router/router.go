package router

import (
	"net/http"

	"./middleware"
	"../handler/health"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 handler
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route")
	})
	// health check handlers
	hcHandlers := g.Group("/health") {
		hcHandlers.GET("/check", health.Check)
		hcHandlers.GET("/disk", health.DiskCheck)
		hcHandlers.GET("/cpu", health.CPUCheck)
		hcHandlers.GET("/mem", health.MemCheck)
	}

	return g
}