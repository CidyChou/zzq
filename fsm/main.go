// main.go
package main

import (
	"fmt"

	fsm "dc-sz/dc/zzq/fsm/index"
	model "dc-sz/dc/zzq/model/attr"
)

var (
	Search = fsm.FSMState("寻找目标")
	Attack = fsm.FSMState("攻击")
	Move   = fsm.FSMState("移动")
	Skill  = fsm.FSMState("释放技能")

	TargetDealEvent    = fsm.FSMEvent("发现目标死亡")
	TargetOutSideEvent = fsm.FSMEvent("发现目标在范围内")
	TargetInSidEvent   = fsm.FSMEvent("发现目标在范围外")

	TargetDealHandler = fsm.FSMHandler(func(hero model.Hero, enemy model.Hero) fsm.FSMState {
		fmt.Println("目标已死亡 => 重新搜寻目标")
		return Search
	})

	TargetOutSideHandler = fsm.FSMHandler(func(hero model.Hero, enemy model.Hero) fsm.FSMState {
		fmt.Println("敌人在攻击范围以外 => 移动")
		return Move
	})

	// TargetInSideHandler is 目标在攻击范围外
	TargetInSideHandler = fsm.FSMHandler(func(hero model.Hero, enemy model.Hero) fsm.FSMState {
		//fmt.Println("Hero:", hero, enemy)
		fmt.Println("敌人在攻击范围以内 => 攻击")
		return Attack
	})
)

// Chessboard is 棋盘
type Chessboard struct {
	*fsm.FSM
}

// NewChessboard is 实例化棋盘
func NewChessboard(initState fsm.FSMState) *Chessboard {
	return &Chessboard{
		FSM: fsm.NewFSM(initState),
	}
}

// 入口函数
func main() {
	chessboard := NewChessboard(Search) // 初始状态
	fmt.Println("初始状态 => 搜寻目标")

	// 寻找目标
	chessboard.AddHandler(Search, TargetOutSideEvent, TargetOutSideHandler)
	chessboard.AddHandler(Search, TargetInSidEvent, TargetInSideHandler)

	// 攻击状态
	chessboard.AddHandler(Attack, TargetDealEvent, TargetDealHandler)

	// 移动状态
	chessboard.AddHandler(Move, TargetDealEvent, TargetDealHandler)
	chessboard.AddHandler(Move, TargetInSidEvent, TargetInSideHandler)

	// 开始测试状态变化
	chessboard.Call(TargetOutSideEvent)
	chessboard.Call(TargetInSidEvent)
	chessboard.Call(TargetDealEvent)
	chessboard.Call(TargetInSidEvent)
}
