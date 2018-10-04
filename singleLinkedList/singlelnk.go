package singlelnk

import "fmt"

//Node list element
type Node struct {
	next *Node
	prev *Node
	Key  interface{}
}

// List type has two pointers Head, Tail
type List struct {
	Head *Node
	Tail *Node
}

// Insert function(Key interface{}); insert element at list top
func (L *List) Insert(Key interface{}) {
	node := &Node{
		next: L.Head,
		Key:  Key,
	}
	if L.Head != nil { // if L not empty
		L.Head.prev = node
	}
	L.Head = node // else node is the first element

	// traverse L and adjust Tail pointer
	l := L.Head
	for l.next != nil {
		l = l.next
	}
	L.Tail = l
}

// Display function(); display list contents
func (L *List) Display() {
	node := L.Head
	for node != nil {
		fmt.Printf("%+v -> ", node.Key)
		node = node.next
	}
	fmt.Println("---")
}

// Display function(node *Node); display nodes forward starting from node
func Display(node *Node) {
	for node != nil {
		fmt.Printf("%v ->", node.Key)
		node = node.next
	}
	fmt.Println("---")
}

// DisplayBwrds function(node *Node); display nodes in backward order starting from node
func DisplayBwrds(node *Node) {
	for node != nil {
		fmt.Printf("%v <- ", node.Key)
		node = node.prev
	}
	fmt.Println("---")
}

// Reverse function(); reverse list
func (L *List) Reverse() {
	curr := L.Head
	prev := &Node{}
	L.Tail = L.Head
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	L.Head = prev
	Display(L.Head)
}
