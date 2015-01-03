package datastore

import "time"

type Item struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Desc      string    `db:"desc" json:"desc"`
	Img       string    `db:"img" json:"img"`
	Serving   float32   `db:"serving" json:"serving"`
	Cost      int       `db:"cost" json:"cost"`
	Unit      string    `db:"unit" json:"unit"`
	Status    int       `db:"status" json:"status"`
	CreatedAt time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt" json:"updatedAt"`
	StaffID   int       `db:"staffID" json:"staffID"`
}

// models table name
func (it Item) TableName() string {
	return "items"
}

func (i Item) FieldNames() []string {
	return []string{"name", "desc", "img", "serving", "cost", "unit", "status", "createdAt", "updatedAt", "staffID"}
}

func (i Item) GetID() int {
	return i.ID
}

// create table query
func (it Item) createTableQuery() string {
	q := "CREATE TABLE IF NOT EXISTS " + it.TableName()
	q += ` ( 
		id INTEGER PRIMARY	KEY AUTOINCREMENT, 
		name TEXT, 
                desc TEXT,
                img TEXT,
                serving REAL,
                cost INTEGER,
                unit TEXT,
		status INTEGER, 
		createdAt INTEGER, 
		updatedAt INTEGER, 
		staffID INTEGER
	)
	`
	return q
}

// defaults to set
func (it *Item) SetDefaults() {
	it.CreatedAt = time.Now()
}

// validation logic for items
func (it Item) Validate() error {
	var err error
	return err
}
