package utils

import (
	"github.com/go-kit/kit/log"
	"github.com/lib/pq"
	"github.com/spf13/viper"
)


// Configurations wraps all the config variables required by the auth service
type Configurations struct {
	ServerAddress              string
	DBHost                     string
	DBName                     string
	DBUser                     string
	DBPass                     string
	DBPort                     string
	DBConn                     string
}

// NewConfigurations returns a new Configuration object
func NewConfigurations(logger log.Logger) *Configurations {

	viper.AutomaticEnv()

	dbURL := viper.GetString("DATABASE_URL")
	conn, _ := pq.ParseURL(dbURL)
	logger.Log("found database url in env, connection string is formed by parsing it")
	logger.Log("db connection string", conn)

	viper.SetDefault("SERVER_ADDRESS", "0.0.0.0:9090")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_NAME", "postgres")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_PASSWORD", "root")
	viper.SetDefault("DB_PORT", "5432")

	configs := &Configurations{
		ServerAddress: viper.GetString("SERVER_ADDRESS"),
		DBHost:        viper.GetString("DB_HOST"),
		DBName:        viper.GetString("DB_NAME"),
		DBUser:        viper.GetString("DB_USER"),
		DBPass:        viper.GetString("DB_PASSWORD"),
		DBPort:        viper.GetString("DB_PORT"),
		DBConn:        conn,
	}

	// reading heroku provided port to handle deployment with heroku
	port := viper.GetString("PORT")
	if port != "" {
		logger.Log("using the port allocated by heroku", port)
		configs.ServerAddress = "0.0.0.0:" + port
	}

	logger.Log("serve port", configs.ServerAddress)
	logger.Log("db host", configs.DBHost)
	logger.Log("db name", configs.DBName)
	logger.Log("db port", configs.DBPort)

	return configs
}