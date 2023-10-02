package main

import "fmt"

type Node struct {
    data int
    next *Node
}

type LinkedList struct {
    head *Node
}

// Insert a new node at the beginning of the linked list
func (ll *LinkedList) insertAtBeginning(data int) {
    newNode := &Node{data: data, next: ll.head}
    ll.head = newNode
}

// Insert a new node at the end of the linked list
func (ll *LinkedList) insertAtEnd(data int) {
    newNode := &Node{data: data, next: nil}

    if ll.head == nil {
        ll.head = newNode
        return
    }

    current := ll.head
    for current.next != nil {
        current = current.next
    }

    current.next = newNode
}

// Print the linked list
func (ll *LinkedList) printList() {
    current := ll.head
    for current != nil {
        fmt.Printf("%d ", current.data)
        current = current.next
    }
    fmt.Println()
}

func main() {
    // Create an empty linked list
    list := LinkedList{}

    // Insert nodes at the beginning
    list.insertAtBeginning(3)
    list.insertAtBeginning(2)
    list.insertAtBeginning(1)

    // Print the linked list: 1 -> 2 -> 3
    list.printList()

    // Insert nodes at the end
    list.insertAtEnd(4)
    list.insertAtEnd(5)

    // Print the linked list: 1 -> 2 -> 3 -> 4 -> 5
    list.printList()
}
