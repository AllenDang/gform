package gform

import (
    "syscall"
    "w32"
)

//Private global variables.
var (
    gAppInstance        w32.HINSTANCE
    gControllerRegistry map[w32.HWND]Controller
    gRegisteredClasses  []string
    gDialogWaiting      *Dialog
)

//Public global variables.
var (
    GeneralWndprocCallBack = syscall.NewCallback(generalWndProc)
    DefaultFont            *Font
)
