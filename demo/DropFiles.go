package main

import (
	"gform"
)

func mainWindow_OnDropFiles(arg *gform.EventArg) {
    if e, ok := arg.Data().(*gform.DropFilesEventData); ok {
    	println("File Count:", len(e.Files))
    	println("Pos", e.X, e.Y)
    	for _, f := range e.Files {
    		println(f)
    	}
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

    mainWindow.OnDropFiles().Attach(mainWindow_OnDropFiles)

    mainWindow.Show()

    gform.RunMainLoop()
}