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

func (this *Canvas) DrawRect(rect *Rect, pen *Pen, brush *Brush) {
    w32Rect := rect.GetW32Rect()

    previousPen := gdi32.SelectObject(this.hdc, w32.HGDIOBJ(pen.GetHPEN()))
    if previousPen == 0 {
        panic("SelectObject for pen failed")
    }
    defer gdi32.SelectObject(this.hdc, previousPen)

    previousBrush := gdi32.SelectObject(this.hdc, w32.HGDIOBJ(brush.GetHBRUSH()))
    if previousBrush == 0 {
        panic("SelectObject for brush failed")
    }
    defer gdi32.SelectObject(this.hdc, previousBrush)

    gdi32.Rectangle(this.hdc, w32Rect.Left, w32Rect.Top, w32Rect.Right, w32Rect.Bottom)
}

func (this *Canvas) FillRect(rect *Rect, brush *Brush) {
    user32.FillRect(this.hdc, rect.GetW32Rect(), brush.GetHBRUSH())
}