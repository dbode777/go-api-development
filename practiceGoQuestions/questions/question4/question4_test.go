package question4

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func validation(result bool, expected bool, err error, test *testing.T) {
	if err != nil && !result {
		fmt.Printf("Test failed as expected with error: %s", err)
	} else if err != nil && expected {
		test.Errorf("Error occurred during testing: %v", err)
	}

	if result != expected {
		test.Fatalf("Result %v, did not equal %v", result, expected)
	}

	fmt.Printf("Test passed: %v, equaled %v\n", result, expected)
}

func TestAddNewItemWithPositiveQuantity(t *testing.T) {
	expected := true

	expectedStoreStock := &Store{
		Name: "Store 2",
		Inventory: map[string]int{
			"Item 1": 15,
			"Item 2": 25,
			"Item 3": 10,
		},
	}

	storeStock := &Store{
		Name: "Store 2",
		Inventory: map[string]int{
			"Item 1": 15,
			"Item 2": 25,
		},
	}

	transaction := Transaction{
		ItemName: "Item 3",
		Quantity: 10,
	}

	result := updateInventory(storeStock, transaction)
	validation(result, expected, nil, t)

	if diff := cmp.Diff(storeStock, expectedStoreStock); diff != "" {
		t.Fatalf("Diff: %s\n", diff)
	}

	fmt.Println("Test passed")
}

func TestRemoveItemsUnsuccessfullyDueToNegativeQuantity(t *testing.T) {
	expected := false

	expectedStoreStock := &Store{
		Name: "Store 3",
		Inventory: map[string]int{
			"Item 1": 10,
			"Item 2": 5,
		},
	}

	storeStock := &Store{
		Name: "Store 3",
		Inventory: map[string]int{
			"Item 1": 10,
			"Item 2": 5,
		},
	}

	transaction := Transaction{
		ItemName: "Item 2",
		Quantity: -10,
	}

	result := updateInventory(storeStock, transaction)
	validation(result, expected, nil, t)

	if diff := cmp.Diff(storeStock, expectedStoreStock); diff != "" {
		t.Fatalf("Diff: %s\n", diff)
	}

	fmt.Println("Test passed")
}

func TestRemoveItemsUnsuccessfullyDueToNegativeQuantity2(t *testing.T) {
	expected := false

	expectedStoreStock := &Store{
		Name: "Store 4",
		Inventory: map[string]int{
			"Item 1": 5,
		},
	}

	storeStock := &Store{
		Name: "Store 4",
		Inventory: map[string]int{
			"Item 1": 5,
		},
	}

	transaction := Transaction{
		ItemName: "Item 1",
		Quantity: -10,
	}

	result := updateInventory(storeStock, transaction)
	validation(result, expected, nil, t)

	if diff := cmp.Diff(storeStock, expectedStoreStock); diff != "" {
		t.Fatalf("Diff: %s\n", diff)
	}

	fmt.Println("Test passed")
}

func TestAddItemsSuccessfully(t *testing.T) {
	expected := true

	expectedStoreStock := &Store{
		Name: "Store 5",
		Inventory: map[string]int{
			"Item 1": 10,
			"Item 2": 25,
		},
	}

	storeStock := &Store{
		Name: "Store 5",
		Inventory: map[string]int{
			"Item 1": 10,
			"Item 2": 5,
		},
	}

	transaction := Transaction{
		ItemName: "Item 2",
		Quantity: 20,
	}

	result := updateInventory(storeStock, transaction)
	validation(result, expected, nil, t)

	if diff := cmp.Diff(storeStock, expectedStoreStock); diff != "" {
		t.Fatalf("Diff: %s\n", diff)
	}

	fmt.Println("Test passed")
}

func TestAddNewItemsSuccessfully2(t *testing.T) {
	expected := true

	expectedStoreStock := &Store{
		Name: "Store 6",
		Inventory: map[string]int{
			"Item 1": 5,
		},
	}

	storeStock := &Store{
		Name:      "Store 6",
		Inventory: map[string]int{},
	}

	transaction := Transaction{
		ItemName: "Item 1",
		Quantity: 5,
	}

	result := updateInventory(storeStock, transaction)
	validation(result, expected, nil, t)

	if diff := cmp.Diff(storeStock, expectedStoreStock); diff != "" {
		t.Fatalf("Diff: %s\n", diff)
	}

	fmt.Println("Test passed")
}

func TestAddNewItemsSuccessfullyWithNegativeQuantity(t *testing.T) {
	expected := true

	expectedStoreStock := &Store{
		Name: "Store 6",
		Inventory: map[string]int{
			"Item 1": 0,
		},
	}

	storeStock := &Store{
		Name:      "Store 6",
		Inventory: map[string]int{},
	}

	transaction := Transaction{
		ItemName: "Item 1",
		Quantity: -12,
	}

	result := updateInventory(storeStock, transaction)
	validation(result, expected, nil, t)

	if diff := cmp.Diff(storeStock, expectedStoreStock); diff != "" {
		t.Fatalf("Diff: %s\n", diff)
	}

	fmt.Println("Test passed")
}
