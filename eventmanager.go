package gform

type EventHandler func(arg *EventArg)

type EventManager struct {
    handlers []EventHandler
}

func (this *EventManager) Fire(arg *EventArg) {
    if this.handlers != nil {
        for _, v := range this.handlers {
            v(arg)
        }
    }
}

func (this *EventManager) Attach(handler EventHandler) {
    if this.handlers == nil {
        this.handlers = make([]EventHandler, 0)
    }

    isExists := false
    for _, v := range this.handlers {
        if &v == &handler {
            isExists = true
            break
        }
    }

    if !isExists {
        this.handlers = append(this.handlers, handler)
    }
}

func (this *EventManager) Detach(handler EventHandler) {
    if this.handlers != nil {
        for i, v := range this.handlers {
            if &v == &handler {
                this.handlers = append(this.handlers[:i], this.handlers[i+1:]...)
                break
            }
        }
    }
}

func (this *EventManager) Clean() {
    if this.handlers != nil {
        for len(this.handlers) > 0 {
            this.handlers = append(this.handlers[:0], this.handlers[1:]...)
        }
    }
}