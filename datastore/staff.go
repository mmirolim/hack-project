package datastore

type Staff struct {
	ID                    int
	Login, Password, Name string
	Role                  Role
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
