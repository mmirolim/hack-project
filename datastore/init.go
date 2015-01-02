package datastore

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mmirolim/hack-project/conf"
)

type Role int
type Status int

var DB *sql.DB

const (
	// define const roles
	_ Role = iota
	RoleAdmin
	RoleManager
	RoleUser
	RoleStaff
	RoleClient
)
const (
	// define const statuses
	_ Status = iota
	StatusIssued
	StatusAccepted
	StatusInProgress
	StatusReady
	StatusDelivered
	StatusPaid
	StatusCanceled
)

func Initialize(ds conf.Datastore) (*sql.DB, error) {
	var err error
	DB, err = sql.Open(ds.SQLite.Name, ds.SQLite.File)
	if err != nil {
		return nil, err
	}

	err = Migrate()
	if err != nil {
		return nil, err
	}
	return DB, err
}

func Migrate() error {
	sqlCreateTableOrders := `
	CREATE TABLE IF NOT EXISTS orders ( 
		ID INTEGER PRIMARY	KEY AUTOINCREMENT, 
		items TEXT, 
		tableID INTEGER , 
		cost INTEGER, 
		percentService REAL, 
		status INTEGER, 
		totalCost INTEGER, 
		createdAt INTEGER, 
		updatedAt INTEGER, 
		closedAt INTEGER, 
		staffID INTEGER
	)
	`

	result, err := DB.Exec(sqlCreateTableOrders)
	if err != nil {
		return err
		fmt.Printf("%+v", result)
	}
	return err
}
