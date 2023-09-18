package db

import (
	"database/sql"
	"fmt"

	"github.com/akshayur0404/go-aws-s3/pkg/config"
	_ "github.com/lib/pq"
)

func ConnectDB(cfg config.Config) (*sql.DB, error) {
	psqlcon := fmt.Sprintf("host=%s port=%s passwrod=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBPassword, cfg.DBName)
	db, err := sql.Open("postgres", psqlcon)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil

}
