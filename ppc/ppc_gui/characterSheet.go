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
	numericalAttributes   []NumericalAttribute
}

func NewCharacterSheet(mainContainer *fyne.Container) CharacterSheet {
	addAttributeButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		// Handle Home action
		println("Add Attribute")
	})
	var numericAttributes []NumericalAttribute
	csh := CharacterSheet{
		mainContainer,
		addAttributeButton,
		numericAttributes,
	}
	return csh
}

func (csh CharacterSheet) AddField() {
	numAttribute := NewNumericalAttribute("Intelligence", 15)
	csh.numericalAttributes = append(csh.numericalAttributes, *numAttribute)
	csh.RefreshMainContent()
}

func (csh CharacterSheet) RefreshMainContent() {
	addNewFieldButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		// Handle Home action
		csh.AddField()
		println("Add Field")
	})
	content := container.NewVBox(
		csh.CreateContainer(),
		widget.NewSeparator(),
		container.NewCenter(addNewFieldButton),
	)
	// mainContent.Objects = csh.CreateContainer().Objects
	csh.mainContainer.Objects = content.Objects
	csh.mainContainer.Refresh()
}

func (csh CharacterSheet) CreateContainer() *fyne.Container {
	var numAtts []fyne.CanvasObject
	for i, nA := range csh.numericalAttributes {
		println(i)
		numAtts = append(numAtts, nA.CreateContainer())
	}
	numAttsContainer := container.NewVBox(numAtts...)

	containerRes := container.NewVBox(
		widget.NewLabel("New CharacterSheet"),
		numAttsContainer,
	)
	csh.mainContainer = containerRes
	return csh.mainContainer
}
