package gform

import (
    "w32"
    "w32/user32"
    "w32/shell32"
)

func genMouseEventArg(wparam, lparam uintptr) *MouseEventArg {
    var arg MouseEventArg
    arg.Button = int(wparam)
    arg.X = int(w32.LOWORD(uint(lparam)))
    arg.Y = int(w32.HIWORD(uint(lparam)))

    return &arg
}

func genDropFilesEventArg(wparam uintptr) *DropFilesEventArg {
    hDrop := w32.HDROP(wparam)

    var arg DropFilesEventArg
    _, fileCount := shell32.DragQueryFile(hDrop, 0xFFFFFFFF)
    arg.Files = make([]string, fileCount)

    var i uint
    for i = 0; i < fileCount; i++ {
        arg.Files[i], _ = shell32.DragQueryFile(hDrop, i)
    }

    arg.X, arg.Y, _ = shell32.DragQueryPoint(hDrop)

    shell32.DragFinish(hDrop)

    return &arg
}

func generalWndProc(hwnd w32.HWND, msg uint, wparam, lparam uintptr) uintptr {
    if msg == w32.WM_INITDIALOG && gDialogWaiting != nil {
        gDialogWaiting.hwnd = hwnd
        RegMsgHandler(gDialogWaiting)
    }

    if msgHandler := GetMsgHandler(hwnd); msgHandler != nil {
        ret := msgHandler.WndProc(msg, wparam, lparam)
        if controller, ok := msgHandler.(Controller); ok {
            switch msg {
            case w32.WM_KILLFOCUS:
                controller.OnKillFocus().Fire(controller)
            case w32.WM_SETFOCUS:
                controller.OnSetFocus().Fire(controller)
            case w32.WM_DROPFILES:
                controller.OnDropFilesA().Fire(controller, genDropFilesEventArg(wparam))
            case w32.WM_LBUTTONDOWN:
                controller.OnLBDownA().Fire(controller, genMouseEventArg(wparam, lparam))
                controller.OnLBDown().Fire(controller)
            case w32.WM_LBUTTONUP:
                controller.OnLBUpA().Fire(controller, genMouseEventArg(wparam, lparam))
                controller.OnLBUp().Fire(controller)
            case w32.WM_MBUTTONDOWN:
                controller.OnMBDownA().Fire(controller, genMouseEventArg(wparam, lparam))
                controller.OnMBDown().Fire(controller)
            case w32.WM_MBUTTONUP:
                controller.OnMBUpA().Fire(controller, genMouseEventArg(wparam, lparam))
                controller.OnMBUp().Fire(controller)
            case w32.WM_RBUTTONDOWN:
                controller.OnRBDownA().Fire(controller, genMouseEventArg(wparam, lparam))
                controller.OnRBDown().Fire(controller)
            case w32.WM_RBUTTONUP:
                controller.OnRBUpA().Fire(controller, genMouseEventArg(wparam, lparam))
                controller.OnRBUp().Fire(controller)
            case w32.WM_PAINT:
                canvas := NewCanvasFromHwnd(hwnd)
                controller.OnPaintA().Fire(controller, &PaintEventArg{Canvas: canvas})
                canvas.Dispose()
            case w32.WM_KEYUP:
                controller.OnKeyUpA().Fire(controller, &KeyUpEventArg{int(wparam), int(lparam)})
            }
        }
        return ret
    }

    return user32.DefWindowProc(hwnd, msg, wparam, lparam)
}
