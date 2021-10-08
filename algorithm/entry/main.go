package main

import (
	"algorithm/substr"
	"fmt"
)

func main() {
	length := substr.MaxSubStrNoRepeat("abcabcbb")
	fmt.Println(length)

	length = substr.MaxSubStrNoRepeatIntl("黑化肥挥发发灰会花飞灰化肥灰发发黑灰飞花")
	fmt.Println(length)
}
