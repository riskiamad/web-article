package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config: load environment variables, return struct of every environment variable.
var Config = loadConfig()

// config: struct hold every environment variables.
type config struct {
	ServerHost string
	DbHost     string
	DbUser     string
	DbPass     string
	DbName     string
	DbTestName string
	EsURL      string
	EsTestURL  string
	EsUsername string
	EsPassword string
}

// loadConfig: read environment variable.
func loadConfig() *config {
	c := new(config)

	_ = godotenv.Load()

	c.ServerHost = os.Getenv("SERVER_HOST")
	c.DbHost = os.Getenv("MYSQL_HOST")
	c.DbUser = os.Getenv("MYSQL_USER")
	c.DbPass = os.Getenv("MYSQL_PASSWORD")
	c.DbName = os.Getenv("MYSQL_DBNAME")
	c.DbTestName = os.Getenv("MYSQL_TEST_DBNAME")
	c.EsURL = os.Getenv("ELASTICSEARCH_URL")
	c.EsTestURL = os.Getenv("ELASTICSEARCH_TEST_URL")
	c.EsUsername = os.Getenv("ELASTICSEARCH_USERNAME")
	c.EsPassword = os.Getenv("ELASTICSEARCH_PASSWORD")

	return c
}
