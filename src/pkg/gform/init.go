package gform

import (
    "w32"
)

func init() {
    msgHandlerRegistry = make(map[w32.HWND]MsgHandler)
}
