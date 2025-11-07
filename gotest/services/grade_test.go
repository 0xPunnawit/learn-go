package services_test

import (
	"fmt"
	"gotest/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Writer block package เดียว
// Block block package คนละที่

func TestCheckGrade(t *testing.T) {

	type testCase struct {
		name     string
		score    int
		expected string
	}

	cases := []testCase{
		{name: "a", score: 80, expected: "A"},
		{name: "b", score: 70, expected: "B"},
		{name: "c", score: 60, expected: "C"},
		{name: "d", score: 50, expected: "D"},
		{name: "f", score: 40, expected: "F"},
	}

	for _, c := range cases {

		t.Run(c.name, func(t *testing.T) {
			grade := services.CheckGrade(c.score)
			// if grade != c.expected {
			// 	t.Errorf("Expected grade %s, got %s", c.expected, grade)
			// }

			assert.Equal(t, c.expected, grade)

		})
	}

	// Sub test ชื่อ Test
	// go test gotest/services -run="TestCheckGrade/Grade A" -v
	// t.Run("Grade A", func(t *testing.T) {
	// 	grade := services.CheckGrade(80)
	// 	expected := "A"
	// 	if grade != expected {
	// 		t.Errorf("Expected grade %s, got %s", expected, grade)
	// 	}
	// })

	// t.Run("B", func(t *testing.T) {
	// 	grade := services.CheckGrade(70)
	// 	expected := "B"
	// 	if grade != expected {
	// 		t.Errorf("Expected grade %s, got %s", expected, grade)
	// 	}
	// })

	// t.Run("C", func(t *testing.T) {
	// 	grade := services.CheckGrade(60)
	// 	expected := "C"
	// 	if grade != expected {
	// 		t.Errorf("Expected grade %s, got %s", expected, grade)
	// 	}
	// })

	// t.Run("D", func(t *testing.T) {
	// 	grade := services.CheckGrade(50)
	// 	expected := "D"
	// 	if grade != expected {
	// 		t.Errorf("Expected grade %s, got %s", expected, grade)
	// 	}
	// })
	// t.Run("F", func(t *testing.T) {
	// 	grade := services.CheckGrade(40)
	// 	expected := "F"
	// 	if grade != expected {
	// 		t.Errorf("Expected grade %s, got %s", expected, grade)
	// 	}
	// })
}

// go test gotest/services -bench=. -benchmem
func BenchmarkCheckGrade(b *testing.B) {
	for i := 0; i < b.N; i++ {
		services.CheckGrade(80)
	}
}

func ExampleCheckGrade() {
	grade := services.CheckGrade(80)
	fmt.Println(grade)
	// Output: A

}
