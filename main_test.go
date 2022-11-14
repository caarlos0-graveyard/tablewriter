package tablewriter

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestRenderEmpty(t *testing.T) {
	if err := Render(os.Stderr, nil, []string{"Item"}, func(item string) ([]string, error) {
		return []string{"doesnt matter"}, nil
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
	}, []string{"Name", "Age"}, func(item Person) ([]string, error) {
		return []string{item.Name, strconv.Itoa(item.Age)}, nil
	}); err != nil {
		t.Fatal("render failed:", err)
	}
}

func TestRenderError(t *testing.T) {
	if err := Render(os.Stderr, []Person{
		{"Carlos", 32},
		{"John", 41},
		{"Jose", 28},
	}, []string{"Name", "Age"}, func(item Person) ([]string, error) {
		return nil, fmt.Errorf("foo bar")
	}); err == nil {
		t.Fatal("expected render to fail, got nil")
	}
}
