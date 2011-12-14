package gform

import (
    "w32"
)

func RegMsgHandler(controller Controller) {
    if _, isExists := gControllerRegistry[controller.Handle()]; !isExists {
        gControllerRegistry[controller.Handle()] = controller
    }
}

func UnRegMsgHandler(hwnd w32.HWND) {
    delete(gControllerRegistry, hwnd)
}

func GetMsgHandler(hwnd w32.HWND) Controller {
    if controller, isExists := gControllerRegistry[hwnd]; isExists {
        return controller
    }

    return nil
}
