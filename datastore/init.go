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
	ID                    int
	Login, Password, Name string
	Role                  Role
}

type Order struct {
	ID             int
	Items          []Item
	Table          Table
	Cost           int
	PercentService float32
	TotalCost      int
	Date           time.Time
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
