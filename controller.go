package gform

import (
    "github.com/AllenDang/w32"
)

type Controller interface {
    Caption() string
    Enabled() bool
    Focus()
    Handle() w32.HWND
    Invalidate(erase bool)
    Parent() Controller
    Pos() (x, y int)
    Size() (w, h int)
    Height() int
    Width() int
    Visible() bool
    Bounds() *Rect
    ClientRect() *Rect
    SetCaption(s string)
    SetEnabled(b bool)
    SetPos(x, y int)
    SetSize(w, h int)
    EnableDragAcceptFiles(b bool)
    Show()
    Hide()
    Font() *Font
    SetFont(font *Font)
    InvokeRequired() bool
    PreTranslateMessage(msg *w32.MSG) bool
    WndProc(msg uint, wparam, lparam uintptr) uintptr

    //Bind w32 message to handler function
    Bind(msg uint, handler EventHandler)
    BindedHandler(msg uint) (EventHandler, bool)

    //General events
    OnCreate() *EventManager
    OnClose() *EventManager

    // Focus events
    OnKillFocus() *EventManager
    OnSetFocus() *EventManager

    //Drag and drop events
    OnDropFiles() *EventManager

    //Mouse events
    OnLBDown() *EventManager
    OnLBUp() *EventManager
    OnMBDown() *EventManager
    OnMBUp() *EventManager
    OnRBDown() *EventManager
    OnRBUp() *EventManager

    OnMouseHover() *EventManager
    OnMouseLeave() *EventManager

    //Keyboard events
    OnKeyUp() *EventManager

    //Paint events
    OnPaint() *EventManager
    OnSize() *EventManager
}
