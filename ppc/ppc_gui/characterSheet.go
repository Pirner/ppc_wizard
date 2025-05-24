package ppc_gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type CharacterSheet struct {
}

func NewCharacterSheet() *CharacterSheet {
	csh := &CharacterSheet{}
	return csh
}

func (csh *CharacterSheet) CreateContainer() *fyne.Container {
	containerRes := container.NewVBox()
	return containerRes
}
