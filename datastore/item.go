package datastore

type Item struct {
	ID      int     `json: id`
	Name    string  `json: name`
	Desc    string  `json: desc`
	Img     string  `json: img`
	Serving float32 `json: serving`
	Cost    int     `json: cost`
	Unit    string  `json: unit`
	Status  int     `json: status`
}
