package gform

import (
    "unsafe"
    "w32"
    "w32/kernel32"
    "w32/user32"
    "w32/comctl32"
)

func GetAppInstance() w32.HINSTANCE {
    return gAppInstance
}

func Init() {
    gAppInstance = kernel32.GetModuleHandle("")
    if gAppInstance == 0 {
        panic("Error occurred in App.Init")
    }

    // Initialize the common controls
    var initCtrls w32.INITCOMMONCONTROLSEX
    initCtrls.DwSize = uint32(unsafe.Sizeof(initCtrls))
    initCtrls.DwICC = w32.ICC_LISTVIEW_CLASSES | w32.ICC_PROGRESS_CLASS | w32.ICC_TAB_CLASSES | w32.ICC_TREEVIEW_CLASSES

    comctl32.InitCommonControlsEx(&initCtrls)
}

func RunMainLoop() int {
    var m w32.MSG
    var ret int
    for {
        ret = user32.GetMessage(&m, 0, 0, 0)
        if ret == 0 {
            break
        }
        if ret == -1 {
            panic("Error in App.Run.GetMessage")
        }

        user32.TranslateMessage(&m)
        user32.DispatchMessage(&m)
    }

    return int(m.WParam)
}

func Exit() {
    user32.PostQuitMessage(0)
}
