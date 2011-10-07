package gform

import (
    "w32"
    "w32/user32"
    "w32/kernel32"
    "w32/shell32"
)

type ControlBase struct {
    hwnd   w32.HWND
    parent Controller

    isForm bool

    //Focus events
    onKillFocus GeneralEventManager
    onSetFocus  GeneralEventManager

    //Drag and drop events
    onDropFilesA DropFilesEventManagerA

    //Mouse events
    onLBDown GeneralEventManager
    onLBUp   GeneralEventManager
    onMBDown GeneralEventManager
    onMBUp   GeneralEventManager
    onRBDown GeneralEventManager
    onRBUp   GeneralEventManager

    onLBDownA MouseEventManagerA
    onLBUpA   MouseEventManagerA
    onMBDownA MouseEventManagerA
    onMBUpA   MouseEventManagerA
    onRBDownA MouseEventManagerA
    onRBUpA   MouseEventManagerA

    //Paint events
    onPaintA PaintEventManagerA
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
    user32.SetWindowText(this.hwnd, caption)
}

func (this *ControlBase) Caption() string {
    return user32.GetWindowText(this.hwnd)
}

func (this *ControlBase) Close() {
    UnRegMsgHandler(this.hwnd)
    user32.DestroyWindow(this.hwnd)
}

func (this *ControlBase) SetSize(width, height int) {
    x, y := this.Pos()
    user32.MoveWindow(this.hwnd, x, y, width, height, true)
}

func (this *ControlBase) Size() (width, height int) {
    rect := user32.GetWindowRect(this.hwnd)
    width = int(rect.Right - rect.Left)
    height = int(rect.Bottom - rect.Top)
    return
}

func (this *ControlBase) SetPos(x, y int) {
    w, h := this.Size()
    if w == 0 {
        w = 100
    }
    if h == 0 {
        h = 25
    }
    user32.MoveWindow(this.hwnd, x, y, w, h, true)
}

func (this *ControlBase) Pos() (x, y int) {
    rect := user32.GetWindowRect(this.hwnd)
    x = int(rect.Left)
    y = int(rect.Top)
    if !this.isForm && this.parent != nil {
        x, y = user32.ScreenToClient(this.parent.Handle(), x, y)
    }
    return
}

func (this *ControlBase) Visible() bool {
    return user32.IsWindowVisible(this.hwnd)
}

func (this *ControlBase) Show() {
    user32.ShowWindow(this.hwnd, w32.SW_SHOWDEFAULT)
}

func (this *ControlBase) Hide() {
    user32.ShowWindow(this.hwnd, w32.SW_HIDE)
}

func (this *ControlBase) Enabled() bool {
    return user32.IsWindowEnabled(this.hwnd)
}

func (this *ControlBase) SetEnabled(b bool) {
    user32.EnableWindow(this.hwnd, b)
}

func (this *ControlBase) Focus() {
    user32.SetFocus(this.hwnd)
}

func (this *ControlBase) Invalidate(erase bool) {
    pRect := user32.GetClientRect(this.hwnd)
    user32.InvalidateRect(this.hwnd, pRect, erase)
}

func (this *ControlBase) Parent() Controller {
    return this.parent
}

func (this *ControlBase) Font() *Font {
    //TODO: Implement Font()
    return nil
}

func (this *ControlBase) SetFont(font *Font) {
    user32.SendMessage(this.hwnd, w32.WM_SETFONT, uintptr(font.hfont), 1)
}

func (this *ControlBase) SetDragAcceptFilesEnabled(b bool) {
    shell32.DragAcceptFiles(this.hwnd, b)
}

func (this *ControlBase) InvokeRequired() bool {
    if this.hwnd == 0 {
        return false
    }

    windowThreadId, _ := user32.GetWindowThreadProcessId(this.hwnd)
    currentThreadId := kernel32.GetCurrentThread()

    return windowThreadId != currentThreadId
}

func (this *ControlBase) PreTranslateMessage(msg *w32.MSG) bool {
    return false
}

//Events
func (this *ControlBase) OnKillFocus() *GeneralEventManager {
    return &this.onKillFocus
}

func (this *ControlBase) OnSetFocus() *GeneralEventManager {
    return &this.onSetFocus
}

func (this *ControlBase) OnDropFilesA() *DropFilesEventManagerA {
    return &this.onDropFilesA
}

func (this *ControlBase) OnLBDown() *GeneralEventManager {
    return &this.onLBDown
}

func (this *ControlBase) OnLBUp() *GeneralEventManager {
    return &this.onLBUp
}

func (this *ControlBase) OnMBDown() *GeneralEventManager {
    return &this.onMBDown
}

func (this *ControlBase) OnMBUp() *GeneralEventManager {
    return &this.onMBUp
}

func (this *ControlBase) OnRBDown() *GeneralEventManager {
    return &this.onRBDown
}

func (this *ControlBase) OnRBUp() *GeneralEventManager {
    return &this.onRBUp
}

func (this *ControlBase) OnLBDownA() *MouseEventManagerA {
    return &this.onLBDownA
}

func (this *ControlBase) OnLBUpA() *MouseEventManagerA {
    return &this.onLBUpA
}

func (this *ControlBase) OnMBDownA() *MouseEventManagerA {
    return &this.onMBDownA
}

func (this *ControlBase) OnMBUpA() *MouseEventManagerA {
    return &this.onMBUpA
}

func (this *ControlBase) OnRBDownA() *MouseEventManagerA {
    return &this.onRBDownA
}

func (this *ControlBase) OnRBUpA() *MouseEventManagerA {
    return &this.onRBUpA
}

func (this *ControlBase) OnPaintA() *PaintEventManagerA {
    return &this.onPaintA
}
