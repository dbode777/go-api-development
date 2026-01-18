package question4

import "fmt"

/*
In Go, write a function named updateInventory that accepts a pointer to a Store struct and a Transaction struct.

The function should update the inventory of the store based on the transaction details.
If the quantity in the Transaction struct is positive, it means items are added to the inventory.
If it's negative, items are removed from the inventory.

If the transaction would cause the quantity of an item in the inventory to become negative,
the function should not update the inventory and should return false.

Otherwise, the function should update the inventory and return true.
If an item does not exist in the store's inventory, the function should treat it as if its quantity is 0.

So, for example, if a transaction tries to remove an item that does not exist in the inventory,
the function should not update the inventory and should return false.
Similarly, if a transaction adds an item that does not exist in the inventory,
the function should add it to the inventory with the correct quantity and should return true.
The function signature in Go is: func updateInventory(s *Store, t Transaction) bool
*/

type Store struct {
	Name      string
	Inventory map[string]int // map item names to the quantity in stock
}

type Transaction struct {
	ItemName string
	Quantity int // positive means adding stock, negative means removing stock
}

func updateInventory(s *Store, t Transaction) bool {
	if _, itemExists := s.Inventory[t.ItemName]; !itemExists {
		fmt.Println("The item does not exist in the inventory. Adding it to the inventory.")
		if t.Quantity < 0 {
			fmt.Println("The quantity provided in the transaction is negative. Setting the inventory quantity to 0.")
			s.Inventory[t.ItemName] = 0
			return true
		}

		fmt.Println("Setting the inventory quantity to the quantity provided in the transaction.")
		s.Inventory[t.ItemName] = t.Quantity
		return true
	} else if s.Inventory[t.ItemName]+t.Quantity < 0 {
		fmt.Println("The item exists in the inventory, but the quantity provided in the transaction would result in a negative quantity. Not updating the inventory.")
		return false
	}

	fmt.Println("The item exists in the inventory. Updating the inventory quantity.")
	s.Inventory[t.ItemName] += t.Quantity
	return true
}
