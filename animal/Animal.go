package animal

type Animal struct {
	Name string
}

func (a *Animal) Eat() {
	println("%v is eating", a.Name)
}
