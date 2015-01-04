package datastore

import (
	"strings"
	"testing"
	"time"

	"github.com/mmirolim/hack-project/conf"
)

func TestCreateNots(t *testing.T) {
	var nt Nots

	nt.TableID = 10
	nt.Msg = "your order accepted"
	nt.StaffID = 2
	nt.Status = 1
	nt.CreatedAt = time.Now()
	nt.UpdatedAt = time.Now()

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

	err = nt.Create()
	if err != nil {
		t.Error(err)
	}
}

func TestFindAllNots(t *testing.T) {
	//get all staffs
	var nt Nots
	nots, err := nt.FindAll(Where{"id", ">", 0}, 0)
	if err != nil {
		t.Error(err)
	}
	want := 10
	if got := nots[0].TableID; got != want {
		t.Errorf("Notification got %d, want %d", got, want)
	}

}

func TestFindOneNots(t *testing.T) {
	//get staff
	var nt Nots
	var err error
	err = nt.FindOne(Where{"tableID", "=", 10})
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateNots(t *testing.T) {

	var nt, nnt Nots

	nnt.Status = 2

	//get staff
	err := nt.FindOne(Where{"tableID", "=", 10})
	if err != nil {
		t.Error(err)
	}

	//update staff
	nnt.TableID = nt.TableID
	nnt.Msg = nt.Msg
	nnt.StaffID = nt.StaffID
	nnt.CreatedAt = nt.CreatedAt
	nnt.UpdatedAt = nt.UpdatedAt
	nnt.ID = nt.ID
	err = nnt.Update()
	if err != nil {
		t.Error(err)
	}
	// find in db and match with wanted
	err = nt.FindOne(Where{"tableID", "=", 10})
	if err != nil {
		t.Error(err)
	}
	want := 2
	if got := nt.Status; got != want {
		t.Errorf("Notification Status got %d, want %d", got, want)
	}
}

func TestDeleteNots(t *testing.T) {
	//staff is deleted, but staff.Get(1) is returns the staff
	var nt Nots
	err := nt.FindOne(Where{"id", ">", 0})
	reportErr(t, err)
	err = nt.Delete()
	reportErr(t, err)
}
