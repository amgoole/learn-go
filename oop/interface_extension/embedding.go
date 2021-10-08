package iexten

import (
	"fmt"
	str "oop/struction"
)

type Worker struct {
	*str.Person
	Work string
}

func (w *Worker) SetWork(work string){
	w.Work = work
}

func (w *Worker) SetName(name string){
	w.Name = name
}

func TestWorker()  {


	w := Worker{&str.Person{}, ""}
	w.SetAge(100)
	fmt.Println(w.Age)

	w.SetName("wenson")
	fmt.Println(w.Name)

	w.Person.SetName("nickname")
	fmt.Println(w.Name)
}
