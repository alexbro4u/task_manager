package main

import "task_manager/internal/server"

func main() {
	s := server.New()
	s.SetupRoutes()

	if err := s.Start("8080"); err != nil {
		// Обработка ошибки запуска сервера
	}
}
