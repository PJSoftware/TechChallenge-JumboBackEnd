package dummydb

import "time"

// Order contains our order information
//
// I feel to be actually functional, above and beyond the specification,
// Order would need to contain UserID so we know who placed the order
//
// Additionally, "Quantity" seems to be an odd value for a pet store;
// we are not selling "three dogs", we are selling pets with individual
// identifications. Given that somebody may choose to buy more than one
// pet in an order, this would need to have the ability to include multiple
// PetID entries per order.
//
// All that said, I've stuck with the model as specified for the purposes
// of this task.
type Order struct {
	ID       int64  `json:"id"`
	PetID    int64  `json:"petId"`
	Quantity int32  `json:"quantity"`
	ShipDate string `json:"shipDate"`
	Status   string `json:"status"`
	Complete bool   `json:"complete"`
}

type orderStatus struct {
	Placed    string
	Approved  string
	Delivered string
}

// OrderStatus used as enum
var OrderStatus = orderStatus{Placed: "placed", Approved: "approved", Delivered: "delivered"}

var orderID int64
var orderTBL []*Order

// NewOrder adds a new order to the db, returns the ID
func NewOrder(petID int64) int64 {
	orderID++
	id := orderID
	o := new(Order)
	o.ID = id
	o.PetID = petID
	o.Quantity = 1
	o.Status = OrderStatus.Placed
	PetByID(petID).SetStatus(PetStatus.Pending)
	orderTBL = append(orderTBL, o)
	return id
}

// OrderByID returns pointer to Order with specified ID
func OrderByID(id int64) *Order {
	for _, order := range orderTBL {
		if order.ID == id {
			return order
		}
	}
	return nil
}

// Shipped changes the order status to shipped
func (o *Order) Shipped() {
	o.Status = OrderStatus.Delivered
	o.ShipDate = time.Now().Format(time.RFC1123)
	o.Complete = true
	PetByID(o.PetID).SetStatus(PetStatus.Sold)
}
