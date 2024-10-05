package main

import (
	"task-master/config"
	"task-master/internal/app"
	logger "task-master/pkg/logs"
)

func main() {
	// Memuat konfigurasi
	config.Load()

	// Menginisialisasi logging
	log := logger.NewLogger()

	// Menginisialisasi aplikasi Fiber
	fiberApp := app.NewApp(log)

	// Memulai server
	log.Info("Starting server on port " + config.AppConfig.Port)
	if err := fiberApp.Listen(":" + config.AppConfig.Port); err != nil {
		log.Fatal(err.Error())
	}
}
