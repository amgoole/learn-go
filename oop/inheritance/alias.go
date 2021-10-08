package inheritance

import "fmt"

type AliasNode Node

//func (c *AliasNode) MidPrint() string {
//	if c == nil {
//		return ""
//	}
//	left := c.Left
//	left.MidPrint() // left为*Node数据类型， 无法调用Alias的MidPrint方法
//	right := c.Right
//	right.MidPrint()
//	c.Print()  // 无法调用Node的Print方法
//	return ""
//}

// 由于Alias与Node为严格的两种类型数据，方法均不可以互相调用
// 单独实现各自的方法,但内部字段类型可以互通

func (a *AliasNode) Print() {
	fmt.Println(a.Val)
}

func TestCalled() {
	a := AliasNode{}
	a.Print()

}
