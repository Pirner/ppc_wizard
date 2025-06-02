package ppc_gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type NumericalFieldEntry struct {
	value      int
	name       string
	ValueEntry *NumericalEntry
	NameEntry  *widget.Entry
}

func (nfe NumericalFieldEntry) CreateContainer() *fyne.Container {
	content := container.New(
		layout.NewGridLayout(2),
		nfe.NameEntry,
		nfe.ValueEntry,
	)
	return content
}

func NewNumericalFieldEntry() *NumericalFieldEntry {
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter Attribute Name")
	numEntry := NewNumericalEntry()
	numFieldEntry := NumericalFieldEntry{
		value:      0,
		name:       "",
		ValueEntry: numEntry,
		NameEntry:  input,
	}
	return &numFieldEntry
}
