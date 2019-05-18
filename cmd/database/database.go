package database

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLConfig struct {
	User string
	Password string
	Host string
	Port int
	Database string
}

type MySQL struct {
	Master *sql.DB
}

func NewMySQL (config MySQLConfig) (*MySQL, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	))

	if err != nil {
		return nil, err
	}

	return &MySQL{Master: db,}, nil
}


