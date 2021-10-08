package inheritance

type CombineNode struct {
	InNode *Node
}

func (c *CombineNode) PrePrint() string {
	if c == nil || c.InNode == nil {
		return ""
	}
	c.InNode.Print() // 需要使用InNode调用Node的方法
	left := CombineNode{c.InNode.Left}
	left.PrePrint()
	right := CombineNode{c.InNode.Right}
	right.PrePrint()
	return ""
}
