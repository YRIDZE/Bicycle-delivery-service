package db_repository

import (
	"database/sql"
	"fmt"

	"github.com/YRIDZE/Bicycle-delivery-service/conf"
	log "github.com/YRIDZE/yolo-log"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	UsersTable        = "users"
	OrdersTable       = "orders"
	OPTable           = "order_products"
	SuppliersTable    = "suppliers"
	ProductsTable     = "products"
	CartTable         = "cart"
	CartProductsTable = "cart_products"
	CacheTokenTable   = "uid_token"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func NewDB(logger *log.Logger, cfg Config) (*sql.DB, error) {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName),
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

func InitDB(cfg *conf.Config) *sql.DB {
	db, err := NewDB(
		cfg.Logger,
		Config{
			Host:     cfg.ConfigDB.Host,
			Port:     cfg.ConfigDB.Port,
			Username: cfg.ConfigDB.Username,
			DBName:   cfg.ConfigDB.DBName,
			Password: cfg.ConfigDB.DbPassword,
		},
	)
	if err != nil {
		cfg.Logger.Fatal("Could not connected to database. Panic!")
		panic(err.Error())
	}

	return db
}
