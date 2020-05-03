package config

import (
	"os"
	"strconv"
	"time"
)

type App struct {
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string
	LogSavePath     string
	LogSaveName     string
	LogFileExt      string
	TimeFormat      string
}

type Server struct {
	RunMode      string
	HTTPPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Database struct {
	Host     string
	Port     string
	Type     string
	Name     string
	User     string
	Password string
}

type Redis struct {
	Host        string
	Port        string
	Password    string
	SSL         bool
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type RabbitMQ struct {
	Host     string
	Port     string
	User     string
	Password string
}

var (
	AppSetting      *App
	ServerSetting   *Server
	DatabaseSetting *Database
	RedisSetting    *Redis
	RabbitMQSetting *RabbitMQ
)

func LoadAppProperties() {
	AppSetting = &App{
		PageSize:        atoi(os.Getenv("PAGE_SIZE")),
		JwtSecret:       os.Getenv("JWT_SECRET"),
		RuntimeRootPath: os.Getenv("RUNTIME_ROOT_PATH"),
		LogSavePath:     os.Getenv("LOG_SAVE_PATH"),
		LogSaveName:     os.Getenv("LOG_SAVE_NAME"),
		LogFileExt:      os.Getenv("LOG_FILE_EXT"),
		TimeFormat:      os.Getenv("TIME_FORMAT"),
	}

	ServerSetting = &Server{
		RunMode:      os.Getenv("RUN_MODE"),
		HTTPPort:     os.Getenv("HTTP_PORT"),
		ReadTimeout:  time.Duration(parseInt(os.Getenv("READ_TIMEOUT"))),
		WriteTimeout: time.Duration(parseInt(os.Getenv("WRITE_TIMEOUT"))),
	}

	DatabaseSetting = &Database{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Type:     os.Getenv("DB_TYPE"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	RedisSetting = &Redis{
		Host:        os.Getenv("REDIS_HOST"),
		Port:        os.Getenv("REDIS_PORT"),
		Password:    os.Getenv("REDIS_PASSWORD"),
		SSL:         parseBool(os.Getenv("REDIS_SSL")),
		MaxIdle:     atoi(os.Getenv("REDIS_MAX_IDLE")),
		MaxActive:   atoi(os.Getenv("REDIS_MAX_ACTIVE")),
		IdleTimeout: time.Duration(parseInt(os.Getenv("REDIS_TIMEOUT"))),
	}

	RabbitMQSetting = &RabbitMQ{
		Host:     os.Getenv("RABBITMQ_HOST"),
		Port:     os.Getenv("RABBITMQ_PORT"),
		User:     os.Getenv("RABBITMQ_USER"),
		Password: os.Getenv("RABBITMQ_PASSWORD"),
	}
}

func atoi(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return value
}

func parseInt(s string) int64 {
	value, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return value
}

func parseBool(s string) bool {
	value, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return value
}
