package main

import (
    "github.com/AllenDang/gform"
    "github.com/AllenDang/w32"
)

var (
    lb *gform.Label
)

func btn_onclick(arg *gform.EventArg) {
    println("Button clicked")
}

func btnOpenFile_onclick(arg *gform.EventArg) {
    file, accepted := gform.ShowOpenFileDlg(arg.Sender().Parent(), "Test open file dialog", "", 0, "")
    if accepted {
        lb.SetCaption(file)
    }
}

func btnBrowseFolder_onclick(arg *gform.EventArg) {
    folder, accepted := gform.ShowBrowseFolderDlg(arg.Sender().Parent(), "Test browse folder")
    if accepted {
        lb.SetCaption(folder)
    }
}

func btnSaveFile_onclick(arg *gform.EventArg) {
    file, accepted := gform.ShowSaveFileDlg(arg.Sender().Parent(), "Test save file dialog", "", 0, "")
    if accepted {
        lb.SetCaption(file)
    }
}

func btnMsgBox_onclick(arg *gform.EventArg) {
    gform.MsgBox(arg.Sender().Parent(), "Message", "Test messagebox from gform", w32.MB_OK|w32.MB_ICONINFORMATION)
}

func main() {
    gform.Init()

    mainWindow := gform.NewForm(nil)
    mainWindow.SetPos(300, 100)
    mainWindow.SetSize(500, 300)
    mainWindow.SetCaption("Controls Demo")

    btn := gform.NewPushButton(mainWindow)
    btn.SetPos(10, 10)
    btn.OnLBDown().Bind(btn_onclick)
    btn.OnLBUp().Bind(btn_onclick)
    btn.OnMBDown().Bind(btn_onclick)
    btn.OnMBUp().Bind(btn_onclick)
    btn.OnRBDown().Bind(btn_onclick)
    btn.OnRBUp().Bind(btn_onclick)

    tooltip := gform.NewToolTip(mainWindow)
    println(tooltip.AddTool(btn, "Hello world"))

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
    btnBrowseFolder.OnLBUp().Bind(btnBrowseFolder_onclick)

    btnOpenFile := gform.NewPushButton(gb1)
    btnOpenFile.SetPos(10, 50)
    btnOpenFile.SetCaption("Open File Dlg")
    btnOpenFile.OnLBUp().Bind(btnOpenFile_onclick)

    btnSaveFile := gform.NewPushButton(gb1)
    btnSaveFile.SetPos(10, 80)
    btnSaveFile.SetCaption("Save File Dlg")
    btnSaveFile.OnLBUp().Bind(btnSaveFile_onclick)

    btnMsgBox := gform.NewPushButton(gb1)
    btnMsgBox.SetPos(10, 110)
    btnMsgBox.SetCaption("Msgbox")
    btnMsgBox.OnLBUp().Bind(btnMsgBox_onclick)

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
