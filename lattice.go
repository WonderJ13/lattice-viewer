package main

import (
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
	renderer := &latticeRenderer{&l, nil}
	renderer.createImage()
	return renderer
}

type latticeRenderer struct {
	lattice *Lattice
	image   *canvas.Image
}

func (l *latticeRenderer) createImage() {
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

	for j := minY; j <= maxY; j++ {
		for i := minX; i <= maxX; i++ {
			translateY := iHaty*float64(i) + jHaty*float64(j)
			translateX := iHatx*float64(i) + jHatx*float64(j)
			img.Point(translateY, translateX)
		}
	}

	img.Arrow(jHaty, jHatx, RED)
	img.Arrow(iHaty, iHatx, BLUE)

	l.image = canvas.NewImageFromImage(img.img)
	l.image.Resize(fyne.NewSize(IMAGE_SIZE, IMAGE_SIZE))
	l.image.Move(fyne.NewPos(0, 0))
}

func (l *latticeRenderer) Layout(size fyne.Size) {
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
}
