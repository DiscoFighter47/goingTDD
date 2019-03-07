package structs

import "testing"

func TestArea(t *testing.T) {
	shapeTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{
			name: "Rectangle",
			shape: Rectangle{
				width:  10.0,
				height: 10.0,
			},
			want: 100.0,
		},
		{
			name: "Circle",
			shape: Circle{
				base:   20.0,
				height: 10.0,
			},
			want: 100.0,
		},
	}

	for _, tt := range shapeTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got '%.2f' want '%.2f'", tt.name, got, tt.want)
		}
	}
}
