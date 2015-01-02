package datastore

import "time"

type Table struct {
	ID        int       `db:"id" json:"id"`
	Alias     string    `db:"alias" json:"alias"`
	Desc      string    `db:"desc" json:"desc"`
	CreatedAt time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt" json:"updatedAt"`
	StaffID   int       `db:"staffID" json:"staffID"`
}

// models table name
func (t Table) TableName() string {
	return "tables"
}

// create table query
func (s Table) createTableQuery() string {
	q := "CREATE TABLE IF NOT EXISTS " + s.TableName()
	q += ` ( 
		id INTEGER PRIMARY	KEY AUTOINCREMENT, 
		alias TEXT, 
		desc TEXT, 
		createdAt INTEGER, 
		updatedAt INTEGER, 
		staffID INTEGER
	)
	`
	return q
}

// defaults to set
func (t *Table) SetDefaults() {
	t.CreatedAt = time.Now()
}

// validation logic for items
func (t Table) Validate() error {
	var err error
	return err
}
