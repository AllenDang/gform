package gform

import (
    "w32"
)

func init() {
    gControllerRegistry = make(map[w32.HWND]Controller)
    gRegisteredClasses = make([]string, 0)
}
