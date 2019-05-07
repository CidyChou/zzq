package main

import (
	"fmt"

	"dc-sz/dc/zzq/battle"
	model "dc-sz/dc/zzq/model/attr"
)

func main() {
	heroA := &model.Hero{1, 700, 0, 1, 100, 5, 1, 1, [2]int{1, 1}}
	heroB := &model.Hero{2, 500, 0, 1, 80, 5, 1, 1, [2]int{4, 4}}

	playerA := []model.Hero{*heroA}
	playerB := []model.Hero{*heroB}

	fmt.Printf("开始,玩家A %v，玩家B %v \n", playerA, playerB)
	battle.Battle(*heroA, *heroB) // 暂时只考虑一个棋子
}
