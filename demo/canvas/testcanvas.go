package main

import (
	"gform"
)

func onpaint(arg *gform.EventArg) {
	if data, ok := arg.Data().(*gform.PaintEventData); ok {
		if bmp, err := gform.NewBitmapFromFile("close.png", gform.RGB(255, 0, 0)); err == nil {
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