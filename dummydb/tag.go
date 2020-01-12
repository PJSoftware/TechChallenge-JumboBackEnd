package dummydb

// Tag is used as searchable pet descriptors
type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

var tagID int64
var tagTBL []*Tag

// NewTag adds a new tag to the db, returns the ID
func NewTag(name string) int64 {
	tagID++
	id := tagID
	t := Tag{id, name}
	tagTBL = append(tagTBL, &t)
	return id
}
