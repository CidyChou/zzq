// package model

// import (
// 	"sync"
// )

// type State struct {
// 	mu       sync.Mutex                           // 排他锁
// 	state    Hero                                 // 当前状态
// 	handlers map[FSMState]map[FSMEvent]FSMHandler // 处理地图集，每一个状态都可以出发有限个事件，执行有限个处理
// }
