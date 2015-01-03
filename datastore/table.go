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
type Tables []Table

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

func (table Table) Create() error {
	sql := ` 
			INSERT INTO tables(alias) VALUES(?)
			`
	_, err := DB.Exec(sql,
		table.Alias,
	)
	if err != nil {
		return err
	}

	return err
}

func (tables Tables) GetAll() (Tables, error) {
	var count int
	getTablesCountSQL := "SELECT COUNT() FROM tables"

	rows, err := DB.Query(getTablesCountSQL)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&count)
	}
	tables = make(Tables, count)
	getTablesSQL := `
			SELECT alias FROM tables 
			`
	rows, _ = DB.Query(getTablesSQL)
	if err != nil {
		return nil, err
	}

	i := 0
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(
			&tables[i].Alias,
		); err != nil {
			return nil, err
		}
		i += 1
	}

	return tables, err
}

func (table Table) Get(id int) (Table, error) {
	getTableByIDSQL := `
			SELECT	
					id, 
					alias
			FROM tables
			WHERE id = ?
			`

	rows, err := DB.Query(getTableByIDSQL, id)
	if err != nil {
		return table, err
	}

	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(
			&table.ID,
			&table.Alias,
		); err != nil {
			return table, err
		}
	}
	return table, err
}

func (table Table) Update(newTable Table) error {
	updateTableSQL := ` UPDATE tables 
						SET	alias = ?
						WHERE id = ?
						`
	_, err := DB.Exec(updateTableSQL,
		newTable.Alias,
		table.ID,
	)
	if err != nil {
		return err
	}
	return err
}

func (table Table) Delete() error {
	deleteTableSQL := "DELETE FROM tables WHERE id = ?"
	_, err := DB.Exec(deleteTableSQL, table.ID)
	if err != nil {
		return err
	}
	return err
}
