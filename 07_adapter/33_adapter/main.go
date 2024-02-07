package main

import (
	"fmt"
	"strings"
)

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

// ↑↑↑ utility functions

// API is vector based

type Line struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

type VectorImage struct {
	Lines []Line
}

func NewRectangle(width int, height int) *VectorImage {
	width -= 1
	height -= 1
	return &VectorImage{
		[]Line{
			Line{0, 0, width, 0},
			Line{0, 0, 0, height},
			Line{width, 0, width, height},
			Line{0, height, width, height},
		},
	}
}

// ^^^ Above is the interface you are given

// vvvv Below is the interface we have

type Point struct {
	X int
	Y int
}

type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
	maxX, maxY := 0, 0
	points := owner.GetPoints()
	for _, pixel := range points {
		if pixel.X > maxX {
			maxX = pixel.X
		}
		if pixel.Y > maxY {
			maxY = pixel.Y
		}
	}
	maxX += 1
	maxY += 1

	// preallocate

	data := make([][]rune, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]rune, maxX)
		for j := range data[i] {
			data[i][j] = ' '
		}
	}

	for _, point := range points {
		data[point.Y][point.X] = '*'
	}

	b := strings.Builder{}
	for _, line := range data {
		b.WriteString(string(line))
		b.WriteRune('\n')
	}

	return b.String()
}

// solution: Adapter
type vectorToRasterAdapter struct {
	points []Point
}

func (a *vectorToRasterAdapter) addLine(line Line) {
	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			a.points = append(a.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			a.points = append(a.points, Point{x, top})
		}
	}

	fmt.Println("we have", len(a.points), "points")
}

func (v vectorToRasterAdapter) GetPoints() []Point {
	return v.points
}

func VectorToRaster(vi *VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}

	for _, line := range vi.Lines {
		adapter.addLine(line)
	}

	return adapter // as RasterImage
}

func main() {
	rc := NewRectangle(6, 4)
	a := VectorToRaster(rc)
	fmt.Println(DrawPoints(a))
}
