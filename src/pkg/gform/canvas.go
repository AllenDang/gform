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
    defer gdi32.SelectObject(this.hdc, previousPen)

    previousBrush := gdi32.SelectObject(this.hdc, w32.HGDIOBJ(brush.GetHBRUSH()))
    defer gdi32.SelectObject(this.hdc, previousBrush)

    gdi32.Rectangle(this.hdc, w32Rect.Left, w32Rect.Top, w32Rect.Right, w32Rect.Bottom)
}

func (this *Canvas) FillRect(rect *Rect, brush *Brush) {
    user32.FillRect(this.hdc, rect.GetW32Rect(), brush.GetHBRUSH())
}

// Refer uFormat parameter for win32 DrawText.
func (this *Canvas) DrawText(text string, rect *Rect, format uint, font *Font, textColor Color) {
    previousFont := gdi32.SelectObject(this.hdc, w32.HGDIOBJ(font.GetHFONT()))
    defer gdi32.SelectObject(this.hdc, w32.HGDIOBJ(previousFont))

    previousBkMode := gdi32.SetBkMode(this.hdc, w32.TRANSPARENT)
    defer gdi32.SetBkMode(this.hdc, previousBkMode)

    previousTextColor := gdi32.SetTextColor(this.hdc, w32.COLORREF(textColor))
    defer gdi32.SetTextColor(this.hdc, previousTextColor)

    user32.DrawText(this.hdc, text, len(text), rect.GetW32Rect(), format)
}