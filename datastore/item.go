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
	CatID     int       `db:"catID" json:"catID"`
}

// models table name
func (it Item) TableName() string {
	return "items"
}

func (i Item) FieldNames() []string {
	return []string{"name", "desc", "img", "serving", "cost", "unit", "status", "createdAt", "updatedAt", "staffID", "catID"}
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
		staffID INTEGER,
		catID INTEGER
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

//Create an order
func (i *Item) Create() error {
	err := create(i,
		i.Name,
		i.Desc,
		i.Img,
		i.Serving,
		i.Cost,
		i.Unit,
		i.Status,
		i.CreatedAt.Unix(),
		i.UpdatedAt.Unix(),
		i.StaffID,
		i.CatID,
	)
	return err
}

func (i *Item) FindOne(wh Where) error {
	var createdAt, updatedAt, closedAt int64
	err := findOne(i, wh, &i.ID,
		&i.Name,
		&i.Desc,
		&i.Img,
		&i.Serving,
		&i.Cost,
		&i.Unit,
		&i.Status,
		&createdAt,
		&updatedAt,
		&i.StaffID,
		&i.CatID,
	)
	i.CreatedAt = time.Unix(createdAt, 0)
	i.UpdatedAt = time.Unix(updatedAt, 0)
	i.CreatedAt = time.Unix(closedAt, 0)

	return err
}

func (i *Item) FindAll(wh Where, lim int) ([]Item, error) {
	itms := []Item{}
	rows, err := findAllRows(i, lim, wh)
	if err != nil {
		return itms, err
	}
	// don't forget to close rows
	defer rows.Close()
	// temp store for scan
	var item Item
	for rows.Next() {
		var createdAt, updatedAt int64
		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Desc,
			&item.Img,
			&item.Serving,
			&item.Cost,
			&item.Unit,
			&item.Status,
			&createdAt,
			&updatedAt,
			&item.StaffID,
			&item.CatID,
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

func (i *Item) Update() error {
	err := update(i,
		i.Name,
		i.Desc,
		i.Img,
		i.Serving,
		i.Cost,
		i.Unit,
		i.Status,
		i.CreatedAt.Unix(),
		i.UpdatedAt.Unix(),
		i.StaffID,
		i.CatID,
	)
	return err
}

func (i *Item) Delete() error {
	err := del(i)
	return err
}
