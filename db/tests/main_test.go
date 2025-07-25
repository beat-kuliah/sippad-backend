package db_test

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	db "github/beat-kuliah/sip_pad_backend/db/sqlc"
	"github/beat-kuliah/sip_pad_backend/utils"
	"log"
	"os"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var testQuery *db.Store

const testDBName = "testdb"
const sslmode = "?sslmode=disable"

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("Could not load env config.", err)
	}

	conn, err := sql.Open(config.DBdriver, config.DB_source+sslmode)
	if err != nil {
		log.Fatalf("Could not connect to %s server %v", config.DBdriver, err)
	}

	// create database for testing purposes
	_, err = conn.Exec(fmt.Sprintf("CREATE DATABASE %s;", testDBName))
	if err != nil {
		log.Fatalf("Encountered an error creating database %v", err)
	}

	tconn, err := sql.Open(config.DBdriver, config.DB_source+testDBName+sslmode)
	if err != nil {
		teardown(conn)
		log.Fatalf("Could not connect to database %v", err)
	}

	driver, err := postgres.WithInstance(tconn, &postgres.Config{})
	if err != nil {
		teardown(conn)
		log.Fatalf("Could not create migrate driver %v", err)
	}

	mig, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", "../migrations"), config.DBdriver, driver)
	if err != nil {
		log.Fatalf("Migration setup failed %v", err)
	}

	if err = mig.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration up failed %v", err)
	}

	testQuery = db.NewStore(tconn)

	code := m.Run()

	tconn.Close()

	teardown(conn)

	os.Exit(code)
}

func teardown(conn *sql.DB) {
	_, err := conn.Exec(fmt.Sprintf("DROP DATABASE %s WITH (FORCE);", testDBName))
	if err != nil {
		log.Fatalf("Failed to drop test database %v", err)
	}
}
