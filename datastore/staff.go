package datastore

import "time"

type Staff struct {
	ID        int       `db:"id" json:"id"`
	Login     string    `db:"login" json:"login"`
	Password  string    `db:"password" json:"password"`
	Name      string    `db:"name" json:"name"`
	Role      Role      `db:"role" json:"role"`
	CreatedAt time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt" json:"updatedAt"`
	StaffID   int       `db:"staffID" json:"staffID"`
}

// models table name
func (s Staff) TableName() string {
	return "staffs"
}

// create table query
func (s Staff) createTableQuery() string {
	q := "CREATE TABLE IF NOT EXISTS " + s.TableName()
	q += ` ( 
		id INTEGER PRIMARY	KEY AUTOINCREMENT, 
		name TEXT, 
		login TEXT, 
		password TEXT, 
                role INTEGER,
		createdAt INTEGER, 
		updatedAt INTEGER, 
		staffID INTEGER
	)
	`
	return q
}

// defaults to set
func (s *Staff) SetDefaults() {
	s.CreatedAt = time.Now()
}

// validation logic for items
func (s Staff) Validate() error {
	var err error
	return err
}
