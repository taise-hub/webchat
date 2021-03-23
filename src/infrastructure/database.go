package infrastructure

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

type Config struct {
	User 	 string
	Password string
	Server	 string
	Port 	 int
	DBName 	 string
}

func (c *Config) setUpDsn() string {
	protocol := fmt.Sprintf("tcp(%s:%d)",c.Server, c.Port)
	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True", c.User, c.Password, protocol, c.DBName)
	return dsn
}

func NewConfig() *Config {
	return &Config {
		User: "",
		Password: "taise",
		Server: "localhost",
		Port: 3306,
		DBName: "chat_db",
	}
}

func NewDB(c *Config) (*gorm.DB, error) {
	fmt.Println("[DEBUG]"+c.setUpDsn())
	return gorm.Open(mysql.Open(c.setUpDsn()), &gorm.Config{})
}