package gform

type GeneralEventManager struct {
    handlers []GeneralEventHandler
}

func (this *GeneralEventManager) Fire(sender Controller) {
    if this.handlers != nil {
        for _, v := range this.handlers {
            v(sender)
        }
    }
}

func (this *GeneralEventManager) Attach(handler GeneralEventHandler) {
    if this.handlers == nil {
        this.handlers = make([]GeneralEventHandler, 0)
    }

    isExists := false
    for _, v := range this.handlers {
        if v == handler {
            isExists = true
            break
        }
    }

    if !isExists {
        this.handlers = append(this.handlers, handler)
    }
}

func (this *GeneralEventManager) Detach(handler GeneralEventHandler) {
    if this.handlers != nil {
        for i, v := range this.handlers {
            if v == handler {
                this.handlers = append(this.handlers[:i], this.handlers[i+1:]...)
                break
            }
        }
    }
}

func (this *GeneralEventManager) Clean() {
    if this.handlers != nil {
        for len(this.handlers) > 0 {
            this.handlers = append(this.handlers[:0], this.handlers[1:]...)
        }
    }
}

type MouseEventManagerA struct {
    handlers []MouseEventHandlerA
}

func (this *MouseEventManagerA) Fire(sender Controller, arg *MouseEventArg) {
    if this.handlers != nil {
        for _, v := range this.handlers {
            v(sender, arg)
        }
    }
}

func (this *MouseEventManagerA) Attach(handler MouseEventHandlerA) {
    if this.handlers == nil {
        this.handlers = make([]MouseEventHandlerA, 0)
    }
    isExists := false
    for _, v := range this.handlers {
        if v == handler {
            isExists = true
            break
        }
    }

    if !isExists {
        this.handlers = append(this.handlers, handler)
    }
}

func (this *MouseEventManagerA) Detach(handler MouseEventHandlerA) {
    if this.handlers != nil {
        for i, v := range this.handlers {
            if v == handler {
                this.handlers = append(this.handlers[:i], this.handlers[i+1:]...)
                break
            }
        }
    }
}

func (this *MouseEventManagerA) Clean() {
    if this.handlers != nil {
        for len(this.handlers) > 0 {
            this.handlers = append(this.handlers[:0], this.handlers[1:]...)
        }
    }
}

type DropFilesEventManagerA struct {
    handlers []DropFilesEventHandlerA
}

func (this *DropFilesEventManagerA) Fire(sender Controller, arg *DropFilesEventArg) {
    if this.handlers != nil {
        for _, v := range this.handlers {
            v(sender, arg)
        }
    }
}

func (this *DropFilesEventManagerA) Attach(handler DropFilesEventHandlerA) {
    if this.handlers == nil {
        this.handlers = make([]DropFilesEventHandlerA, 0)
    }
    isExists := false
    for _, v := range this.handlers {
        if v == handler {
            isExists = true
            break
        }
    }

    if !isExists {
        this.handlers = append(this.handlers, handler)
    }
}

func (this *DropFilesEventManagerA) Detach(handler DropFilesEventHandlerA) {
    if this.handlers != nil {
        for i, v := range this.handlers {
            if v == handler {
                this.handlers = append(this.handlers[:i], this.handlers[i+1:]...)
                break
            }
        }
    }
}

func (this *DropFilesEventManagerA) Clean() {
    if this.handlers != nil {
        for len(this.handlers) > 0 {
            this.handlers = append(this.handlers[:0], this.handlers[1:]...)
        }
    }
}

type PaintEventManagerA struct {
    handlers []PaintEventHandlerA
}

func (this *PaintEventManagerA) Fire(sender Controller, arg *PaintEventArg) {
    if this.handlers != nil {
        for _, v := range this.handlers {
            v(sender, arg)
        }
    }
}

func (this *PaintEventManagerA) Attach(handler PaintEventHandlerA) {
    if this.handlers == nil {
        this.handlers = make([]PaintEventHandlerA, 0)
    }
    isExists := false
    for _, v := range this.handlers {
        if v == handler {
            isExists = true
            break
        }
    }

    if !isExists {
        this.handlers = append(this.handlers, handler)
    }
}

func (this *PaintEventManagerA) Detach(handler PaintEventHandlerA) {
    if this.handlers != nil {
        for i, v := range this.handlers {
            if v == handler {
                this.handlers = append(this.handlers[:i], this.handlers[i+1:]...)
                break
            }
        }
    }
}

func (this *PaintEventManagerA) Clean() {
    if this.handlers != nil {
        for len(this.handlers) > 0 {
            this.handlers = append(this.handlers[:0], this.handlers[1:]...)
        }
    }
}

type LVEndLabelEditEventManagerA struct {
    handlers []LVEndLabelEditEventHandlerA
}

func (this *LVEndLabelEditEventManagerA) Fire(sender *ListView, arg *LVEndLabelEditEventArg) {
    if this.handlers != nil {
        for _, v := range this.handlers {
            v(sender, arg)
        }
    }
}

func (this *LVEndLabelEditEventManagerA) Attach(handler LVEndLabelEditEventHandlerA) {
    if this.handlers == nil {
        this.handlers = make([]LVEndLabelEditEventHandlerA, 0)
    }
    isExists := false
    for _, v := range this.handlers {
        if v == handler {
            isExists = true
            break
        }
    }

    if !isExists {
        this.handlers = append(this.handlers, handler)
    }
}

func (this *LVEndLabelEditEventManagerA) Detach(handler LVEndLabelEditEventHandlerA) {
    if this.handlers != nil {
        for i, v := range this.handlers {
            if v == handler {
                this.handlers = append(this.handlers[:i], this.handlers[i+1:]...)
                break
            }
        }
    }
}

func (this *LVEndLabelEditEventManagerA) Clean() {
    if this.handlers != nil {
        for len(this.handlers) > 0 {
            this.handlers = append(this.handlers[:0], this.handlers[1:]...)
        }
    }
}

type LVDBLClickEventManagerA struct {
    handlers []LVDBLClickEventHandlerA
}

func (this *LVDBLClickEventManagerA) Fire(sender *ListView, arg *LVDBLClickEventArg) {
    if this.handlers != nil {
        for _, v := range this.handlers {
            v(sender, arg)
        }
    }
}

func (this *LVDBLClickEventManagerA) Attach(handler LVDBLClickEventHandlerA) {
    if this.handlers == nil {
        this.handlers = make([]LVDBLClickEventHandlerA, 0)
    }
    isExists := false
    for _, v := range this.handlers {
        if v == handler {
            isExists = true
            break
        }
    }

    if !isExists {
        this.handlers = append(this.handlers, handler)
    }
}

func (this *LVDBLClickEventManagerA) Detach(handler LVDBLClickEventHandlerA) {
    if this.handlers != nil {
        for i, v := range this.handlers {
            if v == handler {
                this.handlers = append(this.handlers[:i], this.handlers[i+1:]...)
                break
            }
        }
    }
}

func (this *LVDBLClickEventManagerA) Clean() {
    if this.handlers != nil {
        for len(this.handlers) > 0 {
            this.handlers = append(this.handlers[:0], this.handlers[1:]...)
        }
    }
}

type KeyUpEventManagerA struct {
    handlers []KeyUpEventHandlerA
}

func (this *KeyUpEventManagerA) Fire(sender Controller, arg *KeyUpEventArg) {
    if this.handlers != nil {
        for _, v := range this.handlers {
            v(sender, arg)
        }
    }
}

func (this *KeyUpEventManagerA) Attach(handler KeyUpEventHandlerA) {
    if this.handlers == nil {
        this.handlers = make([]KeyUpEventHandlerA, 0)
    }
    isExists := false
    for _, v := range this.handlers {
        if v == handler {
            isExists = true
            break
        }
    }

    if !isExists {
        this.handlers = append(this.handlers, handler)
    }
}

func (this *KeyUpEventManagerA) Detach(handler KeyUpEventHandlerA) {
    if this.handlers != nil {
        for i, v := range this.handlers {
            if v == handler {
                this.handlers = append(this.handlers[:i], this.handlers[i+1:]...)
                break
            }
        }
    }
}

func (this *KeyUpEventManagerA) Clean() {
    if this.handlers != nil {
        for len(this.handlers) > 0 {
            this.handlers = append(this.handlers[:0], this.handlers[1:]...)
        }
    }
}
