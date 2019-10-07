package npack

import (
	"fmt"
	"testing"
)

func TestSayMyName(t *testing.T) {
	var v string
	v = SayMyName("Andrey")
	if v != "My Name is Andrey" {
		t.Errorf("Expected My Name is Andrey got %v", v)
	}
}

func TestTableSayMyName(t *testing.T) {
	type test struct {
		name   string
		answer string
	}
	tests := []test{{
		name:   "Andrey",
		answer: "My Name is Andrey",
	}, {
		name:   "Bla",
		answer: "My Name is Bla",
	}}
	for _, v := range tests {
		i := SayMyName(v.name)
		if i != v.answer {
			t.Error("Error", i)
		}
	}
}

func ExampleSayMyName() {
	fmt.Println(SayMyName("Andrey"))
	// Output: My Name is Andrey
}
