package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type Node[T any] struct {
	next  *Node[T]
	value T
}

type LinkedList[T any] struct {
	head, tail *Node[T]
}

func (lst *LinkedList[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &Node[T]{value: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &Node[T]{value: v}
		lst.tail = lst.tail.next
	}
}

// Get all items from the LinkedList
func (lst *LinkedList[T]) GetAll() []T {
	var items []T
	for item := lst.head; item != nil; item = item.next {
		items = append(items, item.value)
	}
	return items
}

func (lst *LinkedList[T]) Length() int {
	var i int
	for node := lst.head; node != nil; node = node.next {
		i += 1
	}
	return i
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	fmt.Println("keys:", MapKeys(m))

	_ = MapKeys[int, string](m)

	StrList := LinkedList[string]{}

	StrList.GetAll()

	lst := LinkedList[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}
