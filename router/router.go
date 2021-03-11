package router

import (
	"net/http"

	_ "go/tiny_http_server/docs"
	"go/tiny_http_server/router/middleware"
	"go/tiny_http_server/handler/health"
	"go/tiny_http_server/handler/user"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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

	// swagger api docs
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// pprof router
	pprof.Register(g)

	g.POST("/login", user.Login)

	// User Handler
	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware()) {
		u.POST("", user.Create)
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
		u.GET("", user.List)
		u.GET("/:user_name", user.Get)
	}

	// Health check handlers
	h := g.Group("/health") {
		h.GET("/check", health.Check)
		h.GET("/disk", health.DiskCheck)
		h.GET("/cpu", health.CPUCheck)
		h.GET("/mem", health.MemCheck)
	}

	return g
}