package test2

import (
	"fmt"
	_ "fmt"
)

type Human struct {
	Age  int
	Name string
}

func (h *Human) Say() {
	fmt.Println("humam " + h.Name + " is say")
}

func (h *Student) Say() {
	fmt.Println("Student " + h.Name + " is say")
}

type Student struct {
	Human
	int
	Name  string
	Value string
	Score int
	Id    string
}

func (s *Student) Println() {
	println("********")
	println(s.Score)
	println(s.Value)
	println(s.Human.Name)
}
