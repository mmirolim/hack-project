package datastore

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mmirolim/hack-project/conf"
)

var DB *sql.DB

type Role int
type Status int

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

type Model interface {
	TableName() string        // get model stored table name
	createTableQuery() string // get query to create appropriate table
	SetDefaults()             // set some defaults like dates
	Validate() error          // @todo refactor
}

func Initialize(ds conf.Datastore) (*sql.DB, error) {
	var err error
	DB, err = sql.Open(ds.SQLite.Name, ds.SQLite.File)
	if err != nil {
		return nil, err
	}

	// create tables for all models if required
	var order Order
	var table Table
	var item Item
	var staff Staff
	err = createTable(&order, DB)
	err = createTable(&table, DB)
	err = createTable(&item, DB)
	err = createTable(&staff, DB)
	if err != nil {
		return nil, err
	}

	return DB, err
}

func createTable(m Model, db *sql.DB) error {
	_, err := db.Exec(m.createTableQuery())
	return err
}

func FindAll(m Model, limit int, dest interface{}) error {
	var err error

	return err
}

func FindOne(m Model, wh map[string]interface{}, dest interface{}) error {
	var err error
	return err
}
