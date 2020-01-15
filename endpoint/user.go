package endpoint

import "net/http"

// GET
// ​/user​/{username}
// Get user by user name

// PUT
// ​/user​/{username}
// Updated user

// DELETE
// ​/user​/{username}
// Delete user

// GET
// ​/user​/login
// Logs user into the system

// GET
// ​/user​/logout
// Logs out current logged in user session

// POST
// ​/user
// Create user

// POST
// ​/user​/createWithArray
// Creates list of users with given input array

// POST
// ​/user​/createWithList
// Creates list of users with given input array

func (r *Request) handleUser() {
	r.status(http.StatusNotImplemented, "user endpoint not yet implemented")
}
