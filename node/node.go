package node

import "fmt"

type Node struct {
	Id   int
	Next *Node
}

func (n *Node) Transverse() {
	fmt.Printf("Node #%d", n.Id)

	if n.Next == nil {
		return
	}

	fmt.Print(" -> ")
	n.Next.Transverse()

}
