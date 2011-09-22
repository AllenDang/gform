package gform

import (
    "syscall"
    "unsafe"
    "w32"
    "w32/user32"
)

type ListView struct {
    W32Control

    onEndLabelEditEventManager LVEndLabelEditEventManagerA
}

func NewListView(parent Controller) *ListView {
    lv := new(ListView)
    lv.init(parent)

    lv.SetFont(DefaultFont)
    lv.SetSize(100, 100)

    return lv
}

func AttachListView(parent Controller, id int) *ListView {
    lv := new(ListView)
    lv.attach(parent, id)
    RegMsgHandler(lv.Handle(), lv)

    user32.SendMessage(lv.Handle(), w32.LVM_SETUNICODEFORMAT, w32.TRUE, 0)
    return lv
}

func (this *ListView) init(parent Controller) {
    this.W32Control.init("SysListView32", parent, 0, w32.WS_CHILD|w32.WS_VISIBLE|w32.WS_BORDER|w32.LVS_REPORT|w32.LVS_EDITLABELS)
    RegMsgHandler(this.Handle(), this)
}

func (this *ListView) SetSingleSelect(enable bool) {
    ToggleStyle(this.Handle(), enable, w32.LVS_SINGLESEL)
}

func (this *ListView) SetSortHeader(enable bool) {
    ToggleStyle(this.Handle(), enable, w32.LVS_NOSORTHEADER)
}

func (this *ListView) SetEditLabels(enable bool) {
    ToggleStyle(this.Handle(), enable, w32.LVS_EDITLABELS)
}

func (this *ListView) SetFullRowSelect(enable bool) {
    if enable {
        user32.SendMessage(this.Handle(), w32.LVM_SETEXTENDEDLISTVIEWSTYLE, 0, w32.LVS_EX_FULLROWSELECT)
    } else {
        user32.SendMessage(this.Handle(), w32.LVM_SETEXTENDEDLISTVIEWSTYLE, w32.LVS_EX_FULLROWSELECT, 0)
    }
}

func (this *ListView) CountItem() int {
    return int(user32.SendMessage(this.Handle(), w32.LVM_GETITEMCOUNT, 0, 0))
}

func (this *ListView) InsertColumn(caption string, width int, iCol int) {
    var lc w32.LVCOLUMN
    lc.Mask = w32.LVCF_TEXT
    if width != 0 {
        lc.Mask = lc.Mask | w32.LVCF_WIDTH
        lc.Cx = int32(width)
    }
    lc.PszText = syscall.StringToUTF16Ptr(caption)

    this.InsertLvColumn(&lc, iCol)
}

func (this *ListView) InsertItem(caption string, index int) {
    var li w32.LVITEM
    li.Mask = w32.LVIF_TEXT
    li.PszText = syscall.StringToUTF16Ptr(caption)
    li.IItem = int32(index)

    this.InsertLvItem(&li)
}

func (this *ListView) AddItem(caption string) {
    this.InsertItem(caption, this.CountItem())
}

func (this *ListView) InsertLvColumn(lvColumn *w32.LVCOLUMN, iCol int) {
    user32.SendMessage(this.Handle(), w32.LVM_INSERTCOLUMN, uintptr(iCol), uintptr(unsafe.Pointer(lvColumn)))
}

func (this *ListView) InsertLvItem(lvItem *w32.LVITEM) {
    if this.InvokeRequired() {
        user32.PostMessage(this.Handle(), w32.LVM_INSERTITEM, 0, uintptr(unsafe.Pointer(lvItem)))
    } else {
        user32.SendMessage(this.Handle(), w32.LVM_INSERTITEM, 0, uintptr(unsafe.Pointer(lvItem)))
    }
}

// Event publishers
func (this *ListView) OnEndLabelEdit() *LVEndLabelEditEventManagerA {
    return &this.onEndLabelEditEventManager
}

func (this *ListView) WndProc(hwnd w32.HWND, msg uint, wparam, lparam uintptr) uintptr {
    switch msg {
    case w32.WM_NOTIFY:
        nm := (*w32.NMHDR)(unsafe.Pointer(lparam))
        switch int(nm.Code) {
        case w32.LVN_BEGINLABELEDITW:
            // println("Begin label edit")
        case w32.LVN_ENDLABELEDITW:
            nmdi := (*w32.NMLVDISPINFO)(unsafe.Pointer(lparam))
            if nmdi.Item.PszText != nil {
                str := UTF16PtrToString(nmdi.Item.PszText)
                this.onEndLabelEditEventManager.Fire(this, int(nmdi.Item.IItem), str)
                return w32.TRUE
            }
        }
    }

    return this.W32Control.WndProc(hwnd, msg, wparam, lparam)
}
