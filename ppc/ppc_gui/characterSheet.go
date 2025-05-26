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
	// Create the field container once
	fieldsContainer := csh.CreateContainer()

	// Button to add a new field
	addNewFieldButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		csh.AddField()
		fieldsContainer.Refresh() // Refresh the UI
		println("Add Field")
	})

	// Make the fields scrollable
	scrollContent := container.NewScroll(fieldsContainer)
	content := container.NewBorder(nil,
		container.NewVBox(widget.NewSeparator(), container.NewCenter(addNewFieldButton)),
		nil, nil,
		scrollContent,
	)

	// Replace main container content
	csh.mainContainer.Objects = []fyne.CanvasObject{content}
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
