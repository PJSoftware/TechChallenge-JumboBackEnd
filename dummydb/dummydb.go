package dummydb

// Obviously in a real world situation our database would be some sort of
// SQL-based RDB. For our purposes here, this will have to do.

// Additionally, the database would be designed such that, for instance,
// category names, usernames, etc, would be unique. For simplicity I have
// not implemented such checks in our dummy database setup code.

// InitDB should be called to initialise our dummy database
func InitDB() {
	// Categories in dummydb
	categoryDog := NewCategory("Dog")
	categoryCat := NewCategory("Cat")
	categoryGerbil := NewCategory("Gerbil")
	categorySnake := NewCategory("Snake")

	// Tags in dummydb
	tagCute := NewTag("Cute")
	tagFluffy := NewTag("Fluffy")
	tagFriendly := NewTag("Friendly")
	tagBig := NewTag("Big")
	tagSmall := NewTag("Small")

	// Pets in dummydb
	petIDFido := NewPet("Fido", *categoryDog, []string{"fido.jpg"}, []Tag{*tagCute, *tagFluffy, *tagFriendly, *tagBig})
	petIDSev := NewPet("Severus", *categorySnake, []string{"severus.jpg"}, []Tag{*tagFriendly, *tagBig})

	NewPet("Moggy", *categoryCat, []string{"moggy1.jpg", "moggy2.jpg"}, []Tag{*tagCute, *tagFluffy, *tagSmall})
	NewPet("Gerry", *categoryGerbil, []string{"gerry.jpg"}, []Tag{*tagFluffy, *tagSmall})
	NewPet("Rover", *categoryDog, []string{"fido.jpg"}, []Tag{*tagCute, *tagFluffy, *tagFriendly})

	// Users in dummydb
	NewUser("Alice", "L00kingGl4ss", "alice@example.com", "Alice", "Liddell", "N/A")
	NewUser("Bob", "EQ4L1s3r", "bob@example.com", "Robert", "McCall", "555-4200")
	NewUser("Charlie", "aNgEl5", "charlie@example.com", "Charles", "Townsend", "1 (857) CHARLIE")

	// Orders in dummydb
	orderIDFido := NewOrder(petIDFido)
	OrderByID(orderIDFido).Shipped()

	NewOrder(petIDSev)
}
