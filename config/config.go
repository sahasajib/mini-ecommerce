package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct{
	Host string
	Port int
	Name string
	User string
	Password string
	EnableSSLMode bool
}

type Config struct{
	Version string
	ServiceName string
	HttpPort int
	JwtSecretKey string
	DB *DBConfig
}



var configuration *Config


func loadConfig(){
	err := godotenv.Load()
	if err != nil{
		fmt.Printf("Failed to load the env variables: %v", err)
		os.Exit(1)
	}
	version := os.Getenv("VERSION")
	if version == ""{
		fmt.Println("Version is required")
		os.Exit(1)
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == ""{
		fmt.Println("Service name is required")
		os.Exit(1)
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == ""{
		fmt.Println("Service port is required")
		os.Exit(1)
	}

	port, err := strconv.Atoi(httpPort)
	if err != nil{
		fmt.Println("port must be number")
		os.Exit(1)
	}
	jwtSercetKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSercetKey == ""{
		fmt.Println("JWT secret key is required")
		os.Exit(1)
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == ""{
		fmt.Println("HOST is required")
		os.Exit(1)
	}
	dbPrt := os.Getenv("DB_PORT")
	if dbPrt == ""{
		fmt.Println("PORT is required")
		os.Exit(1)
	}
	dbPort, err := strconv.Atoi(dbPrt)
	if err != nil{
		fmt.Println("port must be number")
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == ""{
		fmt.Println("NAME is required")
		os.Exit(1)
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == ""{
		fmt.Println("USER is required")
		os.Exit(1)
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == ""{
		fmt.Println("PASSWORD is required")
		os.Exit(1)
	}
	sslMode := os.Getenv("ENABLE_SSL_MODE")
	enableSSLMode, err := strconv.ParseBool(sslMode)
	if err != nil{
		fmt.Println("ENABLE_SSL_MODE must be boolean")
		os.Exit(1)
	}
	
	dbConfig := &DBConfig{
		Host: dbHost,
		Port: dbPort,
		Name: dbName,
		User: dbUser,
		Password: dbPassword,
		EnableSSLMode: enableSSLMode,
	}

	configuration = &Config{
		Version: version,
		ServiceName: serviceName,
		HttpPort: port,
		JwtSecretKey: jwtSercetKey,
		DB: dbConfig,
	}
}

func GetConfig() *Config{
	if configuration == nil{
		loadConfig()
	}
	return configuration
}