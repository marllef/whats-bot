package whatsapp

type Handler interface {
	HandleError(err error)
}
type TextMessageHandler interface {
	Handler
	HandleTextMessage(message TextMessage)
}
type ImageMessageHandler interface {
	Handler
	HandleImageMessage(message ImageMessage)
}

func (wac *conn) AddHandler(handler Handler) {
	wac.dispatcher.handler = append(wac.dispatcher.handler, handler)
}

func (dp *dispatcher) handle(message interface{}) {
	switch m := message.(type) {
	case error:
		for _, h := range dp.handler {
			h.HandleError(m)
		}
	case TextMessage:
		for _, h := range dp.handler {
			x, ok := h.(TextMessageHandler)
			if !ok {
				continue
			}
			go x.HandleTextMessage(m)
		}
	case ImageMessage:
		for _, h := range dp.handler {
			x, ok := h.(ImageMessageHandler)
			if !ok {
				continue
			}
			go x.HandleImageMessage(m)
		}
	}
}