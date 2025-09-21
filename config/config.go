package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct{
	Version string
	ServiceName string
	HttpPort int
	JwtSecretKey string
}

var configuration *Config


func loadCOnfig(){
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

	configuration = &Config{
		Version: version,
		ServiceName: serviceName,
		HttpPort: port,
		JwtSecretKey: jwtSercetKey,
	}
}

func GetConfig() *Config{
	if configuration == nil{
		loadCOnfig()
	}
	return configuration
}