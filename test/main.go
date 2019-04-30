// main.go
package main

import (
	model "dc-sz/dc/zzq/model/attr"
	"dc-sz/dc/zzq/test/fsm"
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

	TargetDealHandler = fsm.FSMHandler(func(hero model.Hero, enemy model.Hero) fsm.FSMState {
		fmt.Println("Hero:", hero, enemy)
		fmt.Println("在移动时目标已死亡,重新搜寻目标")
		return Search
	})

	TargetOutSideHandler = fsm.FSMHandler(func(hero model.Hero, enemy model.Hero) fsm.FSMState {
		fmt.Println("Hero:", hero, enemy)
		fmt.Println("敌人在攻击范围以外,移动")
		return Move
	})

	// TargetInSideHandler is 目标在攻击范围外
	TargetInSideHandler = fsm.FSMHandler(func(hero model.Hero, enemy model.Hero) fsm.FSMState {
		fmt.Println("Hero:", hero, enemy)
		fmt.Println("敌人在攻击范围以内,攻击")
		return Attack
	})
)

// 电风扇
type Chessboard struct {
	*fsm.FSM
}

// 实例化棋盘
func NewChessboard(initState fsm.FSMState) *Chessboard {
	return &Chessboard{
		FSM: fsm.NewFSM(initState),
	}
}

// 入口函数
func main() {
	efan := NewChessboard(Search) // 初始状态
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
