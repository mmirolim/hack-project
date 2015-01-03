package datastore

import (
	"fmt"
	"github.com/mmirolim/hack-project/conf"
	"strings"
	"testing"
)

func TestCreateStaff(t *testing.T) {
	var staff Staff
	role := RoleStaff

	staff.Login = "aziza"
	staff.Password = "password"
	staff.Name = "Aziza Mominova"
	staff.Role = role

	mockConf := `
[ds]
	[ds.sqlite]
	name = "sqlite3"
	file = "../foo.db"
[srv]
	port = "3000"
`

	f := strings.NewReader(mockConf)

	App, _ := conf.Read(f)

	_, err := Initialize(App.DS)
	if err != nil {
		t.Error(err)
	}

	err = staff.Create()
	if err != nil {
		t.Error(err)
	}
}

func TestGetAllStaffs(t *testing.T) {
	var staffs Staffs

	//get all staffs
	staffs, err := staffs.GetAll()
	if err != nil {
		t.Error(err)
	}

	//debug
	fmt.Printf("%+v\n", staffs)
}

func TestGetStaff(t *testing.T) {
	//get staff
	var staff Staff
	var err error
	staff, err = staff.Get(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", staff)
}

func TestUpdateStaff(t *testing.T) {

	var staff Staff
	var newStaff Staff

	newStaff.Login = "Akmal"
	newStaff.Password = "password"
	newStaff.Name = "Akmal Ikromov"
	newStaff.Role = RoleAdmin

	//get staff
	staff, err := staff.Get(1)
	if err != nil {
		t.Error(err)
	}

	//update staff
	err = staff.Update(newStaff)
	if err != nil {
		t.Error(err)
	}

	//check updates
	staff, err = staff.Get(1)
	fmt.Printf("%+v\n", staff)
}

func TestDeleteStaff(t *testing.T) {
	//staff is deleted, but staff.Get(1) is returns the staff
	var staff Staff
	staff, err := staff.Get(1)
	staff.Delete()
	if err != nil {
		t.Error(staff)
	}
}

func TestDropStaff(t *testing.T) {
	err := DropTable()
	if err != nil {
		fmt.Println(err)
	}
}
