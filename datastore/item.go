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

/*
func (o *Order) FindOne(wh Where) error {
	var createdAt, updatedAt, closedAt int64
	var items string
	err := findOne(o, wh, &o.ID,
		&items,
		&o.TableID,
		&o.Cost,
		&o.PercentService,
		&o.Status,
		&o.TotalCost,
		&createdAt,
		&updatedAt,
		&closedAt,
		&o.StaffID,
	)
	o.CreatedAt = time.Unix(createdAt, 0)
	o.UpdatedAt = time.Unix(updatedAt, 0)
	o.CreatedAt = time.Unix(closedAt, 0)

	return err
}

func (o *Order) FindAll(wh Where, lim int) ([]Order, error) {
	ords := []Order{}
	rows, err := findAllRows(o, lim, wh)
	if err != nil {
		return ords, err
	}
	// don't forget to close rows
	defer rows.Close()
	// temp store for scan
	var order Order
	var items string
	for rows.Next() {
		var createdAt, updatedAt, closedAt int64
		err := rows.Scan(
			&order.ID,
			&items,
			&order.TableID,
			&order.Cost,
			&order.PercentService,
			&order.Status,
			&order.TotalCost,
			&createdAt,
			&updatedAt,
			&closedAt,
			&order.StaffID,
		)
		if err != nil {
			return ords, err
		}
		order.CreatedAt = time.Unix(createdAt, 0)
		order.UpdatedAt = time.Unix(updatedAt, 0)
		order.ClosedAt = time.Unix(closedAt, 0)

		err = json.Unmarshal([]byte(items), &order.Items)
		if err != nil {
			return ords, err
		}

		ords = append(ords, order)
	}
	return ords, err
}

func (order *Order) Update() error {
	items, err := json.Marshal(order.Items)
	if err != nil {
		return err
	}

	err = update(order,
		items,
		order.TableID,
		order.Cost,
		order.PercentService,
		order.Status,
		order.TotalCost,
		order.CreatedAt.Unix(),
		order.UpdatedAt.Unix(),
		order.ClosedAt.Unix(),
		order.StaffID,
	)

	return err
}

func (order *Order) Delete() error {
	err := del(order)
	return err
}
*/
