package main

import (
    "github.com/AllenDang/gform"
)

const (
    IDI_ICON1   = 100
    IDD_DIALOG1 = 101
)

func mainform_OnLoad(arg *gform.EventArg) {
    if dlg, ok := arg.Sender().(*gform.Dialog); ok {
        if ico, err := gform.NewIconFromResource(gform.GetAppInstance(), IDI_ICON1); err == nil {
            dlg.SetIcon(0, ico)
        }
    }
}

func main() {
    gform.Init()

    mainform := gform.NewDialogFromResId(nil, IDD_DIALOG1)
    mainform.OnLoad().Attach(mainform_OnLoad)
    mainform.Center()
    mainform.Show()

    gform.RunMainLoop()
}
