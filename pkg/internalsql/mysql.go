package internalsql

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	"github.com/Khatchi/go-tweet/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySQL(cfg *config.Config) (*sql.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		url.QueryEscape("Africa/Lagos"),
	)

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database")
	}

	log.Println("database connected")
	return db, nil
}
