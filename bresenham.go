package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 500, 500))
	xb, yb := 0, 0
	for i := 0; i < 30; i++ {
		col := color.RGBA{uint8(rand.Intn(256)), uint8(rand.Intn(256)), uint8(rand.Intn(256)), 255}
		x, y := rand.Intn(500), rand.Intn(500)
		go drawLine(xb, yb, x, y, img, col)
		xb, yb = x, y
	}
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

func drawLine(x0, y0, x1, y1 int, img *image.RGBA, col color.RGBA) {
	//check if the line is vertical or horizontal
	if x0-x1 == 0 {
		drawVLine(x0, y0, y1, img)
		return
	}
	if y0-y1 == 0 {
		drawHLine(y0, x0, x1, img)
	}

	if abs(x1-x0) > abs(y1-y0) {
		//dx > dy, vary x so there are no gaps
		if x0 > x1 {
			x0, x1 = x1, x0
			y0, y1 = y1, y0
		}

		//check if the slope is negative or positive
		//if slope is negative, we should remove one
		//instead of adding one
		incY := 1
		if y0 > y1 {
			incY = -1
		}

		dx := x1 - x0      //dx > 0
		dy := abs(y1 - y0) //dy > 0
		y := y0
		e := -dx
		incE := dy * 2
		decE := -2 * dx //avoiding the - operation
		//for each turn in the loop

		for x := x0; x <= x1; x++ {
			img.Set(x, y, col)
			e += incE
			if e > 0 {
				y += incY
				e += decE
			}
		}
	} else {
		//dy > dx, vary y so there are no gaps
		if y0 > y1 {
			y0, y1 = y1, y0
			x0, x1 = x1, x0
		}
		//now y0 < y1

		//check if we should increase or decrease
		//the x value
		incX := 1
		if x0 > x1 {
			incX = -1
		}

		dy := y1 - y0      //should be positive
		dx := abs(x1 - x0) //should be positive
		incE := 2 * dx
		decE := -2 * dy
		x := x0
		e := -dy

		for y := y0; y <= y1; y++ {
			img.Set(x, y, col)
			e += incE
			if e > 0 {
				e += decE
				x += incX
			}
		}
	}
}

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

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
