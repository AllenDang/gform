package gform

import (
    "w32"
)

type Controller interface {
    Caption() string
    Close()
    Enabled() bool
    Focus()
    Handle() w32.HWND
    Invalidate(erase bool)
    Parent() Controller
    Pos() (x, y int)
    Size() (w, h int)
    Visible() bool
    SetCaption(s string)
    SetEnabled(b bool)
    SetPos(x, y int)
    SetSize(w, h int)
    SetDragAcceptFilesEnabled(b bool)
    Show()
    Hide()
    Font() *Font
    SetFont(font *Font)
    InvokeRequired() bool
    PreTranslateMessage(msg *w32.MSG) bool
    WndProc(msg uint, wparam, lparam uintptr) uintptr

    // Focus events
    OnKillFocus() *GeneralEventManager
    OnSetFocus() *GeneralEventManager

    //Drag and drop events
    OnDropFilesA() *DropFilesEventManagerA

    //Mouse events
    OnLBDown() *GeneralEventManager
    OnLBUp() *GeneralEventManager
    OnMBDown() *GeneralEventManager
    OnMBUp() *GeneralEventManager
    OnRBDown() *GeneralEventManager
    OnRBUp() *GeneralEventManager

    OnLBDownA() *MouseEventManagerA
    OnLBUpA() *MouseEventManagerA
    OnMBDownA() *MouseEventManagerA
    OnMBUpA() *MouseEventManagerA
    OnRBDownA() *MouseEventManagerA
    OnRBUpA() *MouseEventManagerA

    //Paint events
    OnPaintA() *PaintEventManagerA
}
