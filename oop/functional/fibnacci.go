package functional

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Fibonacci func() int

// Fib 反模式： 函数返回类型为受保护， 函数为非保护类型，可以运行，但不建议
func Fib() Fibonacci {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		if a > 1000 {
			panic(fmt.Sprintf("EOF"))
		}
		return a
	}
}

func (f Fibonacci) Read(p []byte) (n int, err error) {
	next := f()
	s := fmt.Sprintf("%d\n", next)
	// todo：incorrect if p is too small!
	return strings.NewReader(s).Read(p)
}

func PrintFile(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
