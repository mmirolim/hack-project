package datastore

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mmirolim/hack-project/conf"
)

var DB *sql.DB

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
	err = createTable([]Model{&order, &table, &item, &staff})
	if err != nil {
		return nil, err
	}

	return DB, err
}

func createTable(ms []Model) error {
	var err error
	for _, m := range ms {
		_, err = DB.Exec(m.createTableQuery())
	}
	return err
}

// type for where clause
type Where struct {
	Field string
	Value interface{}
}

func (w *Where) trimSpace() (string, interface{}) {
	fld := strings.TrimSpace(w.Field)
	var val interface{}
	switch v := w.Value.(type) {
	case string:
		val = strings.TrimSpace(v)
	default:
		val = v
	}
	return fld, val
}

// returns link to result rows by some field
func findAllRows(m Model, lim int, wh Where) (*sql.Rows, error) {
	var err error
	fld, val := wh.trimSpace()
	if lim == 0 {
		// protection from abuse
		lim = 5000
	}
	q := "SELECT * FROM " + m.TableName() + " WHERE " + fld + "=? LIMIT ?"
	rows, err := DB.Query(q, val, lim)
	return rows, err
}

// simple find one record by one field
// @todo refactor to make more flexible
func findOne(m Model, wh Where, dest interface{}) error {
	var err error
	// trim spaces just in case (injection)
	fld, val := wh.trimSpace()
	q := "SELECT * FROM " + m.TableName() + " WHERE " + fld + "=? LIMIT 1"
	// write data to dest
	err = DB.QueryRow(q, val).Scan(dest)
	return err
}
