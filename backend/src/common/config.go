package common

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	DB                      *DBConfig
	AUTH                    *AuthConfig
	MaxMultipartMemory      int64
	SupportedFileExtensions []string
	Host                    string
	Port                    string
}

type DBConfig struct {
	ConnectionString string
	AutoMigrate      bool
}

type AuthConfig struct {
	HmacSecret        []byte
	ExpirationMinutes uint
}

var config *Config

func InitConfig() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: error loading .env file")
		log.Println("Will use environment variables")
	}

	config = &Config{
		DB:   &DBConfig{},
		AUTH: &AuthConfig{},
	}

	config.DB.ConnectionString = os.Getenv("DB_CONNECTION_STRING")
	config.DB.AutoMigrate = strings.ToLower(os.Getenv("DB_AUTO_MIGRATE")) == "true"
	config.AUTH.HmacSecret = []byte(os.Getenv("AUTH_HMAC_SECRET"))

	expirationMinutes, err := strconv.Atoi(os.Getenv("AUTH_EXPIRATION_MINUTES"))
	if err != nil {
		log.Fatal("AUTH_EXPIRATION_MINUTES is invalid")
	}
	config.AUTH.ExpirationMinutes = uint(expirationMinutes)

	config.MaxMultipartMemory, err = strconv.ParseInt(os.Getenv("MAX_MULTIPART_MEMORY"), 10, 64)
	if err != nil {
		log.Fatal("MAX_FILE_SIZE is invalid")
	}

	config.SupportedFileExtensions = strings.Split(os.Getenv("SUPPORTED_FILE_EXTENSIONS"), "|")

	config.Host = os.Getenv("HOST")
	config.Port = os.Getenv("PORT")
}

func GetConfig() *Config {
	return config
}
