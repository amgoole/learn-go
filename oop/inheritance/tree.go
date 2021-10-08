package inheritance

import (
	"fmt"
)

type Node struct {
	Val         interface{}
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Println(node.Val)
}


func CreateTree(list []int, layer ...*Node) *Node {
	var nodes []*Node
	for _, childNode := range layer {
		if len(list) >= 2 {
			if list[0] != -1{
				childNode.Left = &Node{Val: list[0]}
				nodes = append(nodes, childNode.Left)
			}
			if list[1] != -1{
				childNode.Right = &Node{Val: list[1]}
				nodes = append(nodes, childNode.Right)
			}
			list = list[2:]
		} else if len(list) == 1 {
			childNode.Left = &Node{Val: list[0]}
			nodes = append(nodes, childNode.Left)
			return childNode
		} else {
			return childNode
		}
	}
	CreateTree(list, nodes...)
	return nil
}
