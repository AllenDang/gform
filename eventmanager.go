package gform

type EventHandler func(arg *EventArg)

type EventManager struct {
    handler EventHandler
}

func (this *EventManager) Fire(arg *EventArg) {
    if this.handler != nil {
        this.handler(arg)
    }
}

func (this *EventManager) Bind(handler EventHandler) {
    this.handler = handler
}
