package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/qdriven/qfluent-go/internal/log"
	ginRequestId "github.com/qdriven/qfluent-go/pkg/qgin/middleware/request"
)

var logger = log.ZoneLogger("qgin/router")

// NewRouter creates a new router (a gin.New() router)
// with gin.Recovery() middleware, the log.Logger4Gin middleware,
// the ginRequestId.RequestID() middleware,
// and addon middlewares indicated by the options parameters.
func NewRouter(options ...RouterOption) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery(), log.Logger4Gin, ginRequestId.RequestID())

	for _, option := range options {
		router = option(router).(*gin.Engine)
	}

	return router
}

// RouterOption is an option to construct the router.
type RouterOption func(router gin.IRouter) gin.IRouter

// AllowAllCors allows all origins, methods and headers.
// It's useful for dev. And it is not recommended for the production.
func AllowAllCors() RouterOption {
	return func(router gin.IRouter) gin.IRouter {
		logger.Warn("AllowAllCors: Cors is enabled, this is a security risk!")
		router.Use(cors.Default())
		return router
	}
}

// WithRequestID adds the ginRequestId.RequestID() middleware,
// which adds a request_id in the context for each request.
// And the request_id will be writen to the X-Request-Id response header.
func WithRequestID() RouterOption {
	return func(router gin.IRouter) gin.IRouter {
		router.Use(ginRequestId.RequestID())
		return router
	}
}

// WithMiddleware adds custom middlewares to the router.
func WithMiddleware(middleware ...gin.HandlerFunc) RouterOption {
	return func(router gin.IRouter) gin.IRouter {
		router.Use(middleware...)
		return router
	}
}
