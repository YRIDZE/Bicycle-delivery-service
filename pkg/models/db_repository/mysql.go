package db_repository

import (
	"database/sql"
	"fmt"
	"github.com/YRIDZE/Bicycle-delivery-service/internal"
	_ "github.com/go-sql-driver/mysql"
)

const (
	UsersTable      = "users"
	OrdersTable     = "orders"
	OPTable         = "order_products"
	CacheTokenTable = "uid_token"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func NewDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName),
	)
	if err != nil {
		internal.Log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
