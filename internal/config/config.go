package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-default:"local"`
	StoragePath string `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`
	HTTPServer  `yaml:"http_server" env:"HTTP_SERVER"`
	DB          DB `yaml:"db" env:"DB"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env:"HTTP_ADDR" env-default:"127.0.0.1:8080"`
	Timeout     time.Duration `yaml:"timeout" env:"HTTP_TIMEOUT" env-default:"50s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"IDLE_TIMEOUT" env-default:"600s"`
}

type DB struct {
	Host     string        `yaml:"host" env:"DB_HOST" env-default:"db"`
	Port     string        `yaml:"port" env:"DB_PORT" env-default:"5432"`
	Username string        `yaml:"username" env:"DB_USERNAME" env-default:"postgres"`
	Password string        `yaml:"password" env:"DB_PASSWORD" env-default:"qwerty"`
	DBName   string        `yaml:"db_name" env:"DB_NAME" env-default:"postgres"`
	SSLMode  string        `yaml:"sslmode" env:"DB_SSL_MODE" env-default:"disable"`
	Timeout  time.Duration `yaml:"timeout" env:"DB_TIMEOUT" env-default:"20s"`
	Driver   string        `yaml:"driver" env:"DB_DRIVER" env-default:"postgres"`
}

func MustLoad() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	configPath, _ := os.LookupEnv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("Путь до конфига не найден в енв файле")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("Файл конфига не найден")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Не можем прочитать конфиг %s", err)
	}
	log.Printf("Config: %+v", cfg)
	return &cfg
}
