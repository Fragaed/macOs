package repository

import (
	"Fragaed/internal/config"
	_ "Fragaed/internal/migration"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"go.uber.org/zap"
	"log"
	"time"
)

const (
	userTable = "users"
)

func NewPostgresDB(cfg config.Config) (*sqlx.DB, error) {
	var dsn string
	var err error
	var dbRaw *sql.DB

	dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.Password, cfg.DB.DBName, cfg.DB.SSLMode)
	fmt.Println("Connecting with DSN:", dsn)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	timeoutExceeded := time.After(time.Second * 60)

	for {
		select {
		case <-timeoutExceeded:
			return nil, fmt.Errorf("db connection failed after %d timeout %s", 5, err)
		case <-ticker.C:
			dbRaw, err = sql.Open("postgres", dsn)
			if err != nil {
				return nil, fmt.Errorf("failed to connect to db", zap.Error(err))
			}
			err = dbRaw.Ping()
			if err == nil {

				db := sqlx.NewDb(dbRaw, cfg.DB.Driver)
				db.SetMaxOpenConns(50)
				db.SetMaxIdleConns(50)

				err = goose.Up(dbRaw, "./")
				if err != nil {
					log.Fatal("Goose up failed ", err)
				}
				return db, nil
			}

			log.Fatal("failed to connect to the database", zap.String("dsn", dsn), zap.Error(err))
		}
	}
}
