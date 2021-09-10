package db_repository

import (
	"context"
	"database/sql"
	"fmt"

	log "github.com/YRIDZE/yolo-log"
	_ "github.com/go-sql-driver/mysql"
)

const (
	UsersTable      = "users"
	OrdersTable     = "orders"
	OPTable         = "order_products"
	SuppliersTable  = "suppliers"
	ProductsTable   = "products"
	CacheTokenTable = "uid_token"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func NewDB(ctx context.Context, cfg Config) (*sql.DB, error) {
	logger := ctx.Value("logger").(*log.Logger)

	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName),
	)
	if err != nil {
		logger.Fatalf("database error: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Fatalf("database error: %v", err)
		return nil, err
	}

	return db, nil
}
