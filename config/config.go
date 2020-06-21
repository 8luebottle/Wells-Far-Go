package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

var (
	c Config
	d *gorm.DB
)

type Config struct {
	APIName   string   `yaml:"api_name"`
	APIPrefix string   `yaml:"api_prefix"`
	DB        Database `yaml:"db"`
	Logger    Logger   `yaml:"logger"`
}

type Database struct {
	Driver   string `json:"driver" yaml:"driver"`
	Host     string `json:"host" yaml:"host"`
	Username string `json:"username" yaml:"username"`
	Port     uint   `json:"port" yaml:"port"`
	Verbose  bool   `json:"verbose" yaml:"verbose"`
	DBName   string `json:"dbname" yaml:"dbname"`
	Password string `json:"password" yaml:"password"`
	SSLMode  string `json:"ssl_mode" yaml:"sslmode"`
}

type Logger struct {
	IsProduction bool   `yaml:"is_production"`
	Level        string `yaml:"level"`
	Format       string `yaml:"format"`
	Output       string `yaml:"output"`
}

// ConnetDB returns database connection source as a string.
func (c Config) ConnectDB() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.DB.Host, c.DB.Port, c.DB.Username, c.DB.Password, c.DB.DBName, c.DB.SSLMode,
	)
}

// InitAPIConfig initializes API Configuration setting,
// such Database, and Logger.
func InitAPIConfig(configPath string) error {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return errors.Wrap(err, "is not exist")
	}

	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return errors.Wrap(err, "api config file")
	}

	if err = yaml.Unmarshal(configFile, &c); err != nil {
		return errors.Wrap(err, "unmarshal api config file")
	}

	if err := initDB(c); err != nil {
		return errors.Wrap(err, "init database")
	}
	return nil
}

// initDB initializes a new PostgreSQL database with debug mode.
func initDB(dbConfig Config) error {
	connectionInfo := dbConfig.ConnectDB()
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return errors.Wrap(err, "open database")
	}
	d = db.Debug()
	return nil
}

func GetConfig() Config {
	return c
}

// DBConn handles Database Connection.
func DBConn() *gorm.DB {
	return d
}
