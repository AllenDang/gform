package gform

import (
	"w32"
	"w32/gdi32"
)

type Pen struct {
    hPen w32.HPEN
    style uint
    brush *Brush
}

func NewPen(style uint, width uint, brush *Brush) *Pen {
    if brush == nil {
        panic("Brush cannot be nil")
    }

    hPen := gdi32.ExtCreatePen(style, width, brush.GetLOGBRUSH(), 0, nil)
    if hPen == 0 {
        panic("Failed to create pen")
    }

    return &Pen{hPen, style, brush}
}

func (this *Pen) Style() uint {
    return this.style
}

func (this *Pen) Brush() *Brush {
    return this.brush
}

func (this *Pen) GetHPEN() w32.HPEN {
    return this.hPen
}

func (this *Pen) Dispose() {
    if this.hPen != 0 {
        gdi32.DeleteObject(w32.HGDIOBJ(this.hPen))
        this.hPen = 0
    }
}