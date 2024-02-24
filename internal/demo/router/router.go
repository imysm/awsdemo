package router

import (
	"awsdemo/internal/demo/handler"
	"awsdemo/internal/demo/router/middleware"
	"awsdemo/internal/service/sd"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	u := g.Group("/awsdemo/test")
	u.Use(middleware.AuthMiddleware())
	{
		u.GET("/mysqlInsert", handler.Insert)
		u.GET("/redisInsert", handler.Cache)
	}
	// The health check handlers
	svcd := g.Group("/awsdemo/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
