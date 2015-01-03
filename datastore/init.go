package datastore

import (
	"database/sql"
	"strings"

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
	GetID() int
	FieldNames() []string
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
	Crit  string
	Value interface{}
}

func (w *Where) trimSpace() (string, string, interface{}) {
	fld := strings.TrimSpace(w.Field)
	crit := strings.TrimSpace(w.Crit)
	var val interface{}
	switch v := w.Value.(type) {
	case string:
		val = strings.TrimSpace(v)
	default:
		val = v
	}
	return fld, crit, val
}

// returns link to result rows by some field
func findAllRows(m Model, lim int, wh Where) (*sql.Rows, error) {
	var err error
	fld, crit, val := wh.trimSpace()
	if lim == 0 {
		// protection from abuse
		lim = 5000
	}
	var where string
	if wh.Field != "" {
		where = " WHERE " + fld + crit + "? "
	}
	q := "SELECT * FROM " + m.TableName() + where + " LIMIT ?"
	rows, err := DB.Query(q, val, lim)
	if err != nil {
		return rows, err
	}
	return rows, err
}

// simple find one record by one field
// @todo refactor to make more flexible
func findOne(m Model, wh Where, args ...interface{}) error {
	var err error
	// trim spaces just in case (injection)
	fld, crit, val := wh.trimSpace()
	q := "SELECT *  FROM " + m.TableName() + " WHERE " + fld + crit + "? LIMIT 1"
	// write data to dest
	err = DB.QueryRow(q, val).Scan(args...)
	return err
}

func update(m Model, args ...interface{}) error {
	var err error
	var flds string
	id := m.GetID()
	for _, v := range m.FieldNames() {
		v += "=?,"
		flds += v
	}

	flds = flds[:len(flds)-1]
	q := "UPDATE " +
		m.TableName() +
		" SET " +
		flds +
		" WHERE id = " +
		string(id)

	_, err = DB.Exec(q, args)
	return err
}

func del(m Model) error {
	var err error
	id := m.GetID()
	q := "DELETE FROM " + m.TableName() + " WHERE id = ?"
	_, err = DB.Exec(q, id)
	return err
}

func create(m Model, args ...interface{}) error {
	var err error
	var flds string
	var qm string
	for _, f := range m.FieldNames() {
		f += ", "
		qm += "?, "
		flds += f
	}
	flds = flds[:len(flds)-1]
	qm = flds[:len(qm)-1]
	q := "INSERT INTO " +
		m.TableName() +
		"( " +
		flds +
		" )" +
		" VALUES( " +
		flds +
		")"

	_, err = DB.Exec(q, args)
	return err
}

func get(m Model, id int) (*sql.Rows, error) {
	var err error
	var flds string
	var rows *sql.Rows

	for _, f := range m.FieldNames() {
		f += ", "
		flds += f
	}
	flds = flds[:len(flds)-1]
	q := "SELECT " +
		flds +
		" FROM" +
		m.TableName() +
		" where id = " +
		string(m.GetID())

	rows, err = DB.Query(q, id)
	return rows, err
}
