package gform

type EventArg struct {
	sender Controller
	data interface{}
}

func NewEventArg(sender Controller, data interface{}) *EventArg {
	ea := new(EventArg)
	ea.sender, ea.data = sender, data
	return ea
}

func (this *EventArg) Sender() Controller {
	return this.sender
}

func (this *EventArg) Data() interface{} {
	return this.data
}