package broadcast

import (
	"sync"
	"time"
)

const (
	available = iota
	used
	toClose
)

type channel struct {
	c     chan interface{}
	state int
}

func NewChannel() *channel {
	ret := &channel{
		c:     make(chan interface{}),
		state: available,
	}
	return ret
}

type broadcaster struct {
	done                     chan struct{}
	newSubscribeInput        chan int
	newSubscribeInputAnswer  chan int
	newSubscribeOutput       chan int
	newSubscribeOutputAnswer chan int
	inputs                   []*channel
	outputs                  []*channel
	sync.Mutex
}

func (b *broadcaster) dispatch() {
	for {
		needToClose := false
		select {
		case <-b.done:
			b.done <- struct{}{}
			return
		case <-b.newSubscribeInput:
			found := false
			for i := range b.inputs {
				if b.inputs[i].state == available {
					b.inputs[i].state = used
					b.newSubscribeInputAnswer <- i
					found = true
					break
				}
			}
			if !found {
				b.inputs = append(b.inputs, NewChannel())
				b.inputs[(len(b.inputs) - 1)].state = used
				b.newSubscribeInputAnswer <- (len(b.inputs) - 1)
			}
		case <-b.newSubscribeOutput:
			found := false
			for i := range b.outputs {
				if b.outputs[i].state == available {
					b.outputs[i].state = used
					b.newSubscribeOutputAnswer <- i
					found = true
					break
				}
			}
			if !found {
				b.outputs = append(b.outputs, NewChannel())
				b.outputs[(len(b.outputs) - 1)].state = used
				b.newSubscribeOutputAnswer <- (len(b.outputs) - 1)
			}
		default:
			for i := range b.inputs {
				select {
				case d, ok := <-b.inputs[i].c:
					if !ok {
						needToClose = true
						b.inputs[i].state = toClose
					} else {
						for o := range b.outputs {
							// fmt.Println(o, " ", b.outputs[o].state)
							select {
							// FIXME add unSubcribe method for ouputs channel
							case _, ok := <-b.outputs[o].c:
								if !ok {
									needToClose = true
									b.outputs[o].state = toClose
								}
							default:
								if b.outputs[o].state == used {
									b.outputs[o].c <- d
								}
							}
						}
					}
				case <-time.After(1 * time.Millisecond):
					break
				}
			}
			if needToClose {
				for i := range b.inputs {
					if b.inputs[i].state == toClose {
						b.inputs[i] = NewChannel()
					}
				}
			}
		}
	}
}

func New() *broadcaster {
	ret := &broadcaster{
		done:                     make(chan struct{}),
		newSubscribeInput:        make(chan int),
		newSubscribeOutput:       make(chan int),
		newSubscribeInputAnswer:  make(chan int),
		newSubscribeOutputAnswer: make(chan int),
		inputs:  make([]*channel, 1),
		outputs: make([]*channel, 1),
	}
	for i := range ret.inputs {
		ret.inputs[i] = NewChannel()
	}
	for i := range ret.outputs {
		ret.outputs[i] = NewChannel()
	}
	go ret.dispatch()
	return ret
}

func (b *broadcaster) SubscribeInput() chan<- interface{} {
	b.Lock()
	defer b.Unlock()
	b.newSubscribeInput <- 0
	id := <-b.newSubscribeInputAnswer
	return b.inputs[id].c
}

func (b *broadcaster) SubscribeOuput() (int, <-chan interface{}) {
	b.Lock()
	defer b.Unlock()
	b.newSubscribeOutput <- 0
	id := <-b.newSubscribeOutputAnswer
	return id, b.outputs[id].c
}

func (b *broadcaster) NbOutputChannel() (size int) {
	b.Lock()
	defer b.Unlock()
	size = 0

	for i := range b.outputs {
		if b.outputs[i].state == used {
			size++
		}
	}
	return
}

func (b *broadcaster) Close() {
	b.done <- struct{}{}
	<-b.done
	close(b.done)
	for i := range b.outputs {
		if b.outputs[i].state == used {
			close(b.outputs[i].c)
		}
	}
}
