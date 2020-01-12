package dummydb

// Obviously in a real world situation our database would be some sort of
// SQL-based RDB. For our purposes here, this will have to do.

// Additionally, the database would be designed such that, for instance,
// category names, usernames, etc, would be unique. For simplicity I have
// not implemented such checks in our dummy database setup code.

// InitDB should be called to initialise our dummy database
func InitDB() {
	// Categories
	catIDDog := NewCategory("Dog")
	catIDCat := NewCategory("Cat")
	catIDGerbil := NewCategory("Gerbil")
	catIDSnake := NewCategory("Snake")

	tagIDCute := NewTag("Cute")
	tagIDFluffy := NewTag("Fluffy")
	tagIDFriendly := NewTag("Friendly")

	petIDFido := NewPet("Fido", catIDDog, []string{"fido.jpg"}, []int64{tagIDCute, tagIDFluffy, tagIDFriendly})
	petIDSev := NewPet("Severus", catIDSnake, []string{"severus.jpg"}, []int64{tagIDFriendly})

	NewPet("Moggy", catIDCat, []string{"moggy1.jpg", "moggy2.jpg"}, []int64{tagIDCute, tagIDFluffy})
	NewPet("Gerry", catIDGerbil, []string{"gerry.jpg"}, []int64{tagIDFluffy})
	NewPet("Rover", catIDDog, []string{"fido.jpg"}, []int64{tagIDCute, tagIDFluffy, tagIDFriendly})

	NewUser("Alice", "L00kingGl4ss", "alice@example.com", "Alice", "Liddell", "")
	NewUser("Bob", "EQ4L1s3r", "bob@example.com", "Robert", "McCall", "555-4200")
	NewUser("Charlie", "aNgEl5", "charlie@example.com", "Charles", "Townsend", "1 (857) CHARLIE")

	orderIDFido := NewOrder(petIDFido)
	OrderByID(orderIDFido).Shipped()

	NewOrder(petIDSev)
}
