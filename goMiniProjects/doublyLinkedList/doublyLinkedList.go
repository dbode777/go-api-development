package main

import "fmt"

type Node[T any] struct {
	next  *Node[T]
	prev  *Node[T]
	value T
}

type LinkedList[T any] struct {
	head   *Node[T]
	length uint
}

// Return the value at the given index in the linked list.
//
// If the index is out of range, it will return an undefined value matching the type used and false.
//
// This function does not modify the original linked list.
//
// idx - the index of the value to return.
//
// Returns the value at the given index and true if the index is in range, or an undefined value and false if the index is out of range.
func (ll *LinkedList[T]) index(idx uint) (T, bool) {
	var undefined T // used to return an undefined value matching the type used
	if idx >= ll.length {
		fmt.Println("Index out of range")
		return undefined, false
	}

	for range idx {
		ll.head = ll.head.next
	}

	return ll.head.value, true
}

// Append a new node with the given value to the end of the linked list.
func (ll *LinkedList[T]) append(value T) {
	if ll.length == 0 {
		ll.head = &Node[T]{}
		ll.head.value = value
		ll.head.prev = ll.head
		ll.head.next = ll.head
		ll.length++
		fmt.Printf("Appended %v to the linked list\n", value)
		return
	}

	tail := ll.head.prev
	node := &Node[T]{}
	node.value = value

	tail.next = node
	node.prev = tail
	node.next = ll.head
	ll.head.prev = node
	ll.length++
	fmt.Printf("Appended %v to the linked list\n", value)
}

// Prepend a new node with the given value to the start of the linked list.
//
// The new node will be inserted at the start of the linked list, and the
// existing head will be updated to point to the new node.
//
// value - the value of the new node to prepend.
//
// This function does not return a value, it modifies the original linked list.
//
// Example:
//
// list := &LinkedList[int]{}
// list.prepend(1)
// list.prepend(2)
// list.prepend(3)
//
// The linked list will now have the values 3, 2, 1 in order.
func (ll *LinkedList[T]) prepend(value T) {
	if ll.length == 0 {
		ll.head = &Node[T]{}
		ll.head.value = value
		ll.head.prev = ll.head
		ll.head.next = ll.head
		ll.length++
		fmt.Printf("Prepended %v to the linked list\n", value)
		return
	}

	tail := ll.head.prev
	node := &Node[T]{}
	node.value = value

	node.prev = tail
	node.next = ll.head
	tail.next = node
	ll.head.prev = node
	ll.head = node
	ll.length++
	fmt.Printf("Prepended %v to the linked list\n", value)
}

func (ll *LinkedList[T]) pop() {
	if ll.length == 0 {
		fmt.Println("Linked list is empty. Cannot pop")
		return
	}

	tail := ll.head.prev
	newTail := tail.prev

	newTail.next = &Node[T]{}
	ll.head.prev = newTail
	ll.length--

	fmt.Println("Removed the last node in the linked list")
}

func (ll *LinkedList[T]) shift() {
	if ll.length == 0 {
		fmt.Println("Linked list is empty. Cannot shift")
		return
	}

	ll.head = ll.head.next
	ll.length--
	fmt.Println("Removed the first node in the linked list")
}

// Prints out all the values in the linked list in order.
//
// This function iterates over the linked list and prints out all the values.
//
// Example:
//
// list := &LinkedList[int]{}
// list.append(1)
// list.append(2)
// list.append(3)
// list.printList()
//
// The output will be:
// 1
// 2
// 3
func (ll *LinkedList[T]) printList() {
	for range ll.length {
		fmt.Println(ll.head.value)
		ll.head = ll.head.next
	}
}

func main() {
	ll := LinkedList[uint]{&Node[uint]{}, 0}
	ll.pop()
	ll.shift()
	for i := range uint(5) {
		ll.append(i)
	}
	ll.prepend(100)
	ll.printList()
	ll.shift()
	ll.pop()
	ll.printList()
}
