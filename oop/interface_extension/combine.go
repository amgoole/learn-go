package iexten

import "fmt"

type Retriever interface {
	Get(url string)
}

type Pusher interface {
	Post(url string, data map[string]interface{})
}

type Session interface {
	Retriever
	Pusher
}

type Login struct {
	Username string
	Password string
}

func (l Login) Post(url string, data map[string]interface{}) {
	fmt.Println(url)
	fmt.Println(data)
	//return []byte("Login Success"), nil
}

func (l Login) Get(url string) {
	fmt.Println(url)
	//return []byte("Fetching Success"), nil
}

func Connect(s Session) {
	s.Get("")
	s.Post("", map[string]interface{}{"kay": 1})
}


