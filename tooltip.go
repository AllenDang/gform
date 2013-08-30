package gform

import (
	"github.com/AllenDang/w32"
	"syscall"
	"unsafe"
)

type ToolTip struct {
	W32Control
}

func NewToolTip(parent Controller) *ToolTip {
	ttip := new(ToolTip)
	ttip.init(parent)

	return ttip
}

func (this *ToolTip) init(parent Controller) {
	this.W32Control.init("tooltips_class32", parent, w32.WS_EX_TOPMOST, w32.WS_POPUP|w32.TTS_NOPREFIX|w32.TTS_ALWAYSTIP)
	w32.SetWindowPos(this.Handle(), w32.HWND_TOPMOST, 0, 0, 0, 0, w32.SWP_NOMOVE|w32.SWP_NOSIZE|w32.SWP_NOACTIVATE)
}

func (this *ToolTip) AddTool(tool Controller, tip string) bool {
	var ti w32.TOOLINFO
	ti.CbSize = uint32(unsafe.Sizeof(ti))
	if tool.Parent() != nil {
		ti.Hwnd = tool.Parent().Handle()
	}
	ti.UFlags = w32.TTF_IDISHWND | w32.TTF_SUBCLASS
	ti.UId = uintptr(tool.Handle())
	ti.LpszText = syscall.StringToUTF16Ptr(tip)

	return w32.SendMessage(this.Handle(), w32.TTM_ADDTOOL, 0, uintptr(unsafe.Pointer(&ti))) != w32.FALSE
}
