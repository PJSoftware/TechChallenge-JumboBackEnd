package endpoint

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"../dummydb"
)

func (r *Request) handlePets() {
	if !r.validKey {
		r.status(http.StatusBadRequest, "Invalid ID supplied")
		return
	}

	extractID, err := regexp.Compile("^/(\\d+)(/.+)?$")
	if err != nil {
		r.status(http.StatusInternalServerError, err.Error())
		return
	}

	if extractID.MatchString(r.suburl) {
		matches := extractID.FindStringSubmatch(r.suburl)
		petID, err := strconv.Atoi(matches[1])
		uploadImage := matches[2] == "/uploadImage"
		if uploadImage {
			if r.method != "POST" {
				r.status(http.StatusMethodNotAllowed, "only POST valid for image upload")
				return
			}

		}
		if err != nil {
			r.status(http.StatusInternalServerError, err.Error())
			return
		}

		p := r.lookupPet(int64(petID))
		if p == nil {
			return
		}

		switch r.method {
		case "GET":
			r.findPet(p)
		case "POST":
			if uploadImage {
				r.uploadImage(p)
			} else {
				r.updatePet(p)
			}
		case "DELETE":
			r.deletePet(p)
		default:
			r.status(http.StatusMethodNotAllowed, "unrecognised method for processing pets")
		}
		return
	}

	if r.suburl == "/findByStatus" {
		if r.method == "GET" {
			r.findPetsByStatus()
		} else {
			r.status(http.StatusMethodNotAllowed, "unrecognised method for processing pets")
		}
	} else if r.suburl == "" { // url == "/pet" only
		switch r.method {
		case "POST":
			r.addPet()
		case "PUT":
			r.updatePetByObject()
		default:
			r.status(http.StatusMethodNotAllowed, "unrecognised method for processing pets")
		}
	} else {
		r.status(http.StatusMethodNotAllowed, "unrecognised method for processing pets")
	}

}

func (r *Request) lookupPet(v interface{}) *dummydb.Pet {
	var p *dummydb.Pet
	switch v.(type) {
	case int64:
		p = dummydb.PetByID(v.(int64))
	default:
		msg := fmt.Sprintf("unknown interface '%v' in lookupPet()", v)
		r.status(http.StatusInternalServerError, msg)
		return nil
	}
	if p == nil {
		r.status(http.StatusNotFound, "Pet not found")
	}
	return p
}

// GET
// ​/pet​/{petId}
// Find pet by ID
func (r *Request) findPet(p *dummydb.Pet) {
	r.w.WriteHeader(http.StatusOK)
	r.w.Write(r.jsonOut(*p))
}

// POST
// ​/pet​/{petId}
// Updates a pet in the store with form data
func (r *Request) updatePet(p *dummydb.Pet) {
	err := r.req.ParseForm()
	if err != nil {
		r.status(http.StatusInternalServerError, "Invalid input: form data not found")
		return
	}

	name := r.req.Form.Get("name")
	status := r.req.Form.Get("status")

	if status != "" {
		_, err := p.SetStatus(status)
		if err != nil {
			r.status(http.StatusMethodNotAllowed, "Invalid input: bad status value")
			return
		}
	}

	if name != "" {
		p.SetName(name)
	}

	r.findPet(p) // Return details of changed pet on success
}

// DELETE
// ​/pet​/{petId}
// Deletes a pet
func (r *Request) deletePet(p *dummydb.Pet) {
	p.Delete()
	r.status(http.StatusOK, "pet deleted successfully")
}

// POST
// ​/pet​/{petId}​/uploadImage
// uploads an image
func (r *Request) uploadImage(p *dummydb.Pet) {
	err := r.req.ParseForm()
	if err != nil {
		r.status(http.StatusInternalServerError, "Invalid input: form data not found")
		return
	}

	// metadata not actually stored in current specified model but
	// presumably this would be an image description?
	// metadata := r.req.Form.Get("additionalMetadata")
	file := r.req.Form.Get("file")

	// In a real world situation this would be doing a fair bit more
	if file != "" {
		p.AddImage(file)
	} else {
		r.status(http.StatusBadRequest, "file not found")
	}

	r.status(http.StatusOK, "image uploaded successfully")
}

// POST
// ​/pet
// Add a new pet to the store
func (r *Request) addPet() {
	pet := r.jsonToPet()
	if pet != nil {
		pet.Insert()
		r.status(http.StatusOK, "pet added successfully")
	}
}

// PUT
// ​/pet
// Update an existing pet
func (r *Request) updatePetByObject() {
	pet := r.jsonToPet()
	if pet != nil {
		id := pet.ID
		orig := dummydb.PetByID(id)
		if orig == nil {
			r.status(http.StatusNotFound, "Pet not found")
			return
		}

		orig.Category = pet.Category
		orig.Name = pet.Name
		orig.PhotoURLs = pet.PhotoURLs
		orig.Tags = pet.Tags
		orig.Status = pet.Status
		r.status(http.StatusOK, "pet updated successfully")
	}
}

// GET
// ​/pet​/findByStatus
// Finds Pets by status
func (r *Request) findPetsByStatus() {
	var plist []dummydb.Pet
	if status, ok := r.data["status"]; ok {
		plist = dummydb.PetsByStatus(status)
	}

	r.w.WriteHeader(http.StatusOK)
	r.w.Write(r.jsonOut(plist))
}

func (r *Request) jsonToPet() *dummydb.Pet {
	dec := r.extractJSON()
	if dec == nil {
		r.status(http.StatusMethodNotAllowed, "invalid input")
		return nil
	}

	pet := new(dummydb.Pet)
	err := dec.Decode(&pet)
	if err != nil {
		r.status(http.StatusMethodNotAllowed, "invalid input")
		return nil
	}

	return pet
}
