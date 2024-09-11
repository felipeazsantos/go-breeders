package main

import (
	"database/sql"
	"time"
)

const (
	maxOpenDbConn  = 25
	maxIndleDbConn = 25
	maxDBLifeTime  = 5 * time.Minute
)

func initMySQLDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// test our database
	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIndleDbConn)
	db.SetConnMaxLifetime(maxDBLifeTime)

	return db, nil
}
