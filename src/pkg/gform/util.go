package gform

import (
    "fmt"
    "syscall"
    "unsafe"
    "unicode/utf16"
    "w32"
    "w32/user32"
)

func UTF16PtrToString(cstr *uint16) string {
    if cstr != nil {
        us := make([]uint16, 0, 256)
        for p := uintptr(unsafe.Pointer(cstr)); ; p += 2 {
            u := *(*uint16)(unsafe.Pointer(p))
            if u == 0 {
                return string(utf16.Decode(us))
            }
            us = append(us, u)
        }
    }

    return ""
}

func ToggleStyle(hwnd w32.HWND, b bool, style int) {
    originalStyle := int(user32.GetWindowLongPtr(hwnd, w32.GWL_STYLE))
    if originalStyle != 0 {
        if b {
            originalStyle |= style
        } else {
            originalStyle ^= style
        }
        user32.SetWindowLongPtr(hwnd, w32.GWL_STYLE, uintptr(originalStyle))
    }
}

func ToggleExStyle(hwnd w32.HWND, b bool, style int) {
    originalStyle := int(user32.GetWindowLongPtr(hwnd, w32.GWL_EXSTYLE))
    if originalStyle != 0 {
        if b {
            originalStyle |= style
        } else {
            originalStyle ^= style
        }
        user32.SetWindowLongPtr(hwnd, w32.GWL_EXSTYLE, uintptr(originalStyle))
    }
}

func CreateWindow(className string, parent Controller, exStyle, style uint) w32.HWND {
    instance := GetAppInstance()
    var parentHwnd w32.HWND
    if parent != nil {
        parentHwnd = parent.Handle()
    }
    var hwnd w32.HWND
    hwnd = user32.CreateWindowEx(
        exStyle,
        syscall.StringToUTF16Ptr(className),
        nil,
        style,
        w32.CW_USEDEFAULT,
        w32.CW_USEDEFAULT,
        w32.CW_USEDEFAULT,
        w32.CW_USEDEFAULT,
        parentHwnd,
        0,
        instance,
        nil)

    if hwnd == 0 {
        errStr := fmt.Sprintf("Error occurred in CreateWindow(%s, %v, %d, %d)", className, parent, exStyle, style)
        panic(errStr)
    }

    return hwnd
}

func RegisterClass(className string, wndproc uintptr) {
    instance := GetAppInstance()
    icon := user32.LoadIcon(instance, w32.MakeIntResource(w32.IDI_APPLICATION))

    var wc w32.WNDCLASSEX
    wc.Size = uint(unsafe.Sizeof(wc))
    wc.Style = w32.CS_HREDRAW | w32.CS_VREDRAW
    wc.WndProc = wndproc
    wc.Instance = instance
    wc.Background = w32.COLOR_BTNFACE + 1
    wc.Icon = icon
    wc.Cursor = user32.LoadCursor(0, w32.MakeIntResource(w32.IDC_ARROW))
    wc.ClassName = syscall.StringToUTF16Ptr(className)
    wc.MenuName = nil
    wc.IconSm = icon

    if ret := user32.RegisterClassEx(&wc); ret == 0 {
        panic(syscall.GetLastError())
    }
}

func RegClassOnlyOnce(className string) {
    isExists := false
    for _, class := range gRegisteredClasses {
        if class == className {
            isExists = true
            break
        }
    }

    if !isExists {
        RegisterClass(className, GeneralWndprocCallBack)
        gRegisteredClasses = append(gRegisteredClasses, className)
    }
}
