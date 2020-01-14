package dummydb

import "fmt"

// Pet holds the details of the pet
type Pet struct {
	ID        int64    `json:"id"`
	Category  Category `json:"category"`
	Name      string   `json:"name"`
	PhotoURLs []string `json:"photoUrls"`
	Tags      []Tag    `json:"tags"`
	Status    string
}

const petAvail string = "available"
const petPending string = "pending" // used when pet is ordered
const petSold string = "sold"       // used when order is shipped

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
	p.Status = petAvail
	petTBL = append(petTBL, p)
	return id
}

// PetByID returns pointer to Pet with specified ID
func PetByID(id int64) *Pet {
	for _, pet := range petTBL {
		if pet.ID == id {
			return pet
		}
	}
	return nil
}

// SetStatus used to modify status of Pet
func (p *Pet) SetStatus(status string) (string, error) {
	if status == petAvail || status == petPending || status == petSold {
		p.Status = status
		return p.Status, nil
	}
	return p.Status, fmt.Errorf("status '%s' not recognised; no change made", status)
}
