package emminator

type Emitter struct {
	readyState bool
	ons        map[string][]func()
	onces      map[string][]func()
	readies    []func()
}

func NewEmitter() *Emitter {
	return &Emitter{false, make(map[string][]func()), make(map[string][]func()), make([]func(), 0)}
}

func addCallback(cbmap map[string][]func(), message string, callback func()) {
	_, ok := cbmap[message]
	if !ok {
		cbmap[message] = make([]func(), 0)
	}
	cbmap[message] = append(cbmap[message], callback)
}

func (emitter *Emitter) On(message string, callback func()) {
	addCallback(emitter.ons, message, callback)
}

func (emitter *Emitter) Off(message string) {
        _, ok := emitter.ons[message]
	if ok {
		delete(emitter.ons, message)
	}
	_, ok = emitter.onces[message]
	if ok {
		delete(emitter.onces, message)
	}
}

func (emitter *Emitter) Once(message string, callback func()) {
	addCallback(emitter.onces, message, callback)
}

func (emitter *Emitter) Ready(callback func()) {
	if emitter.readyState {
		if callback != nil {
			callback()
		}
		for _, lcallback := range emitter.readies {
			lcallback()
		}
		emitter.readies = make([]func(), 0)
	} else {
		emitter.readies = append(emitter.readies, callback)
	}
}

func callCallbacks(cbmap map[string][]func(), message string, del bool) {
	cblist, ok := cbmap[message]
	if ok {
		for _, callback := range cblist {
			callback()
		}
		if del {
			delete(cbmap, message)
		}
	}

}

func (emitter *Emitter) Emit(message string) {
	if message == "ready" {
		emitter.readyState = true
		emitter.Ready(nil)
		return
	}
	callCallbacks(emitter.ons, message, false)
	callCallbacks(emitter.onces, message, true)
}
