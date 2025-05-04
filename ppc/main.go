package main

import (
	"fmt"
	"ppc_wizard/ppc/ppc_gui"
)

func main() {
	fmt.Println("Hello World")
	ppcApp := ppc_gui.NewPPWizard()
	fmt.Println(ppcApp.MyApp)
	ppcApp.MyWindow.ShowAndRun()
	//ppcwApp := app.New()
	//ppcwMainW := ppcwApp.NewWindow("Toolbar Widget")
	//ppcwMainW.Resize(fyne.NewSize(800, 600))
	//
	//newBtn := widget.NewButtonWithIcon("New", theme.DocumentCreateIcon(), func() {
	//	// Handle New action
	//	println("New clicked")
	//})
	//
	//openBtn := widget.NewButtonWithIcon("Open", theme.FolderOpenIcon(), func() {
	//	// Handle Open action
	//	println("Open clicked")
	//})
	//
	//saveBtn := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
	//	// Handle Save action
	//	println("Save clicked")
	//})
	//
	//sidebar := container.NewVBox(
	//	newBtn,
	//	widget.NewSeparator(),
	//	openBtn,
	//	widget.NewSeparator(),
	//	saveBtn,
	//)
	//content := container.New(
	//	layout.NewFormLayout(),
	//	sidebar,
	//	widget.NewLabel("Content2"),
	//)
	//ppcwMainW.SetContent(content)
	//ppcwMainW.ShowAndRun()

	//myApp := app.New()
	//myWindow := myApp.NewWindow("List Widget")
	//
	//list := widget.NewList(
	//	func() int {
	//		return len(data)
	//	},
	//	func() fyne.CanvasObject {
	//		return widget.NewLabel("template")
	//	},
	//	func(i widget.ListItemID, o fyne.CanvasObject) {
	//		o.(*widget.Label).SetText(data[i])
	//	})
	//
	//myWindow.SetContent(list)
	//myWindow.ShowAndRun()

	// Create Home Page

	//myApp := app.New()
	//myWindow := myApp.NewWindow("Toolbar Widget")
	//myWindow.Resize(fyne.NewSize(800, 600))
	//
	//toolbar := widget.NewToolbar(
	//	widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
	//		log.Println("New document")
	//	}),
	//	widget.NewToolbarSeparator(),
	//	widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
	//	widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
	//	widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
	//	widget.NewToolbarSpacer(),
	//	widget.NewToolbarAction(theme.HelpIcon(), func() {
	//		log.Println("Display help")
	//	}),
	//)
	//
	//// content := container.NewBorder(toolbar, nil, nil, nil, widget.NewLabel("Content"))
	//content := container.New(layout.NewGridLayoutWithColumns(2), toolbar, layout.NewSpacer())
	//myWindow.SetContent(content)
	//myWindow.ShowAndRun()
	// w.ShowAndRun()
}
