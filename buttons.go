package gform

import (
    "github.com/AllenDang/w32"
    "github.com/AllenDang/w32/user32"
)

type Button struct {
    W32Control
}

func (this *Button) Checked() bool {
    result := user32.SendMessage(this.hwnd, w32.BM_GETCHECK, 0, 0)
    return result == w32.BST_CHECKED
}

func (this *Button) SetChecked(checked bool) {
    wparam := w32.BST_CHECKED
    if !checked {
        wparam = w32.BST_UNCHECKED
    }
    user32.SendMessage(this.hwnd, w32.BM_SETCHECK, uintptr(wparam), 0)
}

type PushButton struct {
    Button
}

func NewPushButton(parent Controller) *PushButton {
    pb := new(PushButton)
    pb.init(parent)

    pb.SetFont(DefaultFont)
    pb.SetCaption("Button")
    pb.SetSize(100, 25)

    return pb
}

func AttachPushButton(parent Controller, id int) *PushButton {
    pb := new(PushButton)
    pb.attach(parent, id)
    RegMsgHandler(pb)

    return pb
}

func (this *PushButton) init(parent Controller) {
    this.W32Control.init("BUTTON", parent, 0, w32.BS_PUSHBUTTON|w32.WS_TABSTOP|w32.WS_VISIBLE|w32.WS_CHILD)
    RegMsgHandler(this)
}

func (this *PushButton) WndProc(msg uint, wparam, lparam uintptr) uintptr {
    switch msg {
    case w32.BN_CLICKED:
        println("Clicked")
    case w32.WM_LBUTTONDOWN:
        user32.SetCapture(this.Handle())
    case w32.WM_LBUTTONUP:
        user32.ReleaseCapture()
    }

    return this.W32Control.WndProc(msg, wparam, lparam)
}

type CheckBox struct {
    Button
}

func NewCheckBox(parent Controller) *CheckBox {
    cb := new(CheckBox)
    cb.init(parent)

    cb.SetFont(DefaultFont)
    cb.SetCaption("CheckBox")
    cb.SetSize(100, 25)

    return cb
}

func AttachCheckBox(parent Controller, id int) *CheckBox {
    cb := new(CheckBox)
    cb.attach(parent, id)
    RegMsgHandler(cb)

    return cb
}

func (this *CheckBox) init(parent Controller) {
    this.W32Control.init("BUTTON", parent, 0, w32.WS_TABSTOP|w32.WS_VISIBLE|w32.WS_CHILD|w32.BS_AUTOCHECKBOX)
    RegMsgHandler(this)
}

type RadioButton struct {
    Button
}

func NewRadioButton(parent Controller) *RadioButton {
    rb := new(RadioButton)
    rb.init(parent)

    rb.SetFont(DefaultFont)
    rb.SetCaption("RadioButton")
    rb.SetSize(100, 25)

    return rb
}

func AttachRadioButton(parent Controller, id int) *RadioButton {
    rb := new(RadioButton)
    rb.attach(parent, id)
    RegMsgHandler(rb)

    return rb
}

func (this *RadioButton) init(parent Controller) {
    this.W32Control.init("BUTTON", parent, 0, w32.WS_TABSTOP|w32.WS_VISIBLE|w32.WS_CHILD|w32.BS_AUTORADIOBUTTON)
    RegMsgHandler(this)
}

type GroupBox struct {
    Button
}

func NewGroupBox(parent Controller) *GroupBox {
    gb := new(GroupBox)
    gb.init(parent)

    gb.SetFont(DefaultFont)
    gb.SetCaption("GroupBox")
    gb.SetSize(100, 100)

    return gb
}

func AttachGroupBox(parent Controller, id int) *GroupBox {
    gb := new(GroupBox)
    gb.attach(parent, id)
    RegMsgHandler(gb)

    return gb
}

func (this *GroupBox) init(parent Controller) {
    this.W32Control.init("BUTTON", parent, 0, w32.WS_CHILD|w32.WS_VISIBLE|w32.WS_GROUP|w32.BS_GROUPBOX)
    RegMsgHandler(this)
}
