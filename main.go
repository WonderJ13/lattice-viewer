package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	window := a.NewWindow("Test :3")

	jHatyBind := binding.NewFloat()
	jHatxBind := binding.NewFloat()
	iHatyBind := binding.NewFloat()
	iHatxBind := binding.NewFloat()

	jHatyBind.Set(1)
	jHatxBind.Set(0)
	iHatyBind.Set(0)
	iHatxBind.Set(1)

	jHatyString := binding.FloatToString(jHatyBind)
	jHatxString := binding.FloatToString(jHatxBind)
	iHatyString := binding.FloatToString(iHatyBind)
	iHatxString := binding.FloatToString(iHatxBind)

	lattice := CreateLattice(jHatyBind, jHatxBind, iHatyBind, iHatxBind)

	updateLattice := binding.NewDataListener(func() {
		lattice.Refresh()
	})

	jHatyBind.AddListener(updateLattice)
	jHatxBind.AddListener(updateLattice)
	iHatyBind.AddListener(updateLattice)
	iHatxBind.AddListener(updateLattice)

	sliderJHaty := widget.NewSliderWithData(-10, 10, jHatyBind)
	sliderJHatx := widget.NewSliderWithData(-10, 10, jHatxBind)
	sliderIHaty := widget.NewSliderWithData(-10, 10, iHatyBind)
	sliderIHatx := widget.NewSliderWithData(-10, 10, iHatxBind)

	sliderJHaty.Step = 0.05
	sliderJHatx.Step = 0.05
	sliderIHaty.Step = 0.05
	sliderIHatx.Step = 0.05

	window.SetContent(
		container.New(
			layout.NewVBoxLayout(),
			container.New(
				layout.NewHBoxLayout(),
				layout.NewSpacer(),
				lattice,
				layout.NewSpacer(),
			),
			layout.NewSpacer(),
			container.New(
				layout.NewHBoxLayout(),
				container.New(
					layout.NewVBoxLayout(),
					widget.NewLabel("JHat y:\t\t\t"),
					widget.NewEntryWithData(jHatyString),
					sliderJHaty,
					widget.NewLabel("JHat x:\t\t\t"),
					widget.NewEntryWithData(jHatxString),
					sliderJHatx,
				),
				layout.NewSpacer(),
				container.New(
					layout.NewVBoxLayout(),
					widget.NewLabel("IHat y:\t\t\t"),
					widget.NewEntryWithData(iHatyString),
					sliderIHaty,
					widget.NewLabel("IHat x:\t\t\t"),
					widget.NewEntryWithData(iHatxString),
					sliderIHatx,
				),
			),
		),
	)
	window.ShowAndRun()
}
