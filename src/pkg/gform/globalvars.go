package gform

import (
    "syscall"
    "container/vector"
    "w32"
)

//Private global variables.
var (
    gAppInstance        w32.HINSTANCE
    gControllerRegistry map[w32.HWND]Controller
    gRegisteredClasses  vector.StringVector
)

//Public global variables.
var (
    GeneralWndprocCallBack = syscall.NewCallback(generalWndProc)
    DefaultFont            *Font
)
