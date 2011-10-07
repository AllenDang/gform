package gform

import (
    "syscall"
    "unsafe"
    "w32"
    "w32/user32"
)

type ListView struct {
    W32Control

    onEndLabelEdit LVEndLabelEditEventManagerA
    onDBLClick     LVDBLClickEventManagerA
    onClick        LVDBLClickEventManagerA
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
    RegMsgHandler(lv)

    user32.SendMessage(lv.Handle(), w32.LVM_SETUNICODEFORMAT, w32.TRUE, 0)
    return lv
}

func (this *ListView) init(parent Controller) {
    this.W32Control.init("SysListView32", parent, 0, w32.WS_CHILD|w32.WS_VISIBLE|w32.WS_BORDER|w32.LVS_REPORT|w32.LVS_EDITLABELS)
    RegMsgHandler(this)
}

func (this *ListView) SetSingleSelect(enable bool) {
    ToggleStyle(this.hwnd, enable, w32.LVS_SINGLESEL)
}

func (this *ListView) SetSortHeader(enable bool) {
    ToggleStyle(this.hwnd, enable, w32.LVS_NOSORTHEADER)
}

func (this *ListView) SetSortAscending(enable bool) {
    ToggleStyle(this.hwnd, enable, w32.LVS_SORTASCENDING)
}

func (this *ListView) SetEditLabels(enable bool) {
    ToggleStyle(this.hwnd, enable, w32.LVS_EDITLABELS)
}

func (this *ListView) SetFullRowSelect(enable bool) {
    if enable {
        user32.SendMessage(this.hwnd, w32.LVM_SETEXTENDEDLISTVIEWSTYLE, 0, w32.LVS_EX_FULLROWSELECT)
    } else {
        user32.SendMessage(this.hwnd, w32.LVM_SETEXTENDEDLISTVIEWSTYLE, w32.LVS_EX_FULLROWSELECT, 0)
    }
}

func (this *ListView) SetEnableDoubleBuffer(enable bool) {
    if enable {
        user32.SendMessage(this.hwnd, w32.LVM_SETEXTENDEDLISTVIEWSTYLE, 0, w32.LVS_EX_DOUBLEBUFFER)
    } else {
        user32.SendMessage(this.hwnd, w32.LVM_SETEXTENDEDLISTVIEWSTYLE, w32.LVS_EX_DOUBLEBUFFER, 0)
    }
}

func (this *ListView) SetItemCount(count int) bool {
    return user32.SendMessage(this.hwnd, w32.LVM_SETITEMCOUNT, uintptr(count), 0) != 0
}

func (this *ListView) GetItemCount() int {
    return int(user32.SendMessage(this.hwnd, w32.LVM_GETITEMCOUNT, 0, 0))
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
    this.InsertItem(caption, this.GetItemCount())
}

func (this *ListView) InsertLvColumn(lvColumn *w32.LVCOLUMN, iCol int) {
    user32.SendMessage(this.hwnd, w32.LVM_INSERTCOLUMN, uintptr(iCol), uintptr(unsafe.Pointer(lvColumn)))
}

func (this *ListView) InsertLvItem(lvItem *w32.LVITEM) {
    user32.SendMessage(this.hwnd, w32.LVM_INSERTITEM, 0, uintptr(unsafe.Pointer(lvItem)))
}

func (this *ListView) DeleteAllItems() bool {
    return user32.SendMessage(this.hwnd, w32.LVM_DELETEALLITEMS, 0, 0) == w32.TRUE
}

func (this *ListView) GetItem(item *w32.LVITEM) bool {
    return user32.SendMessage(this.hwnd, w32.LVM_GETITEM, 0, uintptr(unsafe.Pointer(item))) == w32.TRUE
}

func (this *ListView) GetSelectedCount() uint {
    return uint(user32.SendMessage(this.hwnd, w32.LVM_GETSELECTEDCOUNT, 0, 0))
}

// Event publishers
func (this *ListView) OnEndLabelEdit() *LVEndLabelEditEventManagerA {
    return &this.onEndLabelEdit
}

func (this *ListView) OnDBLClick() *LVDBLClickEventManagerA {
    return &this.onDBLClick
}

func (this *ListView) OnClick() *LVDBLClickEventManagerA {
    return &this.onClick
}

// Message processer
func (this *ListView) WndProc(msg uint, wparam, lparam uintptr) uintptr {
    switch msg {
    case w32.WM_NOTIFY:
        nm := (*w32.NMHDR)(unsafe.Pointer(lparam))
        switch int(nm.Code) {
        case w32.LVN_BEGINLABELEDITW:
            // println("Begin label edit")
        case w32.LVN_ENDLABELEDITW:
            nmdi := (*w32.NMLVDISPINFO)(unsafe.Pointer(lparam))
            if nmdi.Item.PszText != nil {
                this.onEndLabelEdit.Fire(this, &LVEndLabelEditEventArg{Item: &nmdi.Item})
                return w32.TRUE
            }
        case w32.NM_DBLCLK:
            nmItem := (*w32.NMITEMACTIVATE)(unsafe.Pointer(lparam))
            this.onDBLClick.Fire(this, &LVDBLClickEventArg{NmItem: nmItem})
        case w32.NM_CLICK:
            nmItem := (*w32.NMITEMACTIVATE)(unsafe.Pointer(lparam))
            this.onClick.Fire(this, &LVDBLClickEventArg{NmItem: nmItem})
        }
    }

    return this.W32Control.WndProc(msg, wparam, lparam)
}
