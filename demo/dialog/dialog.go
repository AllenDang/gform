package main

import (
    "fmt"
    "syscall"
    "gform"
    "w32"
)

var (
    edt *gform.Edit
)

func onclick() {
    edt.SetCaption("Got you !!!")
}

func main() {
    gform.Init()

    dialog := gform.NewFormFromTemplate(nil, w32.MakeIntResource(101))
    dialog.Center()
    dialog.Show()

    edt = gform.AttachEdit(dialog, 1000)
    edt.SetCaption("Hello")

    btn := gform.AttachPushButton(dialog, 2)
    btn.OnLBDown().Attach(onclick)

    lv := gform.AttachListView(dialog, 1002)

    for i := 0; i < 10; i++ {
        var li w32.LVITEM
        li.Mask = w32.LVIF_TEXT
        li.IItem = int32(i)
        li.PszText = syscall.StringToUTF16Ptr(fmt.Sprintf("Here is item #%d", i))

        lv.InsertItem(&li)
    }

    gform.RunMainLoop()
}
