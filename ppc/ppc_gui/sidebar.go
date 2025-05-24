package ppc_gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type Sidebar struct {
	homeBtn     *widget.Button
	createBtn   *widget.Button
	openBtn     *widget.Button
	saveBtn     *widget.Button
	mainContent *fyne.Container
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

func (sb *Sidebar) ChangeMainContent() {
	// sb.mainContent = container.NewVBox(widget.NewLabel("some text"))
	// sb.mainContent.Refresh()

	sb.mainContent.Objects = []fyne.CanvasObject{
		widget.NewLabel("New Content"),
		widget.NewButton("Another Button and I have changed it", nil),
	}
	sb.mainContent.Refresh()
}

func NewSidebar(mainContent *fyne.Container) *Sidebar {
	homeBtn := widget.NewButtonWithIcon("", theme.HomeIcon(), func() {
		// Handle New action
		println("Home clicked")
	})
	createNewBtn := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		// Handle Open action
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
	}
	return &sidebar
}
