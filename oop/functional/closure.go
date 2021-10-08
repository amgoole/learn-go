package functional

import (
	"fmt"
	"strings"
)

type IAdd func(int) (int, IAdd)

func (I IAdd) Read(p []byte) (n int, err error) {
	i, _ := I(0)
	return strings.NewReader(fmt.Sprintf("%d\n", i)).Read(p)
}

func Operation(base int) IAdd {
	return func(step int) (int, IAdd) {
		return base + step, Operation(base + step)
	}
}

