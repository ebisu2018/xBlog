package config

import (
	"encoding/json"
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DefaultConfig() *Config {
	return &Config{
		MySql: &MySql{
			Host:     "127.0.0.1",
			Port:     3306,
			Username: "",
			Password: "",
			Database: "",
		},
	}
}

type Config struct {
	*MySql `json:"mysql" toml:"mysql"`
}

type MySql struct {
	Host     string `json:"host" toml:"host"`
	Port     int    `json:"port" toml:"port"`
	Username string `json:"username" toml:"username"`
	Password string `json:"password" toml:"password"`
	Database string `json:"database" toml:"database"`

	db   *gorm.DB
	lock sync.Mutex
}

func (m *MySql) dsn() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.Database)
}

func (m *MySql) GetConn() *gorm.DB {
	if m.db == nil {
		m.lock.Lock()
		defer m.lock.Unlock()
		db, err := gorm.Open(mysql.Open(m.dsn()), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		m.db = db
	}
	return m.db.Debug()
}

func (c *Config) String() string {
	data, _ := json.Marshal(c)
	return string(data)
}
