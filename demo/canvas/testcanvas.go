package main

import (
	"gform"
	"w32"
	"syscall"
)

const IDR_PNG1 = 100

func onpaint(arg *gform.EventArg) {
	if data, ok := arg.Data().(*gform.PaintEventData); ok {
		if bmp, err := gform.NewBitmapFromResId(
			gform.GetAppInstance(), 
			w32.MakeIntResource(IDR_PNG1), 
			syscall.StringToUTF16Ptr("PNG"), 
			gform.RGB(255, 0, 0)); err == nil {
			data.Canvas.DrawBitmap(bmp, 10, 10)
			bmp.Dispose()
		} else {
			println(err.Error())
		}
	}
}

func main() {
	gform.Init()

	mf := gform.NewForm(nil)
	mf.SetSize(300, 200)
	mf.Center()

	mf.OnPaint().Attach(onpaint)

	mf.Show()

	gform.RunMainLoop()
}