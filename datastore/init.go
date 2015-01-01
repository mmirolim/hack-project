package datastore

import (
	"fmt"
	//"github.com/jinzhu/gorm"
	"database/sql"
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
	DB, err := sql.Open(ds.SQLite.Name, ds.SQLite.File)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v", DB)
	return DB, err
}

func migrate() {
	var result sql.Result
	sqlCreateTableOrders = `
		CREATE TABLE IF NOT EXISTS orders ( 
			ID INTEGER PRIMARY	KEY AUTOINCREMENT, 
			items			TEXT
			tableID			INTEGER , 
			cost			INTEGER,
			percentService	REAL,
			status			INTEGER,
			totalCost		INTEGER,
			createdAt		TEXT,
			updatedAt		TEXT,
			closedAt		TEXT
			staffID			INTEGER
		)
		`
	result, err = db.Exec(sqlCreateTableOrders)
	if err != nil {
		return err
	}
}
