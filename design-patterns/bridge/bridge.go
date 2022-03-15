package main

import "fmt"

// IDrawShape interface
type IDrawShape interface {
	drawShape(x [5]float32, y [5]float32)
}

// DrawShape struct
type DrawShape struct{}

// drawShape method
func (drawShape DrawShape) drawShape(x [5]float32, y [5]float32) {
	fmt.Println("Drawing shape")
}

// IContour interface
type IContour interface {
	drawShape(x [5]float32, y [5]float32)
	resizeByFactor(factor int)
}

// DrawContour struct
type DrawContour struct {
	x      [5]float32
	y      [5]float32
	shape  DrawShape
	factor int
}

// drawContour method
func (contour DrawContour) drawContour(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Contour")
	contour.shape.drawShape(contour.x, contour.y)
}

// resizeByFactor method
func (contour DrawContour) resizeByFactor(factor int) {
	contour.factor = factor
}

func main() {
	x := [5]float32{1, 2, 3, 4, 5}
	y := [5]float32{1, 2, 3, 4, 5}
	contour := DrawContour{x, y, DrawShape{}, 2}
	contour.drawContour(x, y)
	contour.resizeByFactor(2)
}
