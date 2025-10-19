package customer

func Hello() string {
	return "Hello: " + name
}

type Customer struct {
	Name string
	age  int
}
