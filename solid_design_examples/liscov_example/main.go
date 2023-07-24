package main

import "fmt"

type Size interface {
	GetHeight() int
	SetHeight(height int)
	GetWidth() int
	SetWidth(width int)
}

type Rectangle struct {
	height int
	width  int
}

func (r *Rectangle) GetHeight() int {
	return r.height
}
func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

func (r *Rectangle) GetWidth() int {
	return r.width
}
func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

// böyle bir durumda gidip kare diye bir eleman oluşturursak bu sıkıntı çıkarır çünkü kare her ne kadar bir dörtgen olsa da
// özel şartları vardır ve  height ve witdth i aynıdır liscov böyle özel durumlarda çalışır.
func UseThat(size Size) {
	width := size.GetWidth()
	size.SetHeight(10)
	expectedArea := 10 * width
	actualArea := size.GetHeight() * size.GetWidth()
	fmt.Print("Expected area: ", expectedArea, ", But actual is: ", actualArea)
}

func main() {

	rectangle := Rectangle{
		height: 5,
		width:  5,
	}
	UseThat(&rectangle)
}
