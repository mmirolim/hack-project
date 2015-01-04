package datastore

import "time"

type Nots struct {
	ID        int       `db:"id" json:"id"`
	TableID   int       `db:"tableID" json:"tableID"`
	Msg       string    `db:"msg" json:"msg"`
	StaffID   int       `db:"staffID" json:"staffID"`
	Status    int       `db:"status" json:"status"`
	CreatedAt time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt" json:"updatedAt"`
}

// models table name
func (n Nots) TableName() string {
	return "notifications"
}

func (n Nots) GetID() int {
	return n.ID
}

func (n Nots) FieldNames() []string {
	return []string{"tableID", "msg", "staffID", "status", "createdAt", "updatedAt"}
}

// create table query
func (n Nots) createTableQuery() string {
	q := "CREATE TABLE IF NOT EXISTS " + n.TableName()
	q += ` ( 
		id INTEGER PRIMARY	KEY AUTOINCREMENT, 
		tableID INTEGER, 
		msg TEXT,
                staffID INTEGER, 
                status INTEGER,
		createdAt INTEGER, 
		updatedAt INTEGER
	)
	`
	return q
}

// defaults to set
func (n *Nots) SetDefaults() {
	n.CreatedAt = time.Now()
}

// validation logic for items
func (n Nots) Validate() error {
	var err error
	return err
}

//Create an order
func (n *Nots) Create() error {
	err := create(n,
		n.TableID,
		n.Msg,
		n.StaffID,
		n.Status,
		n.CreatedAt.Unix(),
		n.UpdatedAt.Unix(),
	)
	return err
}

func (n *Nots) FindOne(wh Where) error {
	var createdAt, updatedAt int64
	err := findOne(n, wh, &n.ID,
		&n.TableID,
		&n.Msg,
		&n.StaffID,
		&n.Status,
		&createdAt,
		&updatedAt,
	)
	n.CreatedAt = time.Unix(createdAt, 0)
	n.UpdatedAt = time.Unix(updatedAt, 0)

	return err
}

func (n *Nots) FindAll(wh Where, lim int) ([]Nots, error) {
	nots := []Nots{}
	rows, err := findAllRows(n, lim, wh)
	if err != nil {
		return nots, err
	}
	// don't forget to close rows
	defer rows.Close()
	// temp store for scan
	var nt Nots
	for rows.Next() {
		var createdAt, updatedAt int64
		err := rows.Scan(
			&nt.ID,
			&nt.TableID,
			&nt.Msg,
			&nt.StaffID,
			&nt.Status,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nots, err
		}
		nt.CreatedAt = time.Unix(createdAt, 0)
		nt.UpdatedAt = time.Unix(updatedAt, 0)
		nots = append(nots, nt)
	}
	return nots, err
}

func (n *Nots) Update() error {
	err := update(n,
		n.TableID,
		n.Msg,
		n.StaffID,
		n.Status,
		n.CreatedAt.Unix(),
		n.UpdatedAt.Unix(),
	)

	return err
}

func (n *Nots) Delete() error {
	err := del(n)
	return err
}
