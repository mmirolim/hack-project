package datastore

import (
	"strings"
	"testing"
	"time"

	"github.com/mmirolim/hack-project/conf"
)

func TestCreateStaff(t *testing.T) {
	var st Staff
	role := RoleStaff

	st.Login = "aziza"
	st.Password = "password"
	st.Name = "Aziza Mominova"
	st.Identity = "1234567"
	st.Role = role
	st.CreatedAt = time.Now()
	st.UpdatedAt = time.Now()

	mockConf := `
[ds]
	[ds.sqlite]
	name = "sqlite3"
	file = "foo.db"
[srv]
	port = "3000"
`

	f := strings.NewReader(mockConf)

	App, _ := conf.Read(f)

	_, err := Initialize(App.DS)
	if err != nil {
		t.Error(err)
	}

	err = st.Create()
	if err != nil {
		t.Error(err)
	}
}

func TestFindAllStaffs(t *testing.T) {
	//get all staffs
	var st Staff
	stfs, err := st.FindAll(Where{"id", ">", 0}, 0)
	if err != nil {
		t.Error(err)
	}
	if got := stfs[0].Name; got == "" {
		t.Errorf("Staff got %s, want %s", got, "not-emtpy")
	}

}

func TestFindOneStaff(t *testing.T) {
	//get staff
	var st Staff
	var err error
	err = st.FindOne(Where{"login", "=", "aziza"})
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateStaff(t *testing.T) {

	var st Staff
	var newSt Staff

	newSt.Login = "Akmal"
	newSt.Password = "password"
	newSt.Name = "Akmal Ikromov"
	newSt.Role = RoleAdmin

	//get staff
	err := st.FindOne(Where{"login", "=", "aziza"})
	if err != nil {
		t.Error(err)
	}

	//update staff
	newSt.Identity = st.Identity
	newSt.ID = st.ID
	err = newSt.Update()
	if err != nil {
		t.Error(err)
	}
	// find in db and match with wanted
	err = st.FindOne(Where{"login", "=", "Akmal"})
	if err != nil {
		t.Error(err)
	}
	if st.Name != newSt.Name {
		t.Errorf("Staff Name got %s, want %s", st.Name, newSt.Name)
	}
}

func TestDeleteStaff(t *testing.T) {
	//staff is deleted, but staff.Get(1) is returns the staff
	var st Staff
	err := st.FindOne(Where{"id", ">", 0})
	reportErr(t, err)
	err = st.Delete()
	reportErr(t, err)
}

func reportErr(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
