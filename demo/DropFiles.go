package main

import (
	"gform"
)

func mainWindow_OnDropFilesA(sender gform.Controller, arg *gform.DropFilesEventArg) {
	println("File Count:", len(arg.Files))
	println("Pos", arg.X, arg.Y)
	for _, f := range arg.Files {
		println(f)
	}
}


func main() {
    gform.Init()

    mainWindow := gform.NewForm(nil)
    mainWindow.SetPos(300, 100)
    mainWindow.SetSize(500, 300)
    mainWindow.SetSizable(false)
    mainWindow.SetMinButtonEnabled(false)
    mainWindow.SetMaxButtonEnabled(false)
    mainWindow.SetCaption("Drop Files Demo")
    mainWindow.SetDragAcceptFilesEnabled(true)

    mainWindow.OnDropFilesA().Attach(mainWindow_OnDropFilesA)

    mainWindow.Show()

    gform.RunMainLoop()
}