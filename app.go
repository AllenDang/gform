package gform

import (
	"github.com/AllenDang/w32"
	"unsafe"
)

func Init() {
	gAppInstance = w32.GetModuleHandle("")
	if gAppInstance == 0 {
		panic("Error occurred in App.Init")
	}

	// Initialize the common controls
	var initCtrls w32.INITCOMMONCONTROLSEX
	initCtrls.DwSize = uint32(unsafe.Sizeof(initCtrls))
	initCtrls.DwICC =
		w32.ICC_LISTVIEW_CLASSES | w32.ICC_PROGRESS_CLASS | w32.ICC_TAB_CLASSES |
			w32.ICC_TREEVIEW_CLASSES | w32.ICC_BAR_CLASSES

	w32.InitCommonControlsEx(&initCtrls)
}

func GetAppInstance() w32.HINSTANCE {
	return gAppInstance
}

func PreTranslateMessage(msg *w32.MSG) bool {
	// This functions is called by the MessageLoop. It processes the
	// keyboard accelerator keys and calls Controller.PreTranslateMessage for
	// keyboard and mouse events.

	processed := false

	if (msg.Message >= w32.WM_KEYFIRST && msg.Message <= w32.WM_KEYLAST) ||
		(msg.Message >= w32.WM_MOUSEFIRST && msg.Message <= w32.WM_MOUSELAST) {

		if msg.Hwnd != 0 {
			if controller := GetMsgHandler(msg.Hwnd); controller != nil {
				// Search the chain of parents for pretranslated messages.
				for p := controller; p != nil; p = p.Parent() {
					if processed = p.PreTranslateMessage(msg); processed {
						break
					}
				}
			}
		}
	}

	return processed
}

func RunMainLoop() int {
	var m w32.MSG

	for w32.GetMessage(&m, 0, 0, 0) != 0 {
		if !PreTranslateMessage(&m) {
			w32.TranslateMessage(&m)
			w32.DispatchMessage(&m)
		}
	}

	w32.GdiplusShutdown()

	return int(m.WParam)
}

func Exit() {
	w32.PostQuitMessage(0)
}
