package inheritance

type EmbeddingNode struct {
	*Node
}

func (c *EmbeddingNode) PostPrint() string {
	if c == nil || c.Node == nil {
		return ""
	}
	left := EmbeddingNode{c.Node.Left}
	left.PostPrint()
	right := EmbeddingNode{c.Node.Right}
	right.PostPrint()
	c.Print()  // 可以直接调用Node的方法
	return ""
}
