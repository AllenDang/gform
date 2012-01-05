package gform

import (
	"errors"
	"unsafe"
	"w32"
	"w32/kernel32"
	"w32/ole32"
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

func NewBitmapFromFile(filepath string, background Color) (*Bitmap, error) {
	var gpBitmap *uintptr
	var err error

	gpBitmap, err = gdiplus.GdipCreateBitmapFromFile(filepath)
	if err != nil {
		return nil, err
	}
	defer gdiplus.GdipDisposeImage(gpBitmap)

	var hbitmap w32.HBITMAP
	hbitmap, err = gdiplus.GdipCreateHBITMAPFromBitmap(gpBitmap, uint32(background))
	if err != nil {
		return nil, err
	}

	return assembleBitmapFromHBITMAP(hbitmap)
}

func NewBitmapFromResource(instance w32.HINSTANCE, resName *uint16, resType *uint16, background Color) (*Bitmap, error) {
	var gpBitmap *uintptr
	var err error
	var hRes w32.HRSRC
	
	hRes, err = kernel32.FindResource(w32.HMODULE(instance), resName, resType)
	if err != nil {
		return nil, err
	}
	resSize := kernel32.SizeofResource(w32.HMODULE(instance), hRes)
	pResData := kernel32.LockResource(kernel32.LoadResource(w32.HMODULE(instance), hRes))
	resBuffer := kernel32.GlobalAlloc(w32.GMEM_MOVEABLE, resSize)
	pResBuffer := kernel32.GlobalLock(resBuffer)
	defer kernel32.GlobalUnlock(resBuffer)
	defer kernel32.GlobalFree(resBuffer)

	kernel32.MoveMemory(pResBuffer, pResData, resSize)

	stream := ole32.CreateStreamOnHGlobal(resBuffer, false)
	defer stream.Release()
	gpBitmap, err = gdiplus.GdipCreateBitmapFromStream(stream)
	if err != nil {
		return nil, err
	}
	
	var hbitmap w32.HBITMAP
	hbitmap, err = gdiplus.GdipCreateHBITMAPFromBitmap(gpBitmap, uint32(background))
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