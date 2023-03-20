package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type Lattice struct {
	widget.BaseWidget
	jHaty, jHatx, iHaty, iHatx binding.Float
}

func CreateLattice(jHaty, jHatx, iHaty, iHatx binding.Float) *Lattice {
	lattice := &Lattice{widget.BaseWidget{}, jHaty, jHatx, iHaty, iHatx}
	lattice.ExtendBaseWidget(lattice)
	return lattice
}

func (l Lattice) CreateRenderer() fyne.WidgetRenderer {
	fmt.Println("Renderer created")
	renderer := &latticeRenderer{&l, nil}
	renderer.createImage()
	return renderer
}

type latticeRenderer struct {
	lattice *Lattice
	image   *canvas.Image
}

func (l *latticeRenderer) createImage() {
	fmt.Println("Refresh")

	img := CreateImage()

	//Calculate inverse 2x2 matrix
	jHaty, _ := l.lattice.jHaty.Get()
	jHatx, _ := l.lattice.jHatx.Get()
	iHaty, _ := l.lattice.iHaty.Get()
	iHatx, _ := l.lattice.iHatx.Get()

	var determinant = iHatx*jHaty - iHaty*jHatx
	iHatxInverse := jHaty / determinant
	iHatyInverse := -iHaty / determinant
	jHatxInverse := -jHatx / determinant
	jHatyInverse := iHatx / determinant

	//Calculate range of lattices that hit the viewable area
	minLattice := float64(-LATTICE_SIZE) / 2
	maxLattice := float64(LATTICE_SIZE) / 2
	topLefty := int(iHatyInverse*minLattice + jHatyInverse*maxLattice)
	topLeftx := int(iHatxInverse*minLattice + jHatxInverse*maxLattice)
	topRighty := int(iHatyInverse*maxLattice + jHatyInverse*maxLattice)
	topRightx := int(iHatxInverse*maxLattice + jHatxInverse*maxLattice)
	bottomLefty := int(iHatyInverse*minLattice + jHatyInverse*minLattice)
	bottomLeftx := int(iHatxInverse*minLattice + jHatxInverse*minLattice)
	bottomRighty := int(iHatyInverse*maxLattice + jHatyInverse*minLattice)
	bottomRightx := int(iHatxInverse*maxLattice + jHatxInverse*minLattice)

	fmt.Println("------------------------")
	fmt.Println("------------------------")
	fmt.Println(topLefty, topLeftx)
	fmt.Println(topRighty, topRightx)
	fmt.Println(bottomLefty, bottomLeftx)
	fmt.Println(bottomRighty, bottomRightx)
	fmt.Println("------------------------")

	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	max := func(i, j int) int {
		if i < j {
			return j
		}
		return i
	}

	minY := min(topLefty, min(topRighty, min(bottomLefty, bottomRighty)))
	minX := min(topLeftx, min(topRightx, min(bottomLeftx, bottomRightx)))
	maxY := max(topLefty, max(topRighty, max(bottomLefty, bottomRighty)))
	maxX := max(topLeftx, max(topRightx, max(bottomLeftx, bottomRightx)))

	fmt.Println(minY, minX)
	fmt.Println(maxY, maxX)
	fmt.Println("------------------------")
	fmt.Println("------------------------")

	for j := minY; j <= maxY; j++ {
		for i := minX; i <= maxX; i++ {
			translateY := iHaty*float64(i) + jHaty*float64(j)
			translateX := iHatx*float64(i) + jHatx*float64(j)
			img.Point(translateY, translateX)
		}
	}

	img.Arrow(jHaty, jHatx, color.RGBA{255, 0, 0, 255})
	img.Arrow(iHaty, iHatx, color.RGBA{0, 0, 255, 255})

	l.image = canvas.NewImageFromImage(img.img)
	l.image.Resize(fyne.NewSize(IMAGE_SIZE, IMAGE_SIZE))
	l.image.Move(fyne.NewPos(0, 0))
}

func (l *latticeRenderer) Layout(size fyne.Size) {
	fmt.Println("Layout")
}

func (l *latticeRenderer) MinSize() fyne.Size {
	return fyne.NewSize(IMAGE_SIZE, IMAGE_SIZE)
}

func (l *latticeRenderer) Refresh() {
	l.createImage()
}

func (l *latticeRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{l.image}
}

func (l *latticeRenderer) Destroy() {
	fmt.Println("Destroy")
}
