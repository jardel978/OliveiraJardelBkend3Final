package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var cfg *config

type APIConfig struct {
	Port    string
	Profile string
}

type DBConfig struct {
	DriverName   string
	Host         string
	Port         string
	User         string
	pass         string
	DataBase     string
	ReadTimeout  string
	WriteTimeout string
}

func (c *DBConfig) getURI() string {
	switch c.DriverName {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4,utf8\u0026readTimeout=%s\u0026writeTimeout=%s&parseTime=true",
			c.User, c.pass, c.Host, c.Port, c.DataBase, c.ReadTimeout, c.WriteTimeout)
	case "postgres":
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			c.Host, c.Port, c.User, c.pass, c.DataBase)
	}
	return ""
}

type config struct {
	API APIConfig
	DB  DBConfig
}

func init() {
	// valores default
	viper.SetDefault("api.port", "8080")
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "5433")
}

func Load() {
	// para ler arquivo de config
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalln(err)
		}
	}

	cfg = new(config) // o mesmo que &config{} (vai criar um ponteiro de config para gente)

	cfg.API = APIConfig{
		Port:    viper.GetString("api.port"),
		Profile: viper.GetString("db.profile"),
	}
	cfg.DB = DBConfig{
		DriverName:   viper.GetString("db.driverName"),
		Host:         viper.GetString("db.host"),
		Port:         viper.GetString("db.port"),
		User:         viper.GetString("db.user"),
		pass:         viper.GetString("db.password"),
		DataBase:     viper.GetString("db.name") + "_" + cfg.API.Profile,
		ReadTimeout:  viper.GetString("db.readTimeout"),
		WriteTimeout: viper.GetString("db.writeTimeout"),
	}
}

func GetServerPort() string {
	return cfg.API.Port
}
func getDBConfig() DBConfig {
	return cfg.DB
}
