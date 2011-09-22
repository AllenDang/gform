package gform

import (
    "fmt"
    "w32"
    "w32/user32"
    "w32/gdi32"
)

type Canvas struct {
    hwnd         w32.HWND
    hdc          w32.HDC
    doNotDispose bool
}

func NewCanvasFromHwnd(hwnd w32.HWND) *Canvas {
    hdc := user32.GetDC(hwnd)
    if hdc == 0 {
        panic(fmt.Sprintf("Create canvas from %v failed.", hwnd))
    }

    return &Canvas{hwnd: hwnd, hdc: hdc, doNotDispose: false}
}

func NewCanvasFromHDC(hdc w32.HDC) *Canvas {
    if hdc == 0 {
        panic("Cannot create canvas from invalid HDC.")
    }

    return &Canvas{hdc: hdc, doNotDispose: true}
}

func (this *Canvas) Dispose() {
    if !this.doNotDispose && this.hdc != 0 {
        if this.hwnd == 0 {
            gdi32.DeleteDC(this.hdc)
        } else {
            user32.ReleaseDC(this.hwnd, this.hdc)
        }

        this.hdc = 0
    }
}

func (this *Canvas) DrawIcon(ico *Icon, x, y int) bool {
    return user32.DrawIcon(this.hdc, x, y, ico.Handle())
}