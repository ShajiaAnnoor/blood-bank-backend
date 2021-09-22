package server

import (
	"gitlab.com/Aubichol/hrishi-backend/api/middleware"
	"gitlab.com/Aubichol/hrishi-backend/container"
)

//Middleware registers middleware related providers
func Middleware(c container.Container) {
	c.Register(middleware.NewAuthMiddleware)
	c.Register(middleware.NewAuthMiddlewareURL) // don't know what it does
	c.Register(middleware.MessageNotificationMiddleware)
}
