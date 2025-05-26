package ppc_gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type NumericalFieldEntry struct {
	value      int
	name       string
	valueEntry *NumericalEntry
	nameEntry  *widget.Entry
}

func (nfe NumericalFieldEntry) CreateContainer() *fyne.Container {
	content := container.NewHBox(
		nfe.nameEntry,
		nfe.valueEntry,
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
		valueEntry: numEntry,
		nameEntry:  input,
	}
	return &numFieldEntry
}
