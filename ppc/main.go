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
}
