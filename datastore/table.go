package datastore

type Table struct {
	ID    int
	Alias string
}

type Tables []Table

//Create a table
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
