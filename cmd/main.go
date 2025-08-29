package main

import (
	"github.com/MoodyShoo/GinAPI/internal/database"
	http "github.com/MoodyShoo/GinAPI/internal/http"
	"github.com/MoodyShoo/GinAPI/internal/logger"
	"github.com/MoodyShoo/GinAPI/internal/service"
)

func main() {
	log := logger.InitLogger()

	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Ошибка инициализации БД: %v", err)
	}
	defer db.Close()

	userService := service.NewUserService(db.US())

	r := http.SetupRouter(userService, log)

	log.Info("Сервер запущен на :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
