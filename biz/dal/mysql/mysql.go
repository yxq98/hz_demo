package mysql

import (
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Mysql struct {
	w, r *sqlx.DB
}

var (
	client *Mysql
	once   sync.Once
)

func GetMysql() *Mysql {
	// todo update dsn
	once.Do(func() {
		client = &Mysql{
			w: initSqlxDB("read_dsn"),
			r: initSqlxDB("write_dsn"),
		}
	})
	return client
}

func initSqlxDB(dsn string) *sqlx.DB {
	db, err := sqlx.Open("mysql2", dsn)
	if err != nil {
		panic("init read mysql")
	}
	return db
}
