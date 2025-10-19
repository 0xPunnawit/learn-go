package customer

var name = "Ing customer"

func Sum(a, b int) int {
	return a + b
}

type ClassPerson struct {
	name string
	age  int
}

func (c ClassPerson) GetName() string {
	return c.name
}

func (c *ClassPerson) SetName(name string) {
	c.name = name
}

func (c ClassPerson) GetAge() int {
	return c.age
}
