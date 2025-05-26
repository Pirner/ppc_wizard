package ppc_gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	ppc_gui2 "ppc_wizard/ppc/ppc_gui/fieldEntries"
	"strconv"
)

type NumericalAttribute struct {
	value           *int
	labelWidget     *widget.Label
	input           *ppc_gui2.NumericalEntry
	buttonContainer *fyne.Container
	removeButton    *widget.Button
}

func (na *NumericalAttribute) CreateContainer() *fyne.Container {
	sampleInput := container.New(
		layout.NewGridLayout(4),
		na.labelWidget,
		na.input,
		na.buttonContainer,
		container.NewCenter(na.removeButton),
	)
	return sampleInput
}

func NewNumericalAttribute(attributeName string, defaultValue int) *NumericalAttribute {
	attributeWidget := widget.NewLabel(attributeName)
	intInput := ppc_gui2.NewNumericalEntry()
	intInput.SetPlaceHolder(attributeName)
	intInput.SetText(strconv.Itoa(defaultValue))
	// create up and down button
	upButton := widget.NewButtonWithIcon("", theme.MoveUpIcon(), func() {
		// Handle New action
		defaultValue++
		intInput.SetText(strconv.Itoa(defaultValue))
	})
	doButton := widget.NewButtonWithIcon("", theme.MoveDownIcon(), func() {
		// Handle New action
		defaultValue--
		intInput.SetText(strconv.Itoa(defaultValue))
	})

	removeButton := widget.NewButtonWithIcon("", theme.ContentClearIcon(), func() {
		println("remove this attribute")
	})

	buttons := container.NewHBox(upButton, doButton)
	buttonContainer := container.NewCenter(buttons)

	attributeComponent := &NumericalAttribute{}
	attributeComponent.value = &defaultValue
	attributeComponent.labelWidget = attributeWidget
	attributeComponent.input = intInput
	attributeComponent.buttonContainer = buttonContainer
	attributeComponent.removeButton = removeButton

	return attributeComponent
}
