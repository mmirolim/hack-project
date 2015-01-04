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
	Identity  string    `db:"identity" json:"identity"`
}

// models table name
func (s Staff) TableName() string {
	return "staff"
}

func (s Staff) GetID() int {
	return s.ID
}

func (s Staff) FieldNames() []string {
	return []string{"login", "password", "name", "role", "createdAt", "updatedAt", "identity"}
}

// create table query
func (s Staff) createTableQuery() string {
	q := "CREATE TABLE IF NOT EXISTS " + s.TableName()
	q += ` ( 
		id INTEGER PRIMARY	KEY AUTOINCREMENT, 
		login TEXT, 
		password TEXT,
                name TEXT, 
                role INTEGER,
		createdAt INTEGER, 
		updatedAt INTEGER, 
                identity TEXT
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

//Create an order
func (st *Staff) Create() error {
	st.SetDefaults()
	st.UpdatedAt = time.Now()
	err := create(st,
		st.Login,
		st.Password,
		st.Name,
		st.Role,
		st.CreatedAt.Unix(),
		st.UpdatedAt.Unix(),
		st.Identity,
	)
	return err
}

func (st *Staff) FindOne(wh Where) error {
	var createdAt, updatedAt int64
	err := findOne(st, wh, &st.ID,
		&st.Login,
		&st.Password,
		&st.Name,
		&st.Role,
		&createdAt,
		&updatedAt,
		&st.Identity,
	)
	st.CreatedAt = time.Unix(createdAt, 0)
	st.UpdatedAt = time.Unix(updatedAt, 0)

	return err
}

func (st *Staff) FindAll(wh Where, lim int) ([]Staff, error) {
	stfs := []Staff{}
	rows, err := findAllRows(st, lim, wh)
	if err != nil {
		return stfs, err
	}
	// don't forget to close rows
	defer rows.Close()
	// temp store for scan
	var stf Staff
	for rows.Next() {
		var createdAt, updatedAt int64
		err := rows.Scan(
			&stf.ID,
			&stf.Login,
			&stf.Password,
			&stf.Name,
			&stf.Role,
			&createdAt,
			&updatedAt,
			&stf.Identity,
		)
		if err != nil {
			return stfs, err
		}
		stf.CreatedAt = time.Unix(createdAt, 0)
		stf.UpdatedAt = time.Unix(updatedAt, 0)
		stfs = append(stfs, stf)
	}
	return stfs, err
}

func (st *Staff) Update() error {
	err := update(st,
		st.Login,
		st.Password,
		st.Name,
		st.Role,
		st.CreatedAt.Unix(),
		st.UpdatedAt.Unix(),
		st.Identity,
	)

	return err
}

func (st *Staff) Delete() error {
	err := del(st)
	return err
}
