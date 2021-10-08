package struction

import "fmt"

type Person struct {
	Name  string
	Age   int
	Other [4]int
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func (p Person) SetAge(age int) {
	p.Age = age
}

func (p *Person) SetOther(other [4]int) {
	if p == nil {
		err := fmt.Errorf("%s", "nil pointer")
		fmt.Println(err.Error())
		return
	}
	p.Other = other
}

func (p *Person) GetName() string {
	return p.Name
}

type Teacher struct {
	Subject string
	Class   int
}

func (t Teacher) SetSubject(subject string) {
	t.Subject = subject
}

func (t Teacher) SetClass(class int) (teach Teacher) {
	t.Class = class
	teach = Teacher{"hello", 3}
	return teach
}
