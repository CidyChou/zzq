// main.go
package main

import (
	"cidychou/enjoy-test/fsm"
	"fmt"
)

var (
	Search = fsm.FSMState("寻找目标")
	Attack = fsm.FSMState("攻击")
	Move   = fsm.FSMState("移动")
	Skill  = fsm.FSMState("释放技能")

	TargetDealEvent    = fsm.FSMEvent("目标死亡")
	TargetOutSideEvent = fsm.FSMEvent("目标在范围内")
	TargetInSidEvent   = fsm.FSMEvent("目标在范围外")

	TargetDealHandler = fsm.FSMHandler(func() fsm.FSMState {
		fmt.Println("在移动时目标已死亡,重新搜寻目标")
		return Search
	})

	TargetOutSideHandler = fsm.FSMHandler(func() fsm.FSMState {
		fmt.Println("敌人在攻击范围以外,移动")
		return Move
	})

	// TargetInSideHandler is 目标在攻击范围外
	TargetInSideHandler = fsm.FSMHandler(func() fsm.FSMState {
		fmt.Println("敌人在攻击范围以内,攻击")
		return Attack
	})
)

// 电风扇
type ElectricFan struct {
	*fsm.FSM
}

// 实例化电风扇
func NewElectricFan(initState fsm.FSMState) *ElectricFan {
	return &ElectricFan{
		FSM: fsm.NewFSM(initState),
	}
}

// 入口函数
func main() {
	efan := NewElectricFan(Search) // 初始状态是关闭的
	// 寻找目标
	efan.AddHandler(Search, TargetOutSideEvent, TargetOutSideHandler)
	efan.AddHandler(Search, TargetInSidEvent, TargetInSideHandler)

	// 攻击状态
	efan.AddHandler(Attack, TargetDealEvent, TargetDealHandler)

	// 移动状态
	efan.AddHandler(Move, TargetDealEvent, TargetDealHandler)
	efan.AddHandler(Move, TargetInSidEvent, TargetInSideHandler)

	// 开始测试状态变化
	efan.Call(TargetOutSideEvent)
	efan.Call(TargetInSidEvent)
	efan.Call(TargetDealEvent)
	efan.Call(TargetInSidEvent)
}
