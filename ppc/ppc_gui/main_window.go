package ppc_gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"io/ioutil"
	"path/filepath"
)

type PPWizard struct {
	MyApp    fyne.App
	MyWindow fyne.Window
	Sidebar  fyne.Container
	Header   fyne.Container
	Footer   fyne.Container
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

func NewSidebar() fyne.Container {
	homeBtn := widget.NewButtonWithIcon("", theme.HomeIcon(), func() {
		// Handle New action
		println("Home clicked")
	})

	openBtn := widget.NewButtonWithIcon("", theme.FolderOpenIcon(), func() {
		// Handle Open action
		println("Open clicked")
	})

	saveBtn := widget.NewButtonWithIcon("", theme.DocumentSaveIcon(), func() {
		// Handle Save action
		println("Save clicked")
	})

	vboxSidebar := container.NewVBox(
		homeBtn,
		widget.NewSeparator(),
		openBtn,
		widget.NewSeparator(),
		saveBtn,
	)
	separator := canvas.NewLine(color.Gray{Y: 80})
	separator.StrokeWidth = 1
	sidebar := container.NewHBox(vboxSidebar, separator)
	return *sidebar
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

func NewHomeContent() fyne.Container {
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

	intInput := NewNumericalEntry()
	intInput.SetPlaceHolder("Intelligence")

	strInput := NewNumericalEntry()
	strInput.SetPlaceHolder("Strength")

	agiInput := NewNumericalEntry()
	agiInput.SetPlaceHolder("Agility")

	staInput := NewNumericalEntry()
	staInput.SetPlaceHolder("Stamina")
	sampleInput := container.New(
		layout.NewGridLayout(3),
		intInput,
		strInput,
		agiInput,
		staInput,
	)
	fmt.Println(sampleInput)
	// sampleInput := container.New(layout.NewCenterLayout(), inputs)
	sampleAttribute := NewNumericalAttribute("Sample Attribute 1", 1)
	sampleAttribute2 := NewNumericalAttribute("Sample Attribute 2", 2)

	homeContent := container.NewVBox(
		homeHeaderContent,
		widget.NewSeparator(),
		introductionLabel,
		sampleAttribute.CreateContainer(),
		sampleAttribute2.CreateContainer(),
	)
	// content := container.NewCenter(homeContent)
	return *homeContent
}

func NewPPWizard() PPWizard {
	ppcwApp := app.New()

	ppcwMainW := ppcwApp.NewWindow("PenPaper Agent")
	ppcwMainW.Resize(fyne.NewSize(800, 600))
	r, _ := LoadResourceFromPath("../ressources/icon.png")
	ppcwMainW.SetIcon(r)

	sidebar := NewSidebar()
	header := NewHeader()
	homeContent := NewHomeContent()
	footer := NewFooter()

	content := container.NewBorder(&header, &footer, &sidebar, nil, &homeContent)
	ppcwMainW.SetContent(content)

	ppwiz := PPWizard{
		ppcwApp,
		ppcwMainW,
		sidebar,
		header,
		footer,
	}
	return ppwiz
}
