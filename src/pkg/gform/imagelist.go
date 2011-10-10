package gform

import (
	"w32"
	"w32/comctl32"
)

type ImageList struct {
	handle w32.HIMAGELIST
}

func NewImageList(cx, cy int, flags uint, cInitial, cGrow int) *ImageList {
	imgl := new(ImageList)
	imgl.handle = comctl32.ImageList_Create(cx, cy, flags, cInitial, cGrow)

	return imgl
}

func (this *ImageList) Handle() w32.HIMAGELIST {
	return this.handle
}

func (this *ImageList) Destroy() bool {
	return comctl32.ImageList_Destroy(this.handle)
}

func (this *ImageList) SetImageCount(uNewCount uint) bool {
	return comctl32.ImageList_SetImageCount(this.handle, uNewCount)
}

func (this *ImageList) GetImageCount() int {
	return comctl32.ImageList_GetImageCount(this.handle)
}

func (this *ImageList) AddIcon(icon *Icon) int {
	return comctl32.ImageList_AddIcon(this.handle, icon.Handle())
}

func (this *ImageList) RemoveAll() bool {
	return comctl32.ImageList_RemoveAll(this.handle)
}

func (this *ImageList) Remove(i int) bool {
	return comctl32.ImageList_Remove(this.handle, i)
}