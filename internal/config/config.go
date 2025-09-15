package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// Структура конфигураций сервера
type Config struct {
	Host  string `env:"HOST"`
	Port  int    `env:"PORT"`
	Debug bool   `env:"DEBUG"`
}

// Метод инициализации конфигураций
func MustLoad() *Config {

	// Путь до файла конфигурации
	configPath := "./config/.env"

	// Проверка на существования файла конфигурации
	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("error opening config file: %s", err)
	}

	// Конфигурация
	var cfg Config

	// Чтение файла конфигурации
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("error reading config file: %s", err)
	}

	return &cfg
}
