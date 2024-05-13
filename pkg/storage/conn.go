package storage

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitSqlx() (err error) {
	db, err = sqlx.Connect("mysql", "root:root@tcp(192.168.1.6:3306)/test?charset=utf8")
	if err != nil {
		return fmt.Errorf("connect db err: %w", err)
	}

	return nil
}
