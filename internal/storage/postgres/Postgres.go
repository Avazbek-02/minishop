package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/minishop/internal/config"
	"github.com/minishop/internal/logger"
)

func ConnectPostgres() (*sql.DB, error) {

	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	conf := config.Load()
	dns := "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
	db, err := sql.Open("postgres",
		fmt.Sprintf(dns, conf.DBHOST, conf.DBPORT, conf.DBUSER, conf.DBPASSWORD, conf.DBNAME))
	if err != nil {
		log.Error("Failed to connect to postgres")
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Error("Failed to ping postgres")
		return nil, err
	}
	return db, nil
}
