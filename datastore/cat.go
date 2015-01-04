package datastore

import "time"

type Cat struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Desc      string    `db:"desc" json:"desc"`
	CreatedAt time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt" json:"updatedAt"`
	StaffID   int       `db:"staffID" json:"staffID"`
}

// models table name
func (it Cat) TableName() string {
	return "cats"
}

func (i Cat) FieldNames() []string {
	return []string{"name", "desc", "createdAt", "updatedAt", "staffID"}
}

func (i Cat) GetID() int {
	return i.ID
}

// create table query
func (it Cat) createTableQuery() string {
	q := "CREATE TABLE IF NOT EXISTS " + it.TableName()
	q += ` ( 
		id INTEGER PRIMARY	KEY AUTOINCREMENT, 
		name TEXT, 
        desc TEXT,
		createdAt INTEGER, 
		updatedAt INTEGER, 
		staffID INTEGER
	)
	`
	return q
}

// defaults to set
func (it *Cat) SetDefaults() {
	it.CreatedAt = time.Now()
}

// validation logic for items
func (it Cat) Validate() error {
	var err error
	return err
}

//Create an order
func (i *Cat) Create() error {
	err := create(i,
		i.Name,
		i.Desc,
		i.CreatedAt.Unix(),
		i.UpdatedAt.Unix(),
		i.StaffID,
	)
	return err
}

func (i *Cat) FindOne(wh Where) error {
	var createdAt, updatedAt int64
	err := findOne(i, wh, &i.ID,
		&i.Name,
		&i.Desc,
		&createdAt,
		&updatedAt,
		&i.StaffID,
	)
	i.CreatedAt = time.Unix(createdAt, 0)
	i.UpdatedAt = time.Unix(updatedAt, 0)

	return err
}

func (i *Cat) FindAll(wh Where, lim int) ([]Cat, error) {
	itms := []Cat{}
	rows, err := findAllRows(i, lim, wh)
	if err != nil {
		return itms, err
	}
	// don't forget to close rows
	defer rows.Close()
	// temp store for scan
	var item Cat
	for rows.Next() {
		var createdAt, updatedAt int64
		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Desc,
			&createdAt,
			&updatedAt,
			&item.StaffID,
		)
		if err != nil {
			return itms, err
		}
		item.CreatedAt = time.Unix(createdAt, 0)
		item.UpdatedAt = time.Unix(updatedAt, 0)

		itms = append(itms, item)
	}
	return itms, err
}

func (i *Cat) Update() error {
	err := update(i,
		i.Name,
		i.Desc,
		i.CreatedAt.Unix(),
		i.UpdatedAt.Unix(),
		i.StaffID,
	)
	return err
}

func (i *Cat) Delete() error {
	err := del(i)
	return err
}
