package gform

import (
    "github.com/AllenDang/w32"
)

type ControlBase struct {
    hwnd          w32.HWND
    font          *Font
    parent        Controller
    evtHandlerMap map[uint]EventHandler

    isForm bool

    // General events
    onCreate EventManager
    onClose  EventManager

    // Focus events
    onKillFocus EventManager
    onSetFocus  EventManager

    // Drag and drop events
    onDropFiles EventManager

    // Mouse events
    onLBDown EventManager
    onLBUp   EventManager
    onMBDown EventManager
    onMBUp   EventManager
    onRBDown EventManager
    onRBUp   EventManager

    onMouseHover EventManager
    onMouseLeave EventManager

    // Keyboard events
    onKeyUp EventManager

    // Paint events
    onPaint EventManager
}

func (this *ControlBase) init(parent Controller) {
    if this.hwnd == 0 {
        panic("hwnd cannot be nil")
    }

    if parent != nil {
        this.parent = parent
    }
}

func (this *ControlBase) Handle() w32.HWND {
    return this.hwnd
}

func (this *ControlBase) SetCaption(caption string) {
    w32.SetWindowText(this.hwnd, caption)
}

func (this *ControlBase) Caption() string {
    return w32.GetWindowText(this.hwnd)
}

func (this *ControlBase) Close() {
    UnRegMsgHandler(this.hwnd)
    w32.DestroyWindow(this.hwnd)
}

func (this *ControlBase) SetSize(width, height int) {
    x, y := this.Pos()
    w32.MoveWindow(this.hwnd, x, y, width, height, true)
}

func (this *ControlBase) Size() (width, height int) {
    rect := w32.GetWindowRect(this.hwnd)
    width = int(rect.Right - rect.Left)
    height = int(rect.Bottom - rect.Top)
    return
}

func (this *ControlBase) Width() int {
    rect := w32.GetWindowRect(this.hwnd)
    return int(rect.Right - rect.Left)
}

func (this *ControlBase) Height() int {
    rect := w32.GetWindowRect(this.hwnd)
    return int(rect.Bottom - rect.Top)
}

func (this *ControlBase) SetPos(x, y int) {
    w, h := this.Size()
    if w == 0 {
        w = 100
    }
    if h == 0 {
        h = 25
    }
    w32.MoveWindow(this.hwnd, x, y, w, h, true)
}

func (this *ControlBase) Pos() (x, y int) {
    rect := w32.GetWindowRect(this.hwnd)
    x = int(rect.Left)
    y = int(rect.Top)
    if !this.isForm && this.parent != nil {
        x, y = w32.ScreenToClient(this.parent.Handle(), x, y)
    }
    return
}

func (this *ControlBase) Visible() bool {
    return w32.IsWindowVisible(this.hwnd)
}

func (this *ControlBase) Bounds() *Rect {
    rect := w32.GetWindowRect(this.hwnd)
    if this.isForm {
        return &Rect{*rect}
    }

    return ScreenToClientRect(this.hwnd, rect)
}

func (this *ControlBase) ClientRect() *Rect {
    rect := w32.GetClientRect(this.hwnd)
    return ScreenToClientRect(this.hwnd, rect)
}

func (this *ControlBase) Show() {
    w32.ShowWindow(this.hwnd, w32.SW_SHOWDEFAULT)
}

func (this *ControlBase) Hide() {
    w32.ShowWindow(this.hwnd, w32.SW_HIDE)
}

func (this *ControlBase) Enabled() bool {
    return w32.IsWindowEnabled(this.hwnd)
}

func (this *ControlBase) SetEnabled(b bool) {
    w32.EnableWindow(this.hwnd, b)
}

func (this *ControlBase) Focus() {
    w32.SetFocus(this.hwnd)
}

func (this *ControlBase) Invalidate(erase bool) {
    pRect := w32.GetClientRect(this.hwnd)
    if this.isForm {
        w32.InvalidateRect(this.hwnd, pRect, erase)
    } else {
        rc := ScreenToClientRect(this.hwnd, pRect)
        w32.InvalidateRect(this.hwnd, rc.GetW32Rect(), erase)
    }
}

func (this *ControlBase) Parent() Controller {
    return this.parent
}

func (this *ControlBase) Font() *Font {
    return this.font
}

func (this *ControlBase) SetFont(font *Font) {
    w32.SendMessage(this.hwnd, w32.WM_SETFONT, uintptr(font.hfont), 1)
    this.font = font
}

func (this *ControlBase) EnableDragAcceptFiles(b bool) {
    w32.DragAcceptFiles(this.hwnd, b)
}

func (this *ControlBase) InvokeRequired() bool {
    if this.hwnd == 0 {
        return false
    }

    windowThreadId, _ := w32.GetWindowThreadProcessId(this.hwnd)
    currentThreadId := w32.GetCurrentThread()

    return windowThreadId != currentThreadId
}

func (this *ControlBase) PreTranslateMessage(msg *w32.MSG) bool {
    return false
}

func (this *ControlBase) Bind(msg uint, handler EventHandler) {
    //Check whether map is already created
    if this.evtHandlerMap == nil {
        this.evtHandlerMap = make(map[uint]EventHandler)
    }

    if handler == nil {
        delete(this.evtHandlerMap, msg)
    } else {
        this.evtHandlerMap[msg] = handler
    }
}

// Get binded handlers for specifed message.
func (this *ControlBase) BindedHandler(msg uint) (EventHandler, bool) {
    handler, ok := this.evtHandlerMap[msg]
    return handler, ok
}

//Events
func (this *ControlBase) OnCreate() *EventManager {
    return &this.onCreate
}

func (this *ControlBase) OnClose() *EventManager {
    return &this.onClose
}

func (this *ControlBase) OnKillFocus() *EventManager {
    return &this.onKillFocus
}

func (this *ControlBase) OnSetFocus() *EventManager {
    return &this.onSetFocus
}

func (this *ControlBase) OnDropFiles() *EventManager {
    return &this.onDropFiles
}

func (this *ControlBase) OnLBDown() *EventManager {
    return &this.onLBDown
}

func (this *ControlBase) OnLBUp() *EventManager {
    return &this.onLBUp
}

func (this *ControlBase) OnMBDown() *EventManager {
    return &this.onMBDown
}

func (this *ControlBase) OnMBUp() *EventManager {
    return &this.onMBUp
}

func (this *ControlBase) OnRBDown() *EventManager {
    return &this.onRBDown
}

func (this *ControlBase) OnRBUp() *EventManager {
    return &this.onRBUp
}

func (this *ControlBase) OnMouseHover() *EventManager {
    return &this.onMouseHover
}

func (this *ControlBase) OnMouseLeave() *EventManager {
    return &this.onMouseLeave
}

func (this *ControlBase) OnPaint() *EventManager {
    return &this.onPaint
}

func (this *ControlBase) OnKeyUp() *EventManager {
    return &this.onKeyUp
}
