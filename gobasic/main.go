package main

import (
	"fmt"
	"gobasic/customer"
	"gobasic/user"
	"unicode/utf8"
)

// Struct (โครงสร้างข้อมูล) คล้ายๆ class ในภาษาอื่น มีได้แค่ field
type Person struct {
	Name string
	Age  int
}

func Hello(p Person) string {
	return "Hello " + p.Name
}

// Method (ฟังก์ชันที่ผูกกับ Struct)
func (p2 Person) Hello2() string {
	return "Hello from Method: " + p2.Name
}

func main() {

	// Slice (Array แบบไดนามิก) คือ ขนาดมันเปลี่ยนแปลงได้
	x := []int{1, 2, 3, 5, 8, 10}
	x = append(x, 4)
	y := len(x)
	z := x[1:3]

	fmt.Printf("%d\n", z) // [2 3]

	fmt.Println(x)         // [1 2 3 5 8 10 4]
	fmt.Printf("%#v\n", x) // []int{1, 2, 3, 5, 8, 10, 4}

	fmt.Println(y)         // 7
	fmt.Printf("%#v\n", y) // 7

	name := "ก"
	println(utf8.RuneCountInString(name)) // 1

	// Map
	countries := map[string]string{}
	countries["th"] = "Thailand"
	countries["jp"] = "Japan"
	countries["us"] = "United States"
	countries["de"] = "Germany"

	country, ok := countries["jp"]
	if ok {
		fmt.Println(country) // Japan
	} else {
		fmt.Println("Country not found")
	}

	// Loop
	fmt.Println("For range Loop")
	values := []int{10, 20, 30, 40, 50}

	// For Range Loop
	for _, v := range values {
		fmt.Println(v)
	}

	// For Loop
	for i := 0; i < len(values); i++ {
		fmt.Println(i, values[i])
	}

	// While Loop
	fmt.Println("While Loop")
	count := 0
	for count < len(values) {
		fmt.Println(values[count])
		count++
	}

	// Function
	fmt.Println("Function")
	c := sum()
	fmt.Println(c)

	fmt.Println("Function with Parameters and Multiple Return Values")
	number, character := sum2(15, 25)
	fmt.Println(number, character)

	// Anonymous Function
	// ก็คือ ไม่มีชื่อ function นั่นเอง ให้ aa เป็นตัวแทน
	fmt.Println("Anonymous Function")
	aa := func(a, b int) int {
		return a + b
	}

	bb := aa(10, 20)
	fmt.Println(bb)

	// คล้ายๆ interface oop
	cal(add)
	cal(sub)
	// Anonymous Function แบบไม่ต้องเก็บในตัวแปร
	cal(func(a, b int) int {
		return a * b
	})

	v := sumArray(1, 3, 5, 7, 9)
	fmt.Println(v)

	sumArray()

	ans := customer.Sum(3, 3)
	fmt.Println(ans)

	fmt.Println(customer.Hello())

	fmt.Println(user.Name)

	xx := 10
	yy := xx

	fmt.Println(&xx)
	fmt.Println(&yy)

	aaa := 30
	var bbb *int
	bbb = &aaa
	fmt.Println(aaa)  // value
	fmt.Println(&aaa) // address
	fmt.Println(bbb)  // address

	var ccc int
	ccc = 100

	var ddd *int
	ddd = &ccc
	*ddd = 200

	fmt.Println("\nPointer")
	fmt.Println("c:", ccc)
	fmt.Println("&ccc:", &ccc)
	fmt.Println("ddd: ", ddd)
	fmt.Println("*ddd: ", *ddd)
	fmt.Println("&ddd: ", &ddd)

	// ใช้ Pointer กับ Function
	fmt.Println("\n\nPointer with Function")
	var result int
	sumPointer(&result)
	fmt.Println("Result:", result)
	fmt.Println("&result:", &result)

	// class
	// Struct
	fmt.Println("\n\nStruct")
	classPerson := Person{"Ing", 22}
	fmt.Printf("%#v\n", classPerson) // main.Person{Name:"Ing", Age:22}
	// หรือ ใช้แบบนี้ก็ได้
	var classPerson2 = Person{
		Name: "Punnawit",
		Age:  25,
	}
	fmt.Printf("%#v\n", classPerson2) // main.Person{Name:"Punnawit", Age:25}

	classPerson3 := customer.Customer{
		Name: "Customer 1",
		// age:  30, ใช้ไม่ได้ เพราะเป็นตัวแปร private เป็นตัวเล็ก
	}
	fmt.Printf("%#v\n", classPerson3) // customer.Customer{Name:"Customer 1", age:0} zero value ของ int คือ 0

	fmt.Println("\nStruct with function")
	classPerson4 := Person{
		Name: "Alice",
		Age:  23,
	}
	println(Hello(classPerson4))

	// OOP
	fmt.Println("\nOOP")
	ggg := customer.ClassPerson{}
	ggg.SetName("OOP Customer")
	fmt.Printf("%#v\n", ggg.GetName())
}

// function sum that returns an integer
func sum() int {
	a := 10
	b := 20

	return a + b
}

// function sum2 that takes two integers as parameters and returns an integer and a string
// จะ return กี่ค่าก็ได้
func sum2(a, b int) (int, string) {
	return a + b, "Hello"
}

// f คือ ตัวแปรที่เก็บฟังก์ชัน
// func cal that takes a function as a parameter
// คล้ายๆ interface oop
func cal(f func(int, int) int) {
	sum := f(15, 40)
	fmt.Println(sum)
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

// slice แบบไม่จำกัดจำนวน
func sumArray(y ...int) int {
	s := 0
	for _, v := range y {
		s += v
	}
	return s
}

func sumPointer(result *int) {
	a := 12
	b := 22
	*result = a + b
	fmt.Println("Function &result:", &result)
	fmt.Println("Function *result:", *result)
}
