package main

import (
	"dc-sz/dc/zzq/battle"
	model "dc-sz/dc/zzq/model/attr"
)

func main() {
	heroA := model.Hero{1, 1, 700, 100, 5, 1, 1, [2]int{1, 1}, model.Point{1, 1}}
	heroB := model.Hero{2, 2, 500, 80, 5, 1, 1, [2]int{4, 4}, model.Point{4, 4}}

	var heroes []model.Hero
	heroes = append(heroes, heroA, heroB)

	battle.Battle(heroes)
}
