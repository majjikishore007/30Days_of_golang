package main

import "testing"

func TestPerimeter(t*testing.T){
	rectangle:=Rectangle{10.0,10.0}
	got := Perimeter(rectangle)
	want:=40.0
	if got != want{
		t.Errorf(" got %.2f , want %.2f",got,want)
	} 
}

func TestArea(t *testing.T){
	// helper function helps in decoupling the concrete types 
	checkArea:= func (t*testing.T, shape Shape,want float64)  {
		t.Helper()
		got:=shape.Area()
		if got != want {
			t.Errorf("got %g want %g",want,got)
		}
	}
	t.Run("rectangles", func(t *testing.T){
		rectangle:=Rectangle{12.0,6.0}
		checkArea(t,rectangle,72.0)
		
	})
	t.Run("circles",func(t *testing.T) {
		circle:=Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
	// regular method where the logic is not decoupled 
	// t.Run("rectangel",func(t *testing.T) {
	// 	rectangle:=Rectangle{12.0,6.0}
	// 	got:=rectangle.Area()
	// 	want := 72.0
	// 	if got != want{
	// 		t.Errorf("got %g want %g",want,got)
	// 	}
	// })
	
	//table driven tests 
	// anonymous struct 
	areaTests:=[]struct{
		shape Shape
		want float64
	}{
		{Rectangle{12,6},72.0},
		{Circle{10},314.1592653589793},
		{Triangle{5,6},15.0},
	}
	for _,tt:=range areaTests {
		got := tt.shape.Area()
		if got != tt.want{
			//The %#v format string will print out our struct with the values in its field
			  t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
	} 

}

