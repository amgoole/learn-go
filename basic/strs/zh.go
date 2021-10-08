package strs

import "fmt"

func Str(){
	s := "asdcvb中国"
	for index, ch := range s {
		fmt.Printf("%d-->%x", index, ch)
		fmt.Printf(" %s  ", string(ch))
	}
	fmt.Println()
	for _, ch := range []byte(s) {
		fmt.Printf("%x  ", ch)
	}
	fmt.Println()
	for _, ch := range []rune(s) {
		fmt.Printf("%x", ch)
		fmt.Printf(" %s  ", string(ch))
	}

	fmt.Println()
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i])
		fmt.Printf(" %s  ", string(s[i]))

	}
}