// main.go
package main

import (
	"fmt"
)

var (
	Attack = FSMState("攻击")
	Move   = FSMState("移动")
	Skill  = FSMState("释放技能")

	AttackEvent = FSMEvent("攻击敌方英雄")
	MoveEvent   = FSMEvent("向敌方英雄移动")
	SkillEvent  = FSMEvent("释放技能")

	AttackHandler = FSMHandler(func() FSMState {
		fmt.Println("已攻击敌方英雄")
		return Attack
	})
	MoveHandler = FSMHandler(func() FSMState {
		fmt.Println("已向敌方英雄移动")
		return Move
	})
	SkillHandler = FSMHandler(func() FSMState {
		fmt.Println("已释放技能")
		return Skill
	})
)

// 电风扇
type ElectricFan struct {
	*FSM
}

// 实例化电风扇
func NewElectricFan(initState FSMState) *ElectricFan {
	return &ElectricFan{
		FSM: NewFSM(initState),
	}
}

// 入口函数
func main() {

	efan := NewElectricFan(Poweroff) // 初始状态是关闭的
	// 关闭状态
	efan.AddHandler(Poweroff, PowerOffEvent, PowerOffHandler)
	efan.AddHandler(Poweroff, FirstGearEvent, FirstGearHandler)
	efan.AddHandler(Poweroff, SecondGearEvent, SecondGearHandler)
	efan.AddHandler(Poweroff, ThirdGearEvent, ThirdGearHandler)
	// 1档状态
	efan.AddHandler(FirstGear, PowerOffEvent, PowerOffHandler)
	efan.AddHandler(FirstGear, FirstGearEvent, FirstGearHandler)
	efan.AddHandler(FirstGear, SecondGearEvent, SecondGearHandler)
	efan.AddHandler(FirstGear, ThirdGearEvent, ThirdGearHandler)
	// 2档状态
	efan.AddHandler(SecondGear, PowerOffEvent, PowerOffHandler)
	efan.AddHandler(SecondGear, FirstGearEvent, FirstGearHandler)
	efan.AddHandler(SecondGear, SecondGearEvent, SecondGearHandler)
	efan.AddHandler(SecondGear, ThirdGearEvent, ThirdGearHandler)
	// 3档状态
	efan.AddHandler(ThirdGear, PowerOffEvent, PowerOffHandler)
	efan.AddHandler(ThirdGear, FirstGearEvent, FirstGearHandler)
	efan.AddHandler(ThirdGear, SecondGearEvent, SecondGearHandler)
	efan.AddHandler(ThirdGear, ThirdGearEvent, ThirdGearHandler)

	// 开始测试状态变化
	efan.Call(ThirdGearEvent)  // 按下3档按钮
	efan.Call(FirstGearEvent)  // 按下1档按钮
	efan.Call(PowerOffEvent)   // 按下关闭按钮
	efan.Call(SecondGearEvent) // 按下2档按钮
	efan.Call(PowerOffEvent)   // 按下关闭按钮
}
