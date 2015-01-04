package datastore

//Create a table
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

func (t Table) FieldNames() []string {
	return []string{"alias", "desc", "createdAt", "updatedAt", "staffID"}
}

func (t Table) GetID() int {
	return t.ID
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

func (tbl *Table) Create() error {
	err := create(tbl,
		tbl.Alias,
		tbl.Desc,
		tbl.CreatedAt.Unix(),
		tbl.UpdatedAt.Unix(),
		tbl.StaffID,
	)

	return err
}

func (tbl *Table) FindOne(wh Where) error {
	var createdAt, updatedAt int64
	err := findOne(tbl, wh,
		&tbl.ID,
		&tbl.Alias,
		&tbl.Desc,
		&createdAt,
		&updatedAt,
		&tbl.StaffID,
	)
	tbl.CreatedAt = time.Unix(createdAt, 0)
	tbl.UpdatedAt = time.Unix(updatedAt, 0)

	return err
}

func (tbl *Table) FindAll(wh Where, lim int) ([]Table, error) {
	tbls := []Table{}
	rows, err := findAllRows(tbl, lim, wh)
	if err != nil {
		return tbls, err
	}

	defer rows.Close()
	var t Table
	for rows.Next() {
		var createdAt, updatedAt int64
		err := rows.Scan(
			&t.ID,
			&t.Alias,
			&t.Desc,
			&createdAt,
			&updatedAt,
			&t.StaffID,
		)
		if err != nil {
			return tbls, err
		}
		t.CreatedAt = time.Unix(createdAt, 0)
		t.UpdatedAt = time.Unix(updatedAt, 0)
		tbls = append(tbls, t)
	}

	return tbls, err
}

func (tbl *Table) Update() error {
	err := update(tbl,
		tbl.Alias,
		tbl.Desc,
		tbl.CreatedAt.Unix(),
		tbl.UpdatedAt.Unix(),
		tbl.StaffID,
	)
	return err
}

func (tbl *Table) Delete() error {
	err := del(tbl)
	return err
}
