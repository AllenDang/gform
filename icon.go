package gform

import (
	"errors"
	"fmt"
	"github.com/AllenDang/w32"
	"github.com/AllenDang/w32/shell32"
	"github.com/AllenDang/w32/user32"
	"syscall"
)

type Icon struct {
	handle w32.HICON
}

func NewIconFromFile(path string) (*Icon, error) {
	ico := new(Icon)
	var err error
	if ico.handle = user32.LoadIcon(0, syscall.StringToUTF16Ptr(path)); ico.handle == 0 {
		err = errors.New(fmt.Sprintf("Cannot load icon from %s", path))

	}

	return ico, err
}

func NewIconFromResource(instance w32.HINSTANCE, resId uint16) (*Icon, error) {
	ico := new(Icon)
	var err error
	if ico.handle = user32.LoadIcon(instance, w32.MakeIntResource(resId)); ico.handle == 0 {
		err = errors.New(fmt.Sprintf("Cannot load icon from resource with id %v", resId))
	}

	return ico, err
}

func ExtractIcon(fileName string, index int) (*Icon, error) {
	ico := new(Icon)
	var err error
	if ico.handle = shell32.ExtractIcon(fileName, index); ico.handle == 0 || ico.handle == 1 {
		err = errors.New(fmt.Sprintf("Cannot extract icon from %s at index %v", fileName, index))
	}

	return ico, err
}

func (this *Icon) Destroy() bool {
	return user32.DestroyIcon(this.handle)
}

func (this *Icon) Handle() w32.HICON {
	return this.handle
}
