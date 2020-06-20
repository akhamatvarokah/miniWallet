package controllers

func (s *Server) initializeRoutes() {

	r := s.Router.Group("/api")
	{
		// Login Route
		r.POST("/login", s.Login)
		r.POST("/users", s.CreateUser)
	}
}
