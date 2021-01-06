package main

import "fmt"

/*
Use a queue structure to compute the height iteratively
*/
type queue struct {
	head, tail *qNode
	len        int
}

// qNode is a singly-linked list, so enqueue and dequeue ops are ~= push and pop
type qNode struct {
	v    interface{}
	next *qNode
}

func (q *queue) isEmpty() bool {
	if q.len > 0 {
		return false
	}
	return true
}

func (q *queue) dequeue() interface{} {
	if q.isEmpty() {
		return nil
	}
	n := q.head
	if q.len == 1 {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
	}
	q.len--

	return n.v
}

func (q *queue) enqueue(v interface{}) {
	new := &qNode{
		v:    v,
		next: nil,
	}
	if q.isEmpty() {
		q.head = new
		q.tail = new
	} else {
		q.tail.next = new
		q.tail = new
	}
	q.len++
}

/*
Queues operate on tree nodes
*/
type node struct {
	v        interface{}
	parent   *node
	children []*node
}

func newNode(v interface{}) *node {
	return &node{
		v: v,
	}
}

func (n *node) addParent(parent *node) {
	n.parent = parent
	parent.children = append(parent.children, n)
}

func (n *node) getRoot() *node {
	var res *node
	// walk up the tree
	for res = n; res.parent != nil; {
		res = res.parent
	}
	return res
}

func (n *node) debug() {
	fmt.Printf("{\n")
	fmt.Printf("\tval: %d\n", n.v)
	fmt.Printf("\tparent: %v\n", n.parent)
	fmt.Printf("\tchildren: %v\n", n.children)
	fmt.Printf("}\n")
}

func (n *node) getHeight() int {
	if n != nil {
		var height int
		q := &queue{}
		q.enqueue(n.getRoot())
		q.enqueue(nil)

		for !q.isEmpty() {
			cur := q.dequeue()
			if cur == nil {
				// level up
				height++
				if !q.isEmpty() {
					q.enqueue(nil)
				}
			} else {
				currentNode := cur.(*node)
				for _, c := range currentNode.children {
					q.enqueue(c)
				}
			}
		}
		return height
	}
	return 0
}

func main() {
	// line 1
	var n int
	fmt.Scan(&n)

	// line2
	vs := make([]int, n)
	for i := 0; i < n; i++ {
		var v int
		fmt.Scan(&v)
		vs[i] = v
	}

	nodes := make([]*node, 0)
	for _, v := range vs {
		nodes = append(nodes, newNode(v))
	}
	for child, parent := range vs {
		if parent >= 0 {
			// only try to add parent relationship if not at the root already
			nodes[child].addParent(nodes[parent])
		}
	}

	fmt.Println(nodes[0].getHeight())
}
