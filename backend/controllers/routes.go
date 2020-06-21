package controllers

import "github.com/akhamatvarokah/miniWallet/backend/middlewares"

func (s *Server) initializeRoutes() {

	r := s.Router.Group("/api")
	{
		// Login Route
		r.POST("/login", s.Login)
		r.POST("/users", s.CreateUser)
		r.GET("/users/others", middlewares.TokenAuthMiddleware(), s.OtherUser)

		r.GET("/balance/history", middlewares.TokenAuthMiddleware(), s.History)
		r.POST("/balance/topup", middlewares.TokenAuthMiddleware(), s.TopUp)
		r.POST("/balance/transfer", middlewares.TokenAuthMiddleware(), s.Transfer)
	}
}
