package tablewriter

import (
	"os"
	"strconv"
	"testing"
)

func TestRenderEmpty(t *testing.T) {
	if err := Render(os.Stderr, nil, []string{"Item"}, func(item string) []string {
		return []string{"doesnt matter"}
	}); err != nil {
		t.Fatal("render failed:", err)
	}
}

type Person struct {
	Name string
	Age  int
}

func TestRender(t *testing.T) {
	if err := Render(os.Stderr, []Person{
		{"Carlos", 32},
		{"John", 41},
		{"Jose", 28},
	}, []string{"Name", "Age"}, func(item Person) []string {
		return []string{item.Name, strconv.Itoa(item.Age)}
	}); err != nil {
		t.Fatal("render failed:", err)
	}
}
