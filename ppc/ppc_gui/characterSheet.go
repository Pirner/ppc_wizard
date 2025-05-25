package ppc_gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type CharacterSheet struct {
	mainContainer         *fyne.Container
	addNewAttributeButton *widget.Button
}

func NewCharacterSheet(mainContainer *fyne.Container) CharacterSheet {
	addAttributeButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		// Handle Home action
		println("Add Attribute")
	})

	csh := CharacterSheet{
		mainContainer,
		addAttributeButton,
	}
	return csh
}

func (csh CharacterSheet) CreateContainer() *fyne.Container {
	sampleAttribute := NewNumericalAttribute("Intelligence", 14)

	containerRes := container.NewVBox(
		widget.NewLabel("New CharacterSheet"),
		sampleAttribute.CreateContainer(),
		widget.NewSeparator(),
		container.NewCenter(csh.addNewAttributeButton),
	)
	csh.mainContainer = containerRes
	return csh.mainContainer
}
