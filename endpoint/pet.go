package endpoint

import (
	"regexp"
	"strconv"
)

// GET
// ​/pet​/{petId}
// Find pet by ID

// POST
// ​/pet​/{petId}
// Updates a pet in the store with form data

// DELETE
// ​/pet​/{petId}
// Deletes a pet

// POST
// ​/pet​/{petId}​/uploadImage
// uploads an image

// POST
// ​/pet
// Add a new pet to the store

// PUT
// ​/pet
// Update an existing pet

// GET
// ​/pet​/findByStatus
// Finds Pets by status

func (r *Request) handlePets() {
	extractID := regexp.MustCompile("^/(\\d+)$")
	if extractID.MatchString(r.suburl) {
		matches := extractID.FindStringSubmatch(r.suburl)
		petID, err := strconv.Atoi(matches[1])
		if err != nil {
			r.status(500, "error identifying pet")
		}
		switch r.method {
		case "GET":
			r.findPet(int64(petID))
		case "POST":
			r.updatePet(int64(petID))
		case "DELETE":
			r.deletePet(int64(petID))
		default:
			r.status(405, "unrecognised method for processing pets")
		}
		return
	}

}

func (r *Request) findPet(id int64) {

}
func (r *Request) updatePet(id int64) {

}
func (r *Request) deletePet(id int64) {

}
