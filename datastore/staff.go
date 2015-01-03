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
	return "staff"
}

func (s Staff) GetID() int {
	return s.ID
}

func (s Staff) FieldNames() []string {
	return []string{"login", "password", "name", "role", "createAt", "updatedAt", "staffID"}
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

//staff in plural form is also staff, but...
type Staffs []Staff

//Create a staff
func (staff Staff) Create() error {
	sql := ` 
			INSERT INTO staff(login, password, name, role) VALUES(?,?,?,?)
			`
	_, err := DB.Exec(sql,
		staff.Login,
		staff.Password,
		staff.Name,
		staff.Role,
	)
	if err != nil {
		return err
	}

	return err
}

func (staffs Staffs) GetAll() (Staffs, error) {
	var count int
	getStaffCountSQL := "SELECT COUNT() FROM staff"

	rows, err := DB.Query(getStaffCountSQL)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&count)
	}
	staffs = make(Staffs, count)
	getStaffsSQL := `
			SELECT id, login, password, name, role FROM staff 
			`
	rows, _ = DB.Query(getStaffsSQL)
	if err != nil {
		return nil, err
	}

	i := 0
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(
			&staffs[i].ID,
			&staffs[i].Login,
			&staffs[i].Password,
			&staffs[i].Name,
			&staffs[i].Role,
		); err != nil {
			return nil, err
		}
		i += 1
	}

	return staffs, err
}

func (staff Staff) Get(id int) (Staff, error) {
	getStaffByIDSQL := `
			SELECT	
				id,
				login,
				password,
				name,
				role
			FROM staff
			WHERE id = ?
			`

	rows, err := DB.Query(getStaffByIDSQL, id)
	if err != nil {
		return staff, err
	}

	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(
			&staff.ID,
			&staff.Login,
			&staff.Password,
			&staff.Name,
			&staff.Role,
		); err != nil {
			return staff, err
		}
	}
	return staff, err
}

func (staff Staff) Update(newStaff Staff) error {
	updateStaffSQL := ` UPDATE staff 
						SET	login = ?,
							password = ?,
							name = ?,
							role = ?
						WHERE id = ?
						`
	_, err := DB.Exec(updateStaffSQL,
		newStaff.Login,
		newStaff.Password,
		newStaff.Name,
		newStaff.Role,
		staff.ID,
	)
	if err != nil {
		return err
	}
	return err
}

func (staff Staff) Delete() error {
	deleteStaffSQL := "DELETE FROM staff WHERE id = ?"
	_, err := DB.Exec(deleteStaffSQL, staff.ID)
	if err != nil {
		return err
	}
	return err
}
