package main

import (
	"image"
	"image/color"
	"math"
)

const (
	IMAGE_SIZE   = 600
	LATTICE_SIZE = 50

	IMG_LATTICE_RATIO = IMAGE_SIZE / LATTICE_SIZE
)

var RED = color.RGBA{255, 0, 0, 255}

func Sign(y int) int {
	if y < 0 {
		return -1
	}
	return 1
}

type Image struct {
	img                *image.RGBA
	originY, originX   int
	sizeMinY, sizeMinX float64
	sizeMaxY, sizeMaxX float64
}

func CreateImage() *Image {
	var gray = color.RGBA{200, 200, 200, 255}

	img := image.NewRGBA(
		image.Rectangle{
			image.Point{0, 0},
			image.Point{IMAGE_SIZE, IMAGE_SIZE},
		},
	)

	for i := 0; i < IMAGE_SIZE; i++ {
		for j := 0; j < IMAGE_SIZE; j++ {
			img.Set(i, j, color.White)
		}
	}

	for j := 0; j < IMAGE_SIZE; j += IMG_LATTICE_RATIO {
		for i := 0; i < IMAGE_SIZE; i += 1 {
			img.Set(j, i, gray)
		}
	}

	for j := 0; j < IMAGE_SIZE; j += 1 {
		for i := 0; i < IMAGE_SIZE; i += IMG_LATTICE_RATIO {
			img.Set(j, i, gray)
		}
	}

	for i := 0; i < IMAGE_SIZE; i++ {
		img.Set(i, IMAGE_SIZE/2, color.Black)
	}

	for j := 0; j < IMAGE_SIZE; j++ {
		img.Set(IMAGE_SIZE/2, j, color.Black)
	}

	return &Image{
		img,
		IMAGE_SIZE / 2, IMAGE_SIZE / 2,
		-LATTICE_SIZE / 2, -LATTICE_SIZE / 2,
		LATTICE_SIZE / 2, LATTICE_SIZE / 2,
	}
}

func (i *Image) lineHigh(y0, x0, y1, x1 int, lineColor color.Color) {
	dx := x1 - x0
	dy := y1 - y0

	xi := 1
	if dx < 0 {
		xi = -1
		dx = -dx
	}

	D := (2 * dx) - dy

	x := x0
	for y := y0; y <= y1; y++ {
		i.img.Set(x, y, lineColor)
		if D > 0 {
			x += xi
			D += 2 * (dx - dy)
		} else {
			D += 2 * dx
		}
	}
}

func (i *Image) lineLow(y0, x0, y1, x1 int, lineColor color.Color) {
	dx := x1 - x0
	dy := y1 - y0

	yi := 1
	if dy < 0 {
		yi = -1
		dy = -dy
	}

	D := (2 * dy) - dx

	y := y0
	for x := x0; x <= x1; x++ {
		i.img.Set(x, y, lineColor)
		if D > 0 {
			y += yi
			D += 2 * (dy - dx)
		} else {
			D += 2 * dy
		}
	}
}

func (i *Image) Line(yStart, xStart, yEnd, xEnd int, lineColor color.Color) {
	dy := float64(yEnd) - float64(yStart)
	dx := float64(xEnd) - float64(xStart)

	//Start of Bresenham's Line Algorithm
	//https://en.wikipedia.org/wiki/Bresenham's_line_algorithm
	if math.Abs(dy) < math.Abs(dx) {
		if xStart > xEnd {
			i.lineLow(yEnd, xEnd, yStart, xStart, lineColor)
		} else {
			i.lineLow(yStart, xStart, yEnd, xEnd, lineColor)
		}
	} else {
		if yStart > yEnd {
			i.lineHigh(yEnd, xEnd, yStart, xStart, lineColor)
		} else {
			i.lineHigh(yStart, xStart, yEnd, xEnd, lineColor)
		}
	}
}

func (i *Image) Arrow(yLattice, xLattice float64, lineColor color.Color) {
	yEnd := i.originY - int(yLattice*IMG_LATTICE_RATIO)
	xEnd := i.originX + int(xLattice*IMG_LATTICE_RATIO)

	yStart, xStart := i.originY, i.originX

	i.Line(yStart, xStart, yEnd, xEnd, lineColor)

	//Get Angle of Line
	dy := yStart - yEnd
	dx := xStart - xEnd
	angle := math.Atan2(float64(dy), float64(dx))

	angle1 := angle + math.Pi/4
	if angle1 > 2*math.Pi {
		angle1 -= 2 * math.Pi
	}
	y1 := math.Sin(angle1)*(IMG_LATTICE_RATIO/2) + float64(yEnd)
	x1 := math.Cos(angle1)*(IMG_LATTICE_RATIO/2) + float64(xEnd)
	i.Line(yEnd, xEnd, int(y1), int(x1), lineColor)

	angle2 := angle - math.Pi/4
	if angle2 < 0 {
		angle2 += 2 * math.Pi
	}
	y2 := math.Sin(angle2)*(IMG_LATTICE_RATIO/2) + float64(yEnd)
	x2 := math.Cos(angle2)*(IMG_LATTICE_RATIO/2) + float64(xEnd)
	i.Line(yEnd, xEnd, int(y2), int(x2), lineColor)

}

func (i *Image) Point(yLattice, xLattice float64) {
	var pointColor = color.Black
	//var pointColor = color.RGBA{255, 0, 255, 255}

	/*	if yLattice < i.sizeMinY || yLattice > i.sizeMaxY ||
		xLattice < i.sizeMinX || xLattice > i.sizeMaxX {
		return
	}*/

	yPixel := i.originY - int(yLattice*IMG_LATTICE_RATIO)
	xPixel := i.originX + int(xLattice*IMG_LATTICE_RATIO)

	i.img.Set(xPixel, yPixel, pointColor)
	i.img.Set(xPixel, yPixel+1, pointColor)
	i.img.Set(xPixel+1, yPixel, pointColor)
	i.img.Set(xPixel+1, yPixel+1, pointColor)
}
