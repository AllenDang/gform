package main

import (
    "fmt"
    "github.com/AllenDang/gform"
)

var (
    edt *gform.Edit
)

func onclick(arg *gform.EventArg) {
    edt.SetCaption("Got you !!!")
}

func main() {
    gform.Init()

    dialog := gform.NewDialogFromResId(nil, 101)
    dialog.Center()
    dialog.Show()

    edt = gform.AttachEdit(dialog, 1000)
    edt.SetCaption("Hello")

    btn := gform.AttachPushButton(dialog, 2)
    btn.OnLBDown().Bind(onclick)

    lv := gform.AttachListView(dialog, 1002)

    for i := 0; i < 10; i++ {
        lv.InsertItem(fmt.Sprintf("Here is item #%d", i), i)
    }

    gform.RunMainLoop()
}
