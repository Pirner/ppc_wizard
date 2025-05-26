package ppc_gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	ppcgui "ppc_wizard/ppc/ppc_gui/fieldEntries"
)

type CharacterSheet struct {
	mainApp               fyne.App
	mainContainer         *fyne.Container
	addNewAttributeButton *widget.Button
	numericalAttributes   []NumericalAttribute
}

func NewCharacterSheet(mainContainer *fyne.Container, a fyne.App) CharacterSheet {
	addAttributeButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		// Handle Home action
		println("Add Attribute")
	})
	var numericAttributes []NumericalAttribute
	csh := CharacterSheet{
		a,
		mainContainer,
		addAttributeButton,
		numericAttributes,
	}
	return csh
}

func (csh CharacterSheet) AddField() {
	// create a new window and ask for what attribute to create
	entryWindow := csh.mainApp.NewWindow("Enter Values")
	entryWindow.Resize(fyne.NewSize(400, 300))
	headerLabel := widget.NewLabel("Create New Field")

	// Dropdown items
	options := []string{"NumericalAttribute", "Not Implemented"}

	// Create dropdown (select) widget
	selectWidget := widget.NewSelect(options, func(selected string) {
		// Callback when selection changes
		println("Selected:", selected)
		if selected == options[0] {
			entryContent := ppcgui.NewNumericalFieldEntry()

			content := container.NewVBox(
				container.NewCenter(headerLabel),
				entryContent.CreateContainer(),
			)
			entryWindow.SetContent(content)
			entryWindow.Show()
		}
	})
	selectWidget.PlaceHolder = "Choose an option"

	content := container.NewVBox(
		container.NewCenter(headerLabel),
		selectWidget,
	)
	entryWindow.SetContent(content)
	entryWindow.Show()
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
