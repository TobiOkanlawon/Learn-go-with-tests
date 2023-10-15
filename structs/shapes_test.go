package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{12.0, 12.0}
	got := rectangle.Perimeter()
	want := 48.0

	if got != want {
		t.Errorf("got %.2f, but want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("%#v got %.2f, but want %.2f", shape, got, want)
		}
	}

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectagle", shape: Rectangle{Width: 12.0, Height: 12.0}, want: 144.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.shape, tt.want)
		})
	}
}

// func assertFloatEquality()
