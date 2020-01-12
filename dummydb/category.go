package dummydb

// Category holds the pet category (ie, animal type)
type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

var categoryID int64
var categoryTBL []*Category

// NewCategory adds a new Category to the db, returns the ID
func NewCategory(name string) int64 {
	categoryID++
	id := categoryID
	c := Category{id, name}
	categoryTBL = append(categoryTBL, &c)
	return id
}
