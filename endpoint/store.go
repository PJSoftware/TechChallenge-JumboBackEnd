package endpoint

// GET
// ​/store​/inventory
// Returns pet inventories by status

// GET
// ​/store​/order​/{orderId}
// Find purchase order by ID

// DELETE
// ​/store​/order​/{orderId}
// Delete purchase order by ID

// POST
// ​/store​/order
// Place an order for a pet

func (r *Request) handleStore() {
	r.status(501, "store endpoint not yet implemented")
}
