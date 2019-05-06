package main

import (
	model "dc-sz/dc/zzq/model/attr"
	"fmt"
	"testing"
)

func Test_Main_正常攻击(t *testing.T) {
	heroA := &model.Hero{1, 700, 0, 1, 100, 5, 1, 1, [2]int{4, 1}}
	heroB := &model.Hero{2, 500, 0, 1, 80, 5, 1, 1, [2]int{4, 4}}

	playerA := []model.Hero{*heroA}
	playerB := []model.Hero{*heroB}

	fmt.Println("开始", playerA, playerB)
	call(*heroA, *heroB)
}
