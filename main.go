package main

import (
	"fmt"

	"./dummydb"
)

func main() {
	dummydb.InitDB()
	fmt.Println(dummydb.PetByID(1))
	fmt.Println(dummydb.PetByID(2))
	fmt.Println(dummydb.PetByID(3))
	fmt.Println(dummydb.PetByID(4))
	fmt.Println(dummydb.PetByID(5))
	fmt.Println(dummydb.UserByUserName("alice"))
	fmt.Println(dummydb.UserByUserName("bob"))
	fmt.Println(dummydb.UserByUserName("charlie"))
	fmt.Println(dummydb.OrderByID(1))
	fmt.Println(dummydb.OrderByID(2))
}
