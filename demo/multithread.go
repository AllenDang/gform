package main

import (
    "gform"
    "time"
)

var (
    pb  *gform.ProgressBar
    btn *gform.PushButton
)

func onclick(arg *gform.EventArg) {
    go setProgress()
}

func setProgress() {
    btn.SetEnabled(false)
    for i := 0; i < 100; i++ {
        pb.SetValue(uint(i))
        time.Sleep(50 * 1E6)
    }
    btn.SetEnabled(true)
    pb.SetValue(0)
}

func main() {
    gform.Init()

    mw := gform.NewForm(nil)
    mw.SetPos(300, 100)
    mw.SetSize(500, 300)
    mw.SetCaption("Multi thread demo")

    btn = gform.NewPushButton(mw)
    btn.SetPos(10, 10)
    btn.SetCaption("Click me")
    btn.OnLBUp().Attach(onclick)

    pb = gform.NewProgressBar(mw)
    pb.SetPos(10, 40)
    pb.SetSize(300, 25)

    mw.Show()

    gform.RunMainLoop()
}
