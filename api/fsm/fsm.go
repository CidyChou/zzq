package fsm

import (
	model "dc-sz/dc/zzq/model/attr"
	"fmt"
	"sync"
)

type FSMState string                                             // 状态
type FSMEvent string                                             // 事件
type FSMHandler func(hero model.Hero, enemy model.Hero) FSMState // 处理方法，并返回新的状态

// 有限状态机
type FSM struct {
	mu       sync.Mutex                           // 排他锁
	state    FSMState                             // 当前状态
	handlers map[FSMState]map[FSMEvent]FSMHandler // 处理地图集，每一个状态都可以出发有限个事件，执行有限个处理
}

// 获取当前状态
func (f *FSM) getState() FSMState {
	return f.state
}

// 设置当前状态
func (f *FSM) setState(newState FSMState) {
	f.state = newState
}

// 某状态添加事件处理方法
func (f *FSM) AddHandler(state FSMState, event FSMEvent, handler FSMHandler) *FSM {
	if _, ok := f.handlers[state]; !ok {
		f.handlers[state] = make(map[FSMEvent]FSMHandler)
	}
	if _, ok := f.handlers[state][event]; ok {
		fmt.Printf("[警告] 状态(%s)事件(%s)已定义过", state, event)
	}
	f.handlers[state][event] = handler
	return f
}

// Call is 事件处理
func (f *FSM) Call(event FSMEvent) FSMState {
	f.mu.Lock()
	defer f.mu.Unlock()
	events := f.handlers[f.getState()]
	if events == nil {
		return f.getState()
	}
	if fn, ok := events[event]; ok {
		oldState := f.getState()

		heroA := &model.Hero{1, 700, 0, 1, 100, 5, 1, 1, [2]int{1, 1}}
		heroB := &model.Hero{2, 500, 0, 1, 80, 5, 1, 1, [2]int{4, 4}}

		f.setState(fn(*heroA, *heroB))
		newState := f.getState()
		fmt.Println("状态从 [", oldState, "] 变成 [", newState, "]")
	}
	return f.getState()
}

// 实例化FSM
func NewFSM(initState FSMState) *FSM {
	return &FSM{
		state:    initState,
		handlers: make(map[FSMState]map[FSMEvent]FSMHandler),
	}
}
