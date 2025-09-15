package main

import (
	"fmt"
	"log"
	"microservice-architecture-go/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {

	// Загрузка конфигураций
	cfg := config.MustLoad()

	var logger *zap.Logger

	// Выстраивание среды
	if cfg.Debug {
		logger, _ = zap.NewDevelopment()
		gin.SetMode(gin.DebugMode)
	} else {
		logger, _ = zap.NewProduction()
		gin.SetMode(gin.ReleaseMode)
	}

	logger.Info("Logger initialization is successfully")

	// Запись логов из буфера, если программа завершится
	defer logger.Sync()
	router := gin.New()

	// Формирования адреса, где будет поднято приложение
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("error starting server with address: %s", addr)
	}

}
