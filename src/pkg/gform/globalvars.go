package gform

import (
    "syscall"
    "container/vector"
    "w32"
)

//Private global variables.
var (
    gAppInstance       w32.HINSTANCE
    msgHandlerRegistry map[w32.HWND]MsgHandler
    registeredClasses  vector.StringVector
)

//Public global variables.
var (
    GeneralWndprocCallBack = syscall.NewCallback(generalWndProc)
    DefaultFont            *Font
)
