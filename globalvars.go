package gform

import (
    "github.com/AllenDang/w32"
    "syscall"
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
