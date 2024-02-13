package main

import "fmt"

// A virtual proxy is the kind of proxy that pretends it's really there
// when it's not necessarily.
type image interface {
	Draw()
}

type Bitmap struct {
	filename string
}

func NewBitmap(filename string) *Bitmap {
	fmt.Println("Loading image from", filename)
	return &Bitmap{filename: filename}
}

func (b *Bitmap) Draw() {
	fmt.Println("Drawing image", b.filename)
}

func DrawImage(image image) {
	fmt.Println("About to draw the image")
	image.Draw()
	fmt.Println("Done drawing the image")
}

type LazyBitmap struct {
	filename string
	bitmap   *Bitmap
}

func (l *LazyBitmap) Draw() {
	if l.bitmap == nil {
		l.bitmap = NewBitmap(l.filename)
	}
	l.bitmap.Draw()
}

func NewLazyBitmap(filename string) *LazyBitmap {
	return &LazyBitmap{filename: filename}
}

func main() {
	bmp := NewBitmap("demo.png")
	DrawImage(bmp) // If don't draw the image, no point of loading the image
	fmt.Println("================================")

	bmp2 := NewLazyBitmap("demo.png")
	DrawImage(bmp2)
	fmt.Println("================================")
	DrawImage(bmp2) // Loading happens only once
}
