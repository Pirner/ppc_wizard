package ppc_gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Sidebar struct {
	homeBtn     *widget.Button
	createBtn   *widget.Button
	openBtn     *widget.Button
	saveBtn     *widget.Button
	mainContent *fyne.Container
	activeCSH   CharacterSheet
}

func (sb *Sidebar) CreateContainer() *fyne.Container {
	vboxSidebar := container.NewVBox(
		sb.homeBtn,
		widget.NewSeparator(),
		sb.createBtn,
		widget.NewSeparator(),
		sb.openBtn,
		widget.NewSeparator(),
		sb.saveBtn,
	)
	separator := canvas.NewLine(color.Gray{Y: 80})
	separator.StrokeWidth = 1
	sidebarContainer := container.NewHBox(vboxSidebar, separator)
	return sidebarContainer
}

func NewSidebar(mainContent *fyne.Container, activeCSH CharacterSheet, a fyne.App) *Sidebar {
	homeBtn := widget.NewButtonWithIcon("", theme.HomeIcon(), func() {
		// Handle Home action
		mainContent.Objects = NewHomeContent().Objects
		println("Home clicked")
	})
	createNewBtn := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		// Handle New action
		csh := NewCharacterSheet(mainContent, a)
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

		mainContent.Objects = []fyne.CanvasObject{content}
		mainContent.Refresh()
		println("Create new clicked")
	})
	openBtn := widget.NewButtonWithIcon("", theme.FolderOpenIcon(), func() {
		// Handle Open action
		println("Open clicked")
	})
	saveBtn := widget.NewButtonWithIcon("", theme.DocumentSaveIcon(), func() {
		// Handle Save action
		println("Save clicked")
	})
	sidebar := Sidebar{
		homeBtn,
		createNewBtn,
		openBtn,
		saveBtn,
		mainContent,
		activeCSH,
	}
	return &sidebar
}
