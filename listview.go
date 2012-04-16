package gform

import (
    "github.com/AllenDang/w32"
    "github.com/AllenDang/w32/user32"
    "syscall"
    "unsafe"
)

type ListView struct {
    W32Control

    onEndLabelEdit EventManager
    onDBLClick     EventManager
    onClick        EventManager
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

// Changes the state of an item in a list-view control. Refer LVM_SETITEMSTATE message.
func (this *ListView) setItemState(i int, state, mask uint) {
    var item w32.LVITEM
    item.State, item.StateMask = state, mask

    user32.SendMessage(this.hwnd, w32.LVM_SETITEMSTATE, uintptr(i), uintptr(unsafe.Pointer(&item)))
}

func (this *ListView) EnableSingleSelect(enable bool) {
    ToggleStyle(this.hwnd, enable, w32.LVS_SINGLESEL)
}

func (this *ListView) EnableSortHeader(enable bool) {
    ToggleStyle(this.hwnd, enable, w32.LVS_NOSORTHEADER)
}

func (this *ListView) EnableSortAscending(enable bool) {
    ToggleStyle(this.hwnd, enable, w32.LVS_SORTASCENDING)
}

func (this *ListView) EnableEditLabels(enable bool) {
    ToggleStyle(this.hwnd, enable, w32.LVS_EDITLABELS)
}

func (this *ListView) EnableFullRowSelect(enable bool) {
    if enable {
        user32.SendMessage(this.hwnd, w32.LVM_SETEXTENDEDLISTVIEWSTYLE, 0, w32.LVS_EX_FULLROWSELECT)
    } else {
        user32.SendMessage(this.hwnd, w32.LVM_SETEXTENDEDLISTVIEWSTYLE, w32.LVS_EX_FULLROWSELECT, 0)
    }
}

func (this *ListView) EnableDoubleBuffer(enable bool) {
    if enable {
        user32.SendMessage(this.hwnd, w32.LVM_SETEXTENDEDLISTVIEWSTYLE, 0, w32.LVS_EX_DOUBLEBUFFER)
    } else {
        user32.SendMessage(this.hwnd, w32.LVM_SETEXTENDEDLISTVIEWSTYLE, w32.LVS_EX_DOUBLEBUFFER, 0)
    }
}

func (this *ListView) EnableHotTrack(enable bool) {
    if enable {
        user32.SendMessage(this.hwnd, w32.LVM_SETEXTENDEDLISTVIEWSTYLE, 0, w32.LVS_EX_TRACKSELECT)
    } else {
        user32.SendMessage(this.hwnd, w32.LVM_SETEXTENDEDLISTVIEWSTYLE, w32.LVS_EX_TRACKSELECT, 0)
    }
}

func (this *ListView) SetItemCount(count int) bool {
    return user32.SendMessage(this.hwnd, w32.LVM_SETITEMCOUNT, uintptr(count), 0) != 0
}

func (this *ListView) ItemCount() int {
    return int(user32.SendMessage(this.hwnd, w32.LVM_GETITEMCOUNT, 0, 0))
}

func (this *ListView) InsertColumn(caption string, width int, iCol int) {
    var lc w32.LVCOLUMN
    lc.Mask = w32.LVCF_TEXT
    if width != 0 {
        lc.Mask = lc.Mask | w32.LVCF_WIDTH
        lc.Cx = width
    }
    lc.PszText = syscall.StringToUTF16Ptr(caption)

    this.InsertLvColumn(&lc, iCol)
}

func (this *ListView) AddItem(text ...string) {
    if len(text) > 0 {
        var li w32.LVITEM
        li.Mask = w32.LVIF_TEXT
        li.PszText = syscall.StringToUTF16Ptr(text[0])
        li.IItem = this.ItemCount()

        this.InsertLvItem(&li)

        for i := 1; i < len(text); i++ {
            li.PszText = syscall.StringToUTF16Ptr(text[i])
            li.ISubItem = i

            this.SetLvItem(&li)
        }
    }
}

func (this *ListView) InsertLvColumn(lvColumn *w32.LVCOLUMN, iCol int) {
    user32.SendMessage(this.hwnd, w32.LVM_INSERTCOLUMN, uintptr(iCol), uintptr(unsafe.Pointer(lvColumn)))
}

func (this *ListView) InsertLvItem(lvItem *w32.LVITEM) {
    user32.SendMessage(this.hwnd, w32.LVM_INSERTITEM, 0, uintptr(unsafe.Pointer(lvItem)))
}

func (this *ListView) SetLvItem(lvItem *w32.LVITEM) {
    user32.SendMessage(this.hwnd, w32.LVM_SETITEM, 0, uintptr(unsafe.Pointer(lvItem)))
}

func (this *ListView) DeleteAllItems() bool {
    return user32.SendMessage(this.hwnd, w32.LVM_DELETEALLITEMS, 0, 0) == w32.TRUE
}

func (this *ListView) Item(item *w32.LVITEM) bool {
    return user32.SendMessage(this.hwnd, w32.LVM_GETITEM, 0, uintptr(unsafe.Pointer(item))) == w32.TRUE
}

func (this *ListView) ItemAtIndex(i int) *w32.LVITEM {
    var item w32.LVITEM
    item.Mask = w32.LVIF_PARAM | w32.LVIF_TEXT
    item.IItem = i

    this.Item(&item)
    return &item
}

// mask is used to set the LVITEM.Mask for ListView.GetItem which indicates which attributes you'd like to receive
// of LVITEM.
func (this *ListView) SelectedItems(mask uint) []*w32.LVITEM {
    items := make([]*w32.LVITEM, 0)

    var i int = -1
    for {
        if i = int(user32.SendMessage(this.hwnd, w32.LVM_GETNEXTITEM, uintptr(i), uintptr(w32.LVNI_SELECTED))); i == -1 {
            break
        }

        var item w32.LVITEM
        item.Mask = mask
        item.IItem = i
        if this.Item(&item) {
            items = append(items, &item)
        }
    }
    return items
}

func (this *ListView) SelectedCount() uint {
    return uint(user32.SendMessage(this.hwnd, w32.LVM_GETSELECTEDCOUNT, 0, 0))
}

// Set i to -1 to select all items.
func (this *ListView) SetSelectedItem(i int) {
    this.setItemState(i, w32.LVIS_SELECTED, w32.LVIS_SELECTED)
}

func (this *ListView) SetImageList(imageList *ImageList, imageListType int) *ImageList {
    h := user32.SendMessage(this.hwnd, w32.LVM_SETIMAGELIST, uintptr(imageListType), uintptr(imageList.Handle()))
    if h == 0 {
        return nil
    }

    return &ImageList{w32.HIMAGELIST(h)}
}

func (this *ListView) ImageList(imageListType int) *ImageList {
    h := user32.SendMessage(this.hwnd, w32.LVM_GETIMAGELIST, uintptr(imageListType), 0)
    if h == 0 {
        return nil
    }

    return &ImageList{w32.HIMAGELIST(h)}
}

// Event publishers
func (this *ListView) OnEndLabelEdit() *EventManager {
    return &this.onEndLabelEdit
}

func (this *ListView) OnDBLClick() *EventManager {
    return &this.onDBLClick
}

func (this *ListView) OnClick() *EventManager {
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
                this.onEndLabelEdit.Fire(NewEventArg(this, &LVEndLabelEditEventData{Item: &nmdi.Item}))
                return w32.TRUE
            }
        case w32.NM_DBLCLK:
            nmItem := (*w32.NMITEMACTIVATE)(unsafe.Pointer(lparam))
            this.onDBLClick.Fire(NewEventArg(this, &LVDBLClickEventData{NmItem: nmItem}))
        case w32.NM_CLICK:
            nmItem := (*w32.NMITEMACTIVATE)(unsafe.Pointer(lparam))
            this.onClick.Fire(NewEventArg(this, &LVDBLClickEventData{NmItem: nmItem}))
        }
    }

    return this.W32Control.WndProc(msg, wparam, lparam)
}
