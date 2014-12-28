package datastore

import "time"

type Role int

const (
	// define const roles
	_ Role = iota
	RoleAdmin
	RoleManager
	RoleUser
	RoleStaff
	RoleClient
)

type User struct {
	Login, Password, Name string
	Role                  Role
}

type Order struct {
	Items []Item
	Table Table
	Cost  int
	Date  time.Time
}

type Table struct {
	ID string
}

type Item struct {
	Name, Desc string
	Cost       int
}
