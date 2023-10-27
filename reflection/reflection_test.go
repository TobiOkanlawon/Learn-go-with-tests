package main

import (
	"reflect"
	"testing"
)

type SpyWalk struct {
	Arguments []string
}

func (s *SpyWalk) fn(arg string) {
	s.Arguments = append(s.Arguments, arg)
}

type DummyInterface struct {
	SomeString      string
	SomeNumber      int
	SomeBoolean     bool
	SomeOtherString string
}

type NestedStruct struct {
	SomeString  string
	Innerstruct struct {
		Name string
		Age  int
	}
}

type TestCase struct {
	Data              interface{}
	Name              string
	ExpectedArguments []string
}

func TestWalk(t *testing.T) {

	t.Run("table tests", func(t *testing.T) {
		tests := []TestCase{
			TestCase{
				Name: "calls the right amount of times",
				Data: DummyInterface{
					SomeString:      "string",
					SomeNumber:      1,
					SomeBoolean:     true,
					SomeOtherString: "some_other_string",
				},
				ExpectedArguments: []string{
					"string", "some_other_string"},
			},
			TestCase{
				Name: "nested structs",
				Data: NestedStruct{
					SomeString: "string",
					Innerstruct: struct {
						Name string
						Age  int
					}{"John Doe", 15},
				},
				ExpectedArguments: []string{"string", "John Doe"},
			},
		}

		for _, test := range tests {
			spy := SpyWalk{}
			Walk(test.Data, spy.fn)

			if !reflect.DeepEqual(spy.Arguments, test.ExpectedArguments) {
				t.Errorf("expected %#v, but got %#v", test.ExpectedArguments, spy.Arguments)
			}
		}
	})
}
