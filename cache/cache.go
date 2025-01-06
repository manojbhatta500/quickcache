package cache

import "fmt"

type Node struct {
	Start *Node
	Value string
	End   *Node
}

type LinkedList struct {
	Head     *Node
	Length   int
	Capacity int
}

func CreateNode(value string) Node {
	return Node{
		Value: value,
	}
}

func CreateLinkedList(cap int) LinkedList {
	return LinkedList{
		Capacity: cap,
	}
}

func (ll *LinkedList) IsEmpty() bool {
	if ll.Head == nil {
		return true
	} else {
		return false
	}

}

var CacheData = CreateLinkedList(10)

func (ll *LinkedList) Traverse() {
	res := ll.IsEmpty()
	if res {
		fmt.Println("linked list is empty")
		return
	}
	current := ll.Head
	var items int
	for current != nil {
		fmt.Printf("\n current node value is %s \n", current.Value)
		current = current.End
		items++
	}
	fmt.Println("\nthe length of linked list is ", items)
}
func (ll *LinkedList) Set(val string) {
	res := ll.IsEmpty()
	if res {
		newNode := CreateNode(val)
		ll.Head = &newNode
		ll.Length++
		return
	}
	currentNode := ll.Head
	var previousNode *Node
	for currentNode != nil {
		if currentNode.Value == val {
			if currentNode == ll.Head {
				return
			}
			if previousNode != nil {
				previousNode.End = currentNode.End
			}
			currentNode.End = ll.Head
			ll.Head = currentNode
			return
		}
		previousNode = currentNode
		currentNode = currentNode.End
	}
	newNode := CreateNode(val)
	newNode.End = ll.Head
	ll.Head = &newNode
	ll.Length++
}

func (ll *LinkedList) GetValues() []string {
	var values []string
	res := ll.IsEmpty()
	if res {
		return nil
	}
	current := ll.Head
	for current != nil {
		values = append(values, current.Value)
		current = current.End
	}
	return values
}
