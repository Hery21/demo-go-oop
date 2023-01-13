package testing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRectangle_Area(t *testing.T) {
	type fields struct {
		width  int
		length int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Should return 10 when width is 2 and length is 5",
			fields: fields{
				width:  2,
				length: 5,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rect := &Rectangle{
				width:  tt.fields.width,
				length: tt.fields.length,
			}
			if got := rect.Area(); got != tt.want {
				t.Errorf("Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Perimeter(t *testing.T) {
	type fields struct {
		width  int
		length int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Should return is a rectangle when width is 5 and length is 5",
			fields: fields{
				width:  5,
				length: 5,
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rect := &Rectangle{
				width:  tt.fields.width,
				length: tt.fields.length,
			}
			if got := rect.Perimeter(); got != tt.want {
				t.Errorf("Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_PerimeterTableDriven(t *testing.T) {
	type fields struct {
		width  int
		length int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Should return is a rectangle when width is 5 and length is 5",
			fields: fields{
				width:  5,
				length: 5,
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rect := &Rectangle{
				width:  tt.fields.width,
				length: tt.fields.length,
			}
			if got := rect.Perimeter(); got != tt.want {
				t.Errorf("Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_PerimeterNonTableDriven(t *testing.T) {

	t.Run("Should return 14 when width is 2 and length is 3", func(t *testing.T) {
		rect := Rectangle{
			width:  2,
			length: 5,
		}
		expected := 14

		perimeter := rect.Perimeter()
		assert.Equal(t, expected, perimeter)
	})
}
