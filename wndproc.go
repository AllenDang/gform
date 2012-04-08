package gform

import (
    "github.com/AllenDang/w32"
    "github.com/AllenDang/w32/shell32"
    "github.com/AllenDang/w32/user32"
    "unsafe"
)

func genMouseEventArg(wparam, lparam uintptr) *MouseEventData {
    var data MouseEventData
    data.Button = int(wparam)
    data.X = int(w32.LOWORD(uint(lparam)))
    data.Y = int(w32.HIWORD(uint(lparam)))

    return &data
}

func genDropFilesEventArg(wparam uintptr) *DropFilesEventData {
    hDrop := w32.HDROP(wparam)

    var data DropFilesEventData
    _, fileCount := shell32.DragQueryFile(hDrop, 0xFFFFFFFF)
    data.Files = make([]string, fileCount)

    var i uint
    for i = 0; i < fileCount; i++ {
        data.Files[i], _ = shell32.DragQueryFile(hDrop, i)
    }

    data.X, data.Y, _ = shell32.DragQueryPoint(hDrop)

    shell32.DragFinish(hDrop)

    return &data
}

func generalWndProc(hwnd w32.HWND, msg uint, wparam, lparam uintptr) uintptr {
    if msg == w32.WM_INITDIALOG && gDialogWaiting != nil {
        gDialogWaiting.hwnd = hwnd
        RegMsgHandler(gDialogWaiting)
    }

    if controller := GetMsgHandler(hwnd); controller != nil {
        ret := controller.WndProc(msg, wparam, lparam)
        switch msg {
        case w32.WM_NOTIFY: //Reflect notification to control
            nm := (*w32.NMHDR)(unsafe.Pointer(lparam))
            if controller := GetMsgHandler(nm.HwndFrom); controller != nil {
                ret := controller.WndProc(msg, wparam, lparam)
                if ret != 0 {
                    user32.SetWindowLong(hwnd, w32.DWL_MSGRESULT, uint32(ret))
                    return w32.TRUE
                }
            }
        case w32.WM_COMMAND:
            if lparam != 0 { //Reflect message to control
                h := w32.HWND(lparam)
                if controller := GetMsgHandler(h); controller != nil {
                    ret := controller.WndProc(msg, wparam, lparam)
                    if ret != 0 {
                        user32.SetWindowLong(hwnd, w32.DWL_MSGRESULT, uint32(ret))
                        return w32.TRUE
                    }
                }
            }
        case w32.WM_CLOSE:
            controller.OnClose().Fire(NewEventArg(controller, nil))
        case w32.WM_KILLFOCUS:
            controller.OnKillFocus().Fire(NewEventArg(controller, nil))
        case w32.WM_SETFOCUS:
            controller.OnSetFocus().Fire(NewEventArg(controller, nil))
        case w32.WM_DROPFILES:
            controller.OnDropFiles().Fire(NewEventArg(controller, genDropFilesEventArg(wparam)))
        case w32.WM_LBUTTONDOWN:
            controller.OnLBDown().Fire(NewEventArg(controller, genMouseEventArg(wparam, lparam)))
        case w32.WM_LBUTTONUP:
            controller.OnLBUp().Fire(NewEventArg(controller, genMouseEventArg(wparam, lparam)))
        case w32.WM_MBUTTONDOWN:
            controller.OnMBDown().Fire(NewEventArg(controller, genMouseEventArg(wparam, lparam)))
        case w32.WM_MBUTTONUP:
            controller.OnMBUp().Fire(NewEventArg(controller, genMouseEventArg(wparam, lparam)))
        case w32.WM_RBUTTONDOWN:
            controller.OnRBDown().Fire(NewEventArg(controller, genMouseEventArg(wparam, lparam)))
        case w32.WM_RBUTTONUP:
            controller.OnRBUp().Fire(NewEventArg(controller, genMouseEventArg(wparam, lparam)))
        case w32.WM_PAINT:
            canvas := NewCanvasFromHwnd(hwnd)
            defer canvas.Dispose()
            controller.OnPaint().Fire(NewEventArg(controller, &PaintEventData{Canvas: canvas}))
        case w32.WM_KEYUP:
            controller.OnKeyUp().Fire(NewEventArg(controller, &KeyUpEventData{int(wparam), int(lparam)}))
        }

        //Trigger msg handler registered via "Bind".
        if handler, ok := controller.BindedHandler(msg); ok {
            handler(NewEventArg(controller, &RawMsg{hwnd, msg, wparam, lparam}))
        }

        return ret
    }

    return user32.DefWindowProc(hwnd, msg, wparam, lparam)
}
