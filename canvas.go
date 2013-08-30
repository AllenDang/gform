package gform

import (
	"fmt"
	"github.com/AllenDang/w32"
)

type Canvas struct {
	hwnd         w32.HWND
	hdc          w32.HDC
	doNotDispose bool
}

func NewCanvasFromHwnd(hwnd w32.HWND) *Canvas {
	hdc := w32.GetDC(hwnd)
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
			w32.DeleteDC(this.hdc)
		} else {
			w32.ReleaseDC(this.hwnd, this.hdc)
		}

		this.hdc = 0
	}
}

func (this *Canvas) DrawBitmap(bmp *Bitmap, x, y int) {
	cdc := w32.CreateCompatibleDC(0)
	defer w32.DeleteDC(cdc)

	hbmpOld := w32.SelectObject(cdc, w32.HGDIOBJ(bmp.GetHBITMAP()))
	defer w32.SelectObject(cdc, w32.HGDIOBJ(hbmpOld))

	w, h := bmp.Size()

	w32.BitBlt(this.hdc, x, y, w, h, cdc, 0, 0, w32.SRCCOPY)
}

func (this *Canvas) DrawStretchedBitmap(bmp *Bitmap, rect *Rect) {
	cdc := w32.CreateCompatibleDC(0)
	defer w32.DeleteDC(cdc)

	hbmpOld := w32.SelectObject(cdc, w32.HGDIOBJ(bmp.GetHBITMAP()))
	defer w32.SelectObject(cdc, w32.HGDIOBJ(hbmpOld))

	w, h := bmp.Size()

	rc := rect.GetW32Rect()
	w32.StretchBlt(this.hdc, int(rc.Left), int(rc.Top), int(rc.Right), int(rc.Bottom), cdc, 0, 0, w, h, w32.SRCCOPY)
}

func (this *Canvas) DrawIcon(ico *Icon, x, y int) bool {
	return w32.DrawIcon(this.hdc, x, y, ico.Handle())
}

func (this *Canvas) DrawRect(rect *Rect, pen *Pen, brush *Brush) {
	w32Rect := rect.GetW32Rect()

	previousPen := w32.SelectObject(this.hdc, w32.HGDIOBJ(pen.GetHPEN()))
	defer w32.SelectObject(this.hdc, previousPen)

	previousBrush := w32.SelectObject(this.hdc, w32.HGDIOBJ(brush.GetHBRUSH()))
	defer w32.SelectObject(this.hdc, previousBrush)

	w32.Rectangle(this.hdc, int(w32Rect.Left), int(w32Rect.Top), int(w32Rect.Right), int(w32Rect.Bottom))
}

func (this *Canvas) FillRect(rect *Rect, brush *Brush) {
	w32.FillRect(this.hdc, rect.GetW32Rect(), brush.GetHBRUSH())
}

// Refer win32 DrawText document for uFormat.
func (this *Canvas) DrawText(text string, rect *Rect, format uint, font *Font, textColor Color) {
	previousFont := w32.SelectObject(this.hdc, w32.HGDIOBJ(font.GetHFONT()))
	defer w32.SelectObject(this.hdc, w32.HGDIOBJ(previousFont))

	previousBkMode := w32.SetBkMode(this.hdc, w32.TRANSPARENT)
	defer w32.SetBkMode(this.hdc, previousBkMode)

	previousTextColor := w32.SetTextColor(this.hdc, w32.COLORREF(textColor))
	defer w32.SetTextColor(this.hdc, previousTextColor)

	w32.DrawText(this.hdc, text, len(text), rect.GetW32Rect(), format)
}
