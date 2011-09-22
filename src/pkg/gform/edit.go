package gform

import (
    "w32"
    "w32/user32"
)

type Edit struct {
    W32Control
}

func NewEdit(parent Controller) *Edit {
    edt := new(Edit)
    edt.init(parent)

    edt.SetFont(DefaultFont)
    edt.SetSize(200, 20)

    return edt
}

func AttachEdit(parent Controller, id int) *Edit {
    edt := new(Edit)
    edt.attach(parent, id)
    RegMsgHandler(edt.Handle(), edt)

    return edt
}

func (this *Edit) init(parent Controller) {
    this.W32Control.init("EDIT", parent, w32.WS_EX_CLIENTEDGE, w32.WS_CHILD|w32.WS_VISIBLE|w32.WS_TABSTOP|w32.ES_LEFT)
    RegMsgHandler(this.hwnd, this)
}

func (this *Edit) SetReadOnly(isReadOnly bool) {
    user32.SendMessage(this.hwnd, w32.EM_SETREADONLY, uintptr(w32.BoolToBOOL(isReadOnly)), 0)
}