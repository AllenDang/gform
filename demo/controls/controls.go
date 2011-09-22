package main

import (
    "gform"
    "w32"
)

var (
    lb *gform.Label
)

func btn_onclick() {
    println("Button clicked")
}

func btnOpenFile_onclick(sender gform.Controller, arg *gform.MouseEventArg) {
    file, accepted := gform.ShowOpenFileDlg(sender.Parent(), "Test open file dialog", "", 0, "")
    if accepted {
        lb.SetCaption(file)
    }
}

func btnBrowseFolder_onclick(sender gform.Controller, arg *gform.MouseEventArg) {
    folder, accepted := gform.ShowBrowseFolderDlg(sender.Parent(), "Test browse folder")
    if accepted {
        lb.SetCaption(folder)
    }
}

func btnSaveFile_onclick(sender gform.Controller, arg *gform.MouseEventArg) {
    file, accepted := gform.ShowSaveFileDlg(sender.Parent(), "Test save file dialog", "", 0, "")
    if accepted {
        lb.SetCaption(file)
    }
}

func btnMsgBox_onclick(sender gform.Controller, arg *gform.MouseEventArg) {
    gform.MsgBox(sender.Parent(), "Message", "Test messagebox from gform", w32.MB_OK|w32.MB_ICONINFORMATION)
}

func main() {
    gform.Init()

    mainWindow := gform.NewForm(nil)
    mainWindow.SetPos(300, 100)
    mainWindow.SetSize(500, 300)
    mainWindow.SetCaption("Controls Demo")

    btn := gform.NewPushButton(mainWindow)
    btn.SetPos(10, 10)
    btn.OnLBDown().Attach(btn_onclick)
    btn.OnLBUp().Attach(btn_onclick)
    btn.OnMBDown().Attach(btn_onclick)
    btn.OnMBUp().Attach(btn_onclick)
    btn.OnRBDown().Attach(btn_onclick)
    btn.OnRBUp().Attach(btn_onclick)

    gb := gform.NewGroupBox(mainWindow)
    gb.SetCaption("GroupBox1")
    gb.SetSize(150, 100)
    gb.SetPos(10, 40)

    cb := gform.NewCheckBox(gb)
    cb.SetPos(10, 15)

    rb1 := gform.NewRadioButton(gb)
    rb1.SetPos(10, 40)

    rb2 := gform.NewRadioButton(gb)
    rb2.SetPos(10, 70)

    gb1 := gform.NewGroupBox(mainWindow)
    gb1.SetCaption("Dialogs")
    gb1.SetPos(240, 40)
    gb1.SetSize(150, 160)

    btnBrowseFolder := gform.NewPushButton(gb1)
    btnBrowseFolder.SetPos(10, 20)
    btnBrowseFolder.SetCaption("Browse Folder Dlg")
    btnBrowseFolder.OnLBUpA().Attach(btnBrowseFolder_onclick)

    btnOpenFile := gform.NewPushButton(gb1)
    btnOpenFile.SetPos(10, 50)
    btnOpenFile.SetCaption("Open File Dlg")
    btnOpenFile.OnLBUpA().Attach(btnOpenFile_onclick)

    btnSaveFile := gform.NewPushButton(gb1)
    btnSaveFile.SetPos(10, 80)
    btnSaveFile.SetCaption("Save File Dlg")
    btnSaveFile.OnLBUpA().Attach(btnSaveFile_onclick)

    btnMsgBox := gform.NewPushButton(gb1)
    btnMsgBox.SetPos(10, 110)
    btnMsgBox.SetCaption("Msgbox")
    btnMsgBox.OnLBUpA().Attach(btnMsgBox_onclick)

    lb = gform.NewLabel(mainWindow)
    lb.SetPos(130, 10)
    lb.SetSize(200, 25)

    edt := gform.NewEdit(mainWindow)
    edt.SetPos(10, 200)

    pb := gform.NewProgressBar(mainWindow)
    pb.SetPos(10, 225)
    pb.SetValue(50)

    mainWindow.Show()

    gform.RunMainLoop()
}
