package ppc_gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"io/ioutil"
	"path/filepath"
)

type PPWizard struct {
	MyApp       fyne.App
	MyWindow    fyne.Window
	Sidebar     Sidebar
	Header      fyne.Container
	Footer      fyne.Container
	mainContent fyne.Container
}

func LoadResourceFromPath(path string) (fyne.Resource, error) {
	bytes, err := ioutil.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	name := filepath.Base(path)
	staticRes := fyne.StaticResource{StaticName: name, StaticContent: bytes}
	return &staticRes, nil
}

func NewHeader() fyne.Container {
	separator := canvas.NewLine(color.Gray{Y: 80})
	separator.StrokeWidth = 1
	topCenter := container.NewCenter(canvas.NewText("Pen&Paper Agent", color.White))
	header := container.NewVBox(
		topCenter,
		separator,
	)
	return *header
}

func NewFooter() fyne.Container {
	separator := canvas.NewLine(color.Gray{Y: 80})
	separator.StrokeWidth = 1
	footerContent := container.NewCenter(canvas.NewText("1.0.0", color.White))
	footer := container.NewVBox(
		separator,
		footerContent,
	)
	return *footer
}

func NewHomeContent() *fyne.Container {
	homeHeaderText := "Welcome to the Pen&Paper Agent, this application supports people in playing Pen&Paper games"
	homeHeader := widget.NewLabel(homeHeaderText)
	homeHeader.Wrapping = fyne.TextWrapWord // optional: ensures word wrapping
	homeHeader.Alignment = fyne.TextAlignCenter

	homeHeaderContent := container.New(layout.NewVBoxLayout(), homeHeader)

	introductionText := "This agent allows to create character sheets with dynamic content that can be customized" +
		"and saved for further use. Every single character sheet can be derived from a template which can be reused." +
		" For final usage, the filled character template can be exported and then printed to be used in a party." +
		" All templates are stored as xml files."
	introductionLabel := widget.NewLabel(introductionText)
	introductionLabel.Wrapping = fyne.TextWrapWord
	introductionLabel.Alignment = fyne.TextAlignCenter

	sampleAttribute := NewNumericalAttribute("Intelligence", 14)
	sampleAttribute2 := NewNumericalAttribute("Stamina", 12)
	sampleAttribute3 := NewNumericalAttribute("Charisma", 12)

	homeContent := container.NewVBox(
		homeHeaderContent,
		widget.NewSeparator(),
		introductionLabel,
		sampleAttribute.CreateContainer(),
		sampleAttribute2.CreateContainer(),
		sampleAttribute3.CreateContainer(),
	)
	// content := container.NewCenter(mainContent)
	return homeContent
}

//func (ppw *PPWizard) UpdateMainContent(newContent fyne.Container) {
//	ppw.mainContent = newContent
//	content := container.NewBorder(
//		&ppw.Header,
//		&ppw.Footer,
//		&ppw.Sidebar,
//		nil,
//		&ppw.mainContent,
//	)
//	ppw.MyWindow.SetContent(content)
//}

func NewPPWizard() PPWizard {
	ppcwApp := app.New()

	ppcwMainW := ppcwApp.NewWindow("PenPaper Agent")
	ppcwMainW.Resize(fyne.NewSize(800, 600))
	r, _ := LoadResourceFromPath("../ressources/icon.png")
	ppcwMainW.SetIcon(r)

	header := NewHeader()
	mainContent := NewHomeContent()
	footer := NewFooter()
	sidebar := NewSidebar(mainContent)

	content := container.NewBorder(&header, &footer, sidebar.CreateContainer(), nil, mainContent)
	ppcwMainW.SetContent(content)

	//mainContent = *container.NewVBox(widget.NewLabel("something"))
	//mainContent.Refresh()
	sidebar.ChangeMainContent()

	ppwiz := PPWizard{
		ppcwApp,
		ppcwMainW,
		*sidebar,
		header,
		footer,
		*mainContent,
	}
	return ppwiz
}
