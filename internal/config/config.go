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

func MustLoad() *Config {

	// Получение пути до файла конфигурации
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATN enviroment variable is not set")
	}

	// Проверка на существования файла конфигурации
	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("error opening config file: %s", err)
	}

	var cfg Config

	// Чтение файла конфигурации
	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		log.Fatalf("error reading config file: %s", err)
	}

	return &cfg
}
