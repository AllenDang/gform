package gform

import (
	"errors"
	"unsafe"
	"w32"
	"w32/gdi32"
	"w32/gdiplus"
)

type Bitmap struct {
	handle w32.HBITMAP
	width, height int
}

func assembleBitmapFromHBITMAP(hbitmap w32.HBITMAP) (*Bitmap, error) {
	var dib w32.DIBSECTION
	if gdi32.GetObject(w32.HGDIOBJ(hbitmap), unsafe.Sizeof(dib), unsafe.Pointer(&dib)) == 0 {
		return nil, errors.New("GetObject for HBITMAP failed")
	}

	return &Bitmap{
		handle: hbitmap,
		width: dib.DsBmih.BiWidth,
		height: dib.DsBmih.BiHeight,
	}, nil
}

func NewBitmapFromFile(filepath string) (*Bitmap, error) {
	var gpBitmap *uintptr
	var err error

	gpBitmap, err = gdiplus.GdipCreateBitmapFromFile(filepath)
	if err != nil {
		return nil, err
	}
	defer gdiplus.GdipDisposeImage(gpBitmap)

	var hbitmap w32.HBITMAP
	hbitmap, err = gdiplus.GdipCreateHBITMAPFromBitmap(gpBitmap)
	if err != nil {
		return nil, err
	}

	return assembleBitmapFromHBITMAP(hbitmap)
}

func NewBitmapFromResId(instance w32.HINSTANCE, resId *uint16) (*Bitmap, error) {
	var gpBitmap *uintptr
	var err error
	
	gpBitmap, err = gdiplus.GdipCreateBitmapFromResource(instance, resId)
	if err != nil {
		return nil, err
	}
	defer gdiplus.GdipDisposeImage(gpBitmap)
	
	var hbitmap w32.HBITMAP
	hbitmap, err = gdiplus.GdipCreateHBITMAPFromBitmap(gpBitmap)
	if err != nil {
		return nil, err
	}

	return assembleBitmapFromHBITMAP(hbitmap)
}

func (this *Bitmap) Dispose() {
	if this.handle != 0 {
		gdi32.DeleteObject(w32.HGDIOBJ(this.handle))
		this.handle = 0
	}
}

func (this *Bitmap) GetHBITMAP() w32.HBITMAP {
	return this.handle
}

func (this *Bitmap) Size() (int, int) {
	return this.width, this.height
}