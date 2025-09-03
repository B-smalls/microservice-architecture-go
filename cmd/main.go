package main

import (
	"microservice-architecture-go/internal/config"
)

func main() {

	// Загрузка конфигураций
	cfg := config.MustLoad()

	_ = cfg
}
