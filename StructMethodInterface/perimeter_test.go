package StructMethodInterface

import "testing"


func TestPrimeter(t *testing.T) {
	ractangle := Rectangle{10.0,10.0}
	got := Perimeter(ractangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

//func TestArea(t *testing.T) {
//	t.Run("test rangle", func(t *testing.T) {
//		ractangle := Rectangle{12.0,6.0}
//		got := ractangle.Area()
//		want := 72.0
//
//		if got != want {
//			t.Errorf("got %.2f want %.2f", got, want)
//		}
//	})
//	t.Run("test circles", func(t *testing.T) {
//		circle := Circle{10}
//		got := circle.Area()
//		want := 314.16
//
//		if got != want {
//			t.Errorf("got %.2f want %.2f", got, want)
//		}
//
//
//	})
//
//}

//func TestArea(t *testing.T) {
//	checkArea := func(t *testing.T, shape Shape, want float64) {
//		t.Helper()
//		got := shape.Area()
//		if got != want {
//			t.Errorf("got %.2f want %.2f", got, want)
//		}
//	}
//
//	t.Run("rectangel",func(t *testing.T){
//		rectangle := Rectangle{12,6}
//		checkArea(t,rectangle,72.0)
//	})
//	t.Run("circles",func(t *testing.T){
//		circle := Circle{10}
//		checkArea(t,circle,314.1592653589793)
//	})
//}


func TestArea(t *testing.T) {
	areaTests := []struct{
		shape Shape
		want float64
	}{
		{Rectangle{12,6},72.0},
		{Circle{10},314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.Area()
		if got != tt.want {
			t.Errorf("got %.2f want %.2f", got, tt.want)
		}
	}

}