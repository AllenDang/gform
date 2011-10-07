package gform

import (
	"w32"
	"w32/user32"
	"syscall"
)

type Icon struct {
	handle w32.HICON
}

func NewIconFromFile(path string) *Icon {
	ico := new(Icon)
	if ico.handle = user32.LoadIcon(0, syscall.StringToUTF16Ptr(path)); ico.handle == 0 {
		panic("Cannot load icon from " + path)
	}

	return ico
}

func NewIconFromResource(instance w32.HINSTANCE, resId uint16) *Icon {
	ico := new(Icon)
	if ico.handle = user32.LoadIcon(instance, w32.MakeIntResource(resId)); ico.handle == 0 {
		panic("Cannot load icon from resource")
	}

	return ico
}

func (this *Icon) Handle() w32.HICON {
	return this.handle
}