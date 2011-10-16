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

func (this *GeneralEventManager) Clean() {
    for this.handlers.Len() > 0 {
        this.handlers.Delete(0)
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

func (this *MouseEventManagerA) Clean() {
    for this.handlers.Len() > 0 {
        this.handlers.Delete(0)
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

func (this *DropFilesEventManagerA) Clean() {
    for this.handlers.Len() > 0 {
        this.handlers.Delete(0)
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

func (this *PaintEventManagerA) Clean() {
    for this.handlers.Len() > 0 {
        this.handlers.Delete(0)
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

func (this *LVEndLabelEditEventManagerA) Clean() {
    for this.handlers.Len() > 0 {
        this.handlers.Delete(0)
    }
}

type LVDBLClickEventManagerA struct {
    handlers vector.Vector
}

func (this *LVDBLClickEventManagerA) Fire(sender *ListView, arg *LVDBLClickEventArg) {
    for _, v := range this.handlers {
        if f, ok := v.(LVDBLClickEventHandlerA); ok {
            f(sender, arg)
        }
    }
}

func (this *LVDBLClickEventManagerA) Attach(handler LVDBLClickEventHandlerA) {
    isExists := false
    for _, v := range this.handlers {
        if f, ok := v.(LVDBLClickEventHandlerA); ok {
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

func (this *LVDBLClickEventManagerA) Detach(handler LVDBLClickEventHandlerA) {
    for i, v := range this.handlers {
        if f, ok := v.(LVDBLClickEventHandlerA); ok {
            if f == handler {
                this.handlers.Delete(i)
                break
            }
        }
    }
}

func (this *LVDBLClickEventManagerA) Clean() {
    for this.handlers.Len() > 0 {
        this.handlers.Delete(0)
    }
}

type KeyUpEventManagerA struct {
    handlers vector.Vector
}

func (this *KeyUpEventManagerA) Fire(sender Controller, arg *KeyUpEventArg) {
    for _, v := range this.handlers {
        if f, ok := v.(KeyUpEventHandlerA); ok {
            f(sender, arg)
        }
    }
}

func (this *KeyUpEventManagerA) Attach(handler KeyUpEventHandlerA) {
    isExists := false
    for _, v := range this.handlers {
        if f, ok := v.(KeyUpEventHandlerA); ok {
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

func (this *KeyUpEventManagerA) Detach(handler KeyUpEventHandlerA) {
    for i, v := range this.handlers {
        if f, ok := v.(KeyUpEventHandlerA); ok {
            if f == handler {
                this.handlers.Delete(i)
                break
            }
        }
    }
}

func (this *KeyUpEventManagerA) Clean() {
    for this.handlers.Len() > 0 {
        this.handlers.Delete(0)
    }
}
