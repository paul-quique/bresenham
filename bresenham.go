package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 500, 500))
	drawHLine(200, 0, 400, img)
	drawHLine(300, 200, 0, img)
	drawVLine(250, 50, 450, img)
	drawVLine(350, 400, 200, img)
	save(img, "test.png")
}

func save(img image.Image, fname string) {
	f, err := os.Create(fname)
	checkErr(err)

	png.Encode(f, img)
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

/*func drawLine(x0, y0, x1, y1 int, img *image.RGBA) {
	dx := math.Abs(x0)
}*/

func drawHLine(y, x0, x1 int, img *image.RGBA) {
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	for x := x0; x < x1; x++ {
		img.Set(x, y, color.RGBA{80, 80, 200, 255})
	}
}

func drawVLine(x, y0, y1 int, img *image.RGBA) {
	if y0 > y1 {
		y0, y1 = y1, y0
	}
	for y := y0; y < y1; y++ {
		img.Set(x, y, color.RGBA{80, 80, 200, 255})
	}
}
