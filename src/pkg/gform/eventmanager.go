package gform

import (
    "container/vector"
)

type GeneralEventManager struct {
    handlers vector.Vector
}

func (this *GeneralEventManager) Fire(sender Controller) {
    for _, v := range this.handlers {
        if f, ok := v.(GeneralEventHandler); ok {
            f(sender)
        }
    }
}

func (this *GeneralEventManager) Attach(handler GeneralEventHandler) {
    isExists := false
    for _, v := range this.handlers {
        if f, ok := v.(GeneralEventHandler); ok {
            if f == handler {
                isExists = true
                break
            }
        }
    }

    if !isExists {
        this.handlers.Push(handler)
    }
}

func (this *GeneralEventManager) Detach(handler GeneralEventHandler) {
    for i, v := range this.handlers {
        if f, ok := v.(GeneralEventHandler); ok {
            if f == handler {
                this.handlers.Delete(i)
                break
            }
        }
    }
}

type MouseEventManagerA struct {
    handlers vector.Vector
}

func (this *MouseEventManagerA) Fire(sender Controller, arg *MouseEventArg) {
    for _, v := range this.handlers {
        if f, ok := v.(MouseEventHandlerA); ok {
            f(sender, arg)
        }
    }
}

func (this *MouseEventManagerA) Attach(handler MouseEventHandlerA) {
    isExists := false
    for _, v := range this.handlers {
        if f, ok := v.(MouseEventHandlerA); ok {
            if f == handler {
                isExists = true
                break
            }
        }
    }

    if !isExists {
        this.handlers.Push(handler)
    }
}

func (this *MouseEventManagerA) Detach(handler MouseEventHandlerA) {
    for i, v := range this.handlers {
        if f, ok := v.(MouseEventHandlerA); ok {
            if f == handler {
                this.handlers.Delete(i)
                break
            }
        }
    }
}

type DropFilesEventManagerA struct {
    handlers vector.Vector
}

func (this *DropFilesEventManagerA) Fire(sender Controller, arg *DropFilesEventArg) {
    for _, v := range this.handlers {
        if f, ok := v.(DropFilesEventHandlerA); ok {
            f(sender, arg)
        }
    }
}

func (this *DropFilesEventManagerA) Attach(handler DropFilesEventHandlerA) {
    isExists := false
    for _, v := range this.handlers {
        if f, ok := v.(DropFilesEventHandlerA); ok {
            if f == handler {
                isExists = true
                break
            }
        }
    }

    if !isExists {
        this.handlers.Push(handler)
    }
}

func (this *DropFilesEventManagerA) Detach(handler DropFilesEventHandlerA) {
    for i, v := range this.handlers {
        if f, ok := v.(DropFilesEventHandlerA); ok {
            if f == handler {
                this.handlers.Delete(i)
                break
            }
        }
    }
}


type PaintEventManagerA struct {
    handlers vector.Vector
}

func (this *PaintEventManagerA) Fire(sender Controller, arg *PaintEventArg) {
    for _, v := range this.handlers {
        if f, ok := v.(PaintEventHandlerA); ok {
            f(sender, arg)
        }
    }
}

func (this *PaintEventManagerA) Attach(handler PaintEventHandlerA) {
    isExists := false
    for _, v := range this.handlers {
        if f, ok := v.(PaintEventHandlerA); ok {
            if f == handler {
                isExists = true
                break
            }
        }
    }

    if !isExists {
        this.handlers.Push(handler)
    }
}

func (this *PaintEventManagerA) Detach(handler PaintEventHandlerA) {
    for i, v := range this.handlers {
        if f, ok := v.(PaintEventHandlerA); ok {
            if f == handler {
                this.handlers.Delete(i)
                break
            }
        }
    }
}

type LVEndLabelEditEventManagerA struct {
    handlers vector.Vector
}

func (this *LVEndLabelEditEventManagerA) Fire(sender *ListView, arg *LVEndLabelEditEventArg) {
    for _, v := range this.handlers {
        if f, ok := v.(LVEndLabelEditEventHandlerA); ok {
            f(sender, arg)
        }
    }
}

func (this *LVEndLabelEditEventManagerA) Attach(handler LVEndLabelEditEventHandlerA) {
    isExists := false
    for _, v := range this.handlers {
        if f, ok := v.(LVEndLabelEditEventHandlerA); ok {
            if f == handler {
                isExists = true
                break
            }
        }
    }

    if !isExists {
        this.handlers.Push(handler)
    }
}

func (this *LVEndLabelEditEventManagerA) Detach(handler LVEndLabelEditEventHandlerA) {
    for i, v := range this.handlers {
        if f, ok := v.(LVEndLabelEditEventHandlerA); ok {
            if f == handler {
                this.handlers.Delete(i)
                break
            }
        }
    }
}