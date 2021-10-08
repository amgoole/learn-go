package main

import "oop/inheritance"

func main() {
	// 组合继承调用
	//node := &inheritance.Node{Val: 0}
	//slice := []int{1, 2, 3, 4, 5, -1, 6, -1, 9, 8}
	//inheritance.CreateTree(slice, node)
	//myNode := inheritance.CombineNode{InNode: node}
	//myNode.PrePrint()

	// 内嵌继承
	node := &inheritance.Node{Val: 0}
	slice := []int{1, 2, 3, 4, 5, -1, 6, -1, 9, 8}
	inheritance.CreateTree(slice, node)
	myNode := inheritance.EmbeddingNode{Node: node}
	myNode.PostPrint()

}
