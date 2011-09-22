package gform

import (
    "w32"
)

type MsgHandler interface {
    WndProc(hwnd w32.HWND, msg uint, wparam ,lparam uintptr) uintptr
}

func RegMsgHandler(hwnd w32.HWND, msgHandler MsgHandler) {
    if _, isExists := msgHandlerRegistry[hwnd]; !isExists {
        msgHandlerRegistry[hwnd] = msgHandler
    }
}

func UnRegMsgHandler(hwnd w32.HWND) {
    msgHandlerRegistry[hwnd] = nil, false
}

func GetMsgHandler(hwnd w32.HWND) MsgHandler {
    if msgHandler, isExists := msgHandlerRegistry[hwnd]; isExists {
        return msgHandler
    }

    return nil
}
