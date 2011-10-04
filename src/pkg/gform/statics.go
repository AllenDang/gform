package gform

import (
    "w32"
)

type Label struct {
    W32Control
}

func NewLabel(parent Controller) *Label {
    lb := new(Label)
    lb.init(parent)

    lb.SetFont(DefaultFont)
    lb.SetCaption("Label")
    lb.SetSize(100, 25)

    return lb
}

func AttachLabel(parent Controller, id int) *Label {
    lb := new(Label)
    lb.attach(parent, id)
    RegMsgHandler(lb)

    return lb
}

func (this *Label) init(parent Controller) {
    this.W32Control.init("STATIC", parent, 0, w32.WS_CHILD|w32.WS_VISIBLE|w32.SS_LEFTNOWORDWRAP)
    RegMsgHandler(this)
}
