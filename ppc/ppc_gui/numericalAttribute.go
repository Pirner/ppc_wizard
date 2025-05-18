package ppc_gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type NumericalAttribute struct {
	value          int
	labelWidget    *widget.Label
	input          *NumericalEntry
	placeholder    *widget.Label
	placeholderSep *widget.Label
}

func (na *NumericalAttribute) CreateContainer() *fyne.Container {
	sampleInput := container.New(
		layout.NewGridLayout(4),
		na.labelWidget,
		na.input,
		na.placeholder,
		na.placeholderSep,
	)
	return sampleInput
}

func NewNumericalAttribute(attributeName string, defaultValue int) *NumericalAttribute {
	attributeWidget := widget.NewLabel(attributeName)
	intInput := NewNumericalEntry()
	intInput.SetPlaceHolder(attributeName)
	sampleWidget1 := widget.NewLabel("1")
	sampleWidget2 := widget.NewLabel("2")

	attributeComponent := &NumericalAttribute{}
	attributeComponent.value = defaultValue
	attributeComponent.labelWidget = attributeWidget
	attributeComponent.input = intInput
	attributeComponent.placeholder = sampleWidget1
	attributeComponent.placeholderSep = sampleWidget2

	return attributeComponent
}
