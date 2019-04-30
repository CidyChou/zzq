package v1

import "sync"

//type FSMState string // 状态
//type FSMEvent string // 事件

//State is 英雄状态机
type HeroFSM struct {
	Mu       sync.Mutex                           // 排他锁
	State    FSMState                             // 当前状态
	Handlers map[FSMState]map[FSMEvent]FSMHandler // 处理地图集，每一个状态都可以出发有限个事件，执行有限个处理
}

// 获取当前状态
func (h *HeroFSM) getHero() Hero {
	return h.State
}

// 设置当前状态
func (h *HeroFSM) setHero(newState Hero) {
	h.State = newState
}

type State int

// iota 初始化后会自动递增
const (
	Alive  State = iota // value --> 0
	Dead                // value --> 1
	Move                // value --> 2
	Attack              // value --> 3
)
