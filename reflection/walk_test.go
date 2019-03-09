package reflection

import (
	"reflect"
	"sort"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name  string
		Input interface{}
		Want  []string
	}{
		{
			Name: "struct with one string field",
			Input: struct {
				Name string
			}{"Zahid"},
			Want: []string{"Zahid"},
		},
		{
			Name: "struct with two string field",
			Input: struct {
				Name string
				City string
			}{
				Name: "Zahid",
				City: "Dhaka",
			},
			Want: []string{"Zahid", "Dhaka"},
		},
		{
			Name: "struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{
				Name: "Zahid",
				Age:  24,
			},
			Want: []string{"Zahid"},
		},
		{
			Name:  "struct with no field",
			Input: struct{}{},
			Want:  []string{},
		},
		{
			Name: "nested field",
			Input: Person{
				Name: "Zahid",
				Profile: Profile{
					Age:  24,
					City: "Dhaka",
				},
			},
			Want: []string{"Zahid", "Dhaka"},
		},
		{
			Name: "pointers",
			Input: &Person{
				Name: "Zahid",
				Profile: Profile{
					Age:  24,
					City: "Dhaka",
				},
			},
			Want: []string{"Zahid", "Dhaka"},
		},
		{
			Name: "slices",
			Input: []Profile{
				{
					Age:  24,
					City: "Dhaka",
				},
				{
					Age:  30,
					City: "London",
				},
			},
			Want: []string{"Dhaka", "London"},
		},
		{
			Name: "arrays",
			Input: [2]Profile{
				{
					Age:  24,
					City: "Dhaka",
				},
				{
					Age:  30,
					City: "London",
				},
			},
			Want: []string{"Dhaka", "London"},
		},
		{
			Name: "maps",
			Input: map[string]string{
				"foo": "bar",
				"baz": "baz",
			},
			Want: []string{"bar", "baz"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := []string{}
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			assertCalls(t, got, test.Want)
		})
	}
}

func assertCalls(t *testing.T, got, want []string) {
	t.Helper()
	sort.Strings(got)
	sort.Strings(want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}
