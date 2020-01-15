package dummydb

import (
	"fmt"
	"strings"
)

// Pet holds the details of the pet
type Pet struct {
	ID        int64    `json:"id"`
	Category  Category `json:"category"`
	Name      string   `json:"name"`
	PhotoURLs []string `json:"photoUrls"`
	Tags      []Tag    `json:"tags"`
	Status    string   `json:"status"`
}

type petStatus struct {
	Available string
	Pending   string
	Sold      string
	Deleted   string
}

// PetStatus used as enum
var PetStatus = petStatus{Available: "available", Pending: "pending", Sold: "sold", Deleted: "deleted"}

var petID int64
var petTBL []*Pet

// NewPet adds a new Pet to the db, returns the ID
func NewPet(name string, category Category, photoURLs []string, tags []Tag) int64 {
	petID++
	id := petID
	p := new(Pet)
	p.ID = id
	p.Name = name
	p.Category = category
	for _, u := range photoURLs {
		p.PhotoURLs = append(p.PhotoURLs, u)
	}
	for _, t := range tags {
		p.Tags = append(p.Tags, t)
	}
	p.Status = PetStatus.Available
	petTBL = append(petTBL, p)
	return id
}

// PetByID returns pointer to Pet with specified ID
func PetByID(id int64) *Pet {
	for _, pet := range petTBL {
		// Ensure we do not return DELETED Pets
		if pet.ID == id && pet.Status != PetStatus.Deleted {
			return pet
		}
	}
	return nil
}

// PetsByStatus returns list of pets matching specified status(es)
func PetsByStatus(st string) []Pet {
	var plist []Pet
	slist := strings.Split(st, ",")
	for _, status := range slist {
		if status == PetStatus.Available || status == PetStatus.Pending || status == PetStatus.Sold {
			for _, pet := range petTBL {
				if pet.Status == status {
					plist = append(plist, *pet)
				}
			}
		}
	}

	return plist
}

// Delete deletes Pet with specified ID
// In an RDB we would probably use DELETE which is an atomic function
// so for this task we shall use the "Deleted" status instead;
// this should be safer (and quicker) than slice manipulation
func (p *Pet) Delete() {
	p.Status = PetStatus.Deleted
}

// SetStatus used to modify status of Pet
func (p *Pet) SetStatus(status string) (string, error) {
	if status == PetStatus.Available || status == PetStatus.Pending || status == PetStatus.Sold {
		p.Status = status
		return p.Status, nil
	}
	return p.Status, fmt.Errorf("status '%s' not recognised; no change made", status)
}

// SetName used to change name of Pet
func (p *Pet) SetName(name string) {
	p.Name = name
}

// AddImage adds the specified filename to our Pet
// In the real world this would handle uploading the file and storing it
// on the server, but for this task we can simply append to the list
func (p *Pet) AddImage(file string) {
	p.PhotoURLs = append(p.PhotoURLs, file)
}

// Insert new Pet entry into table; calls setNewID() because we don't trust
// the existing one
func (p *Pet) Insert() {
	p.setNewID()
	petTBL = append(petTBL, p)
}

func (p *Pet) setNewID() {
	petID++
	p.ID = petID
}
