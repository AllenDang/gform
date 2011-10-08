package gform

import (
	"os"
	"fmt"
	"w32"
	"w32/user32"
	"w32/shell32"
	"syscall"
)

type Icon struct {
	handle w32.HICON
}

func NewIconFromFile(path string) (*Icon, os.Error) {
	ico := new(Icon)
	var err os.Error
	if ico.handle = user32.LoadIcon(0, syscall.StringToUTF16Ptr(path)); ico.handle == 0 {
		err = os.NewError(fmt.Sprintf("Cannot load icon from %s", path))

	}

	return ico, err
}

func NewIconFromResource(instance w32.HINSTANCE, resId uint16) (*Icon, os.Error) {
	ico := new(Icon)
	var err os.Error
	if ico.handle = user32.LoadIcon(instance, w32.MakeIntResource(resId)); ico.handle == 0 {
		err = os.NewError(fmt.Sprintf("Cannot load icon from resource with id %v", resId))
	}

	return ico, err
}

func ExtractIcon(fileName string, index int) (*Icon, os.Error) {
	ico := new(Icon)
	var err os.Error
	if ico.handle = shell32.ExtractIcon(fileName, index); ico.handle == 0 || ico.handle == 1 {
		err = os.NewError(fmt.Sprintf("Cannot extract icon from %s at index %v", fileName, index))
	}

	return ico, err
}

func (this *Icon) Destroy() bool {
	return user32.DestroyIcon(this.handle)
}

func (this *Icon) Handle() w32.HICON {
	return this.handle
}