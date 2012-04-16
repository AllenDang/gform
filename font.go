package gform

import (
    "github.com/AllenDang/w32"
    "syscall"
)

const (
    FontBold      byte = 0x01
    FontItalic    byte = 0x02
    FontUnderline byte = 0x04
    FontStrikeOut byte = 0x08
)

func init() {
    DefaultFont = NewFont("MS Shell Dlg 2", 8, 0)
}

type Font struct {
    hfont     w32.HFONT
    family    string
    pointSize int
    style     byte
}

func NewFont(family string, pointSize int, style byte) *Font {
    if style > FontBold|FontItalic|FontUnderline|FontStrikeOut {
        panic("Invalid font style")
    }

    //Retrive screen DPI
    hDC := w32.GetDC(0)
    defer w32.ReleaseDC(0, hDC)
    screenDPIY := w32.GetDeviceCaps(hDC, w32.LOGPIXELSY)

    font := Font{
        family:    family,
        pointSize: pointSize,
        style:     style,
    }

    font.hfont = font.createForDPI(screenDPIY)
    if font.hfont == 0 {
        panic("CreateFontIndirect failed")
    }

    return &font
}

func (this *Font) createForDPI(dpi int) w32.HFONT {
    var lf w32.LOGFONT

    lf.Height = -w32.MulDiv(this.pointSize, dpi, 72)
    if this.style&FontBold > 0 {
        lf.Weight = w32.FW_BOLD
    } else {
        lf.Weight = w32.FW_NORMAL
    }
    if this.style&FontItalic > 0 {
        lf.Italic = 1
    }
    if this.style&FontUnderline > 0 {
        lf.Underline = 1
    }
    if this.style&FontStrikeOut > 0 {
        lf.StrikeOut = 1
    }
    lf.CharSet = w32.DEFAULT_CHARSET
    lf.OutPrecision = w32.OUT_TT_PRECIS
    lf.ClipPrecision = w32.CLIP_DEFAULT_PRECIS
    lf.Quality = w32.CLEARTYPE_QUALITY
    lf.PitchAndFamily = w32.VARIABLE_PITCH | w32.FF_SWISS

    src := syscall.StringToUTF16(this.family)
    dest := lf.FaceName[:]
    copy(dest, src)

    return w32.CreateFontIndirect(&lf)
}

func (this *Font) GetHFONT() w32.HFONT {
    return this.hfont
}

func (this *Font) Bold() bool {
    return this.style&FontBold > 0
}

func (this *Font) Dispose() {
    if this.hfont != 0 {
        w32.DeleteObject(w32.HGDIOBJ(this.hfont))
    }
}

func (this *Font) Family() string {
    return this.family
}

func (this *Font) Italic() bool {
    return this.style&FontItalic > 0
}

func (this *Font) StrikeOut() bool {
    return this.style&FontStrikeOut > 0
}

func (this *Font) Underline() bool {
    return this.style&FontUnderline > 0
}

func (this *Font) Style() byte {
    return this.style
}
