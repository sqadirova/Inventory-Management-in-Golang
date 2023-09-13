package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

var connection *sql.DB

func DBConn() (*sql.DB, error) {
	if connection == nil {
		dbDriver := "postgres"
		dbConfig := Configuration.Database
		newConnectionDB, err := sql.Open(dbDriver, dbConfig.URL())
		newConnectionDB.SetMaxIdleConns(dbConfig.MaxIdleConn)
		newConnectionDB.SetMaxOpenConns(dbConfig.MaxOpenConn)
		newConnectionDB.SetConnMaxLifetime(time.Duration(dbConfig.MaxConnLifetimeHour))

		if err != nil {
			fmt.Println("err happens: ", err)
			return nil, err
		}

		connection = newConnectionDB

		return newConnectionDB, nil
	}

	return connection, nil
}
