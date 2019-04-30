package v1

import (
	"fmt"
)

// 处理方法，并返回新的状态
type FSMHandler func() Hero

// 某状态添加事件处理方法
func (h *HeroFSM) AddHandler(hero Hero, event IHero, handler FSMHandler) *HeroFSM {

	if _, ok := h.Handlers[hero]; !ok {
		h.Handlers[hero] = make(map[IHero]FSMHandler)
	}

	if _, ok := h.Handlers[hero][event]; ok {
		fmt.Printf("[警告] 事件已定义过")
	}
	h.Handlers[hero][event] = handler

	return nil
}

// Call is 事件处理
func (h *HeroFSM) Call(event IHero) Hero {
	h.Mu.Lock()
	defer h.Mu.Unlock()

	events := h.Handlers[h.getHero()]

	if events == nil {
		return h.getHero()
	}

	if hn, ok := events[event]; ok {
		oldState := h.getHero()
		h.setHero(hn())
		newState := h.getHero()
		fmt.Println("状态从 [", oldState, "] 变成 [", newState, "]")
	}
	return h.getHero()
}

func New(initState Hero) *HeroFSM {
	return &HeroFSM{
		State:    initState,
		Handlers: make(map[Hero]map[IHero]FSMHandler),
	}
}
