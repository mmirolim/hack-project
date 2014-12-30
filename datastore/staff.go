package datastore

type Staff struct {
	ID                    int
	Login, Password, Name string
	Role                  Role
}
