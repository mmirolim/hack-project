package datastore

import (
	"time"

	"github.com/mmirolim/hack-project/conf"
)

type Role int
type Status int

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

func Initialize(ds conf.Datastore) {

}

type Staff struct {
	ID                    int
	Login, Password, Name string
	Role                  Role
}

type Order struct {
	ID             int
	Items          []Item
	TableID        int
	Cost           int
	PercentService float32
	Status         Status
	TotalCost      int
	Date           time.Time
	StaffID        int
}

type Table struct {
	ID    int
	Alias string
}

type Item struct {
	ID              int
	Name, Desc, Img string
	Serving         float32
	Cost            int
}
