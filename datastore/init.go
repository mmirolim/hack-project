package datastore

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mmirolim/hack-project/conf"
)

type Role int
type Status int

var DB gorm.DB

const (
	// define const roles
	_ Role = iota
	RoleAdmin
	RoleManager
	RoleUser
	RoleStaff
	RoleClient
)
const (
	// define const statuses
	_ Status = iota
	StatusIssued
	StatusAccepted
	StatusInProgress
	StatusReady
	StatusPaid
	StatusCanceled
)

func Initialize(ds conf.Datastore) error {
	DB, err := gorm.Open(ds.SQLite.Name, ds.SQLite.File)
	if err != nil {
		return err
	}

	DB.CreateTable(&Item{})
	DB.CreateTable(&Order{})
	DB.CreateTable(&Staff{})
	DB.CreateTable(&Table{})

	return err
}
