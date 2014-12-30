package datastore

type Item struct {
	ID              int
	Name, Desc, Img string
	Serving         float32
	Cost            int
	Unit            string
	Status          int
}
