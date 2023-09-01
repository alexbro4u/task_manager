package server

// SetupRoutes настраивает маршруты сервера
func (s *Server) SetupRoutes() {
	r := s.router
	r.POST("/create-task", s.CreateTaskHandler)
	r.GET("/list-tasks", s.ListTasksHandler)
}
