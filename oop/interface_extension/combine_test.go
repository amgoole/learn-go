package iexten

import "testing"

func TestLogin(t *testing.T) {
	l := Login{"liming", "123456"}
	Connect(l)
}
