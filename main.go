package main

import (
	"fmt"
	"math"
	"sort"

	model "dc-sz/dc/zzq/model/attr"
)

func main() {
	heroA := &model.Hero{1, 700, 0, 1, 100, 5, 1, 1, [2]int{1, 1}}
	heroB := &model.Hero{2, 500, 0, 1, 80, 5, 1, 1, [2]int{4, 4}}

	playerA := []model.Hero{*heroA}
	playerB := []model.Hero{*heroB}

	fmt.Println("开始", playerA, playerB)
	call(*heroA, *heroB)
}

// 主流程
//func call(playerA []model.Hero, playerB []model.Hero) {
func call(hero model.Hero, anemyHero model.Hero) {
	//fmt.Println("棋子间的距离   :", getDistanceByHero(hero, anemyHero))

	if hero.HP <= 0 {
		fmt.Printf("英雄%v  获胜 \n", hero.Id)
		return
	}

	if anemyHero.HP <= 0 {
		fmt.Printf("英雄%v  获胜 \n", anemyHero.Id)
		return
	}

	//TODO 检索目标
	isHeroAttack := false
	isAnemyHeroAttack := false

	for _, attackRange := range getAttackRange(hero.Position, hero.AttackRange) {
		if attackRange == anemyHero.Position {
			//攻击
			anemyHero = attack(hero, anemyHero)
			isHeroAttack = true
			break
		}
	}

	if !isHeroAttack {
		hero = move(hero, SearchPath(hero, anemyHero))
	}

	for _, attackRange := range getAttackRange(anemyHero.Position, anemyHero.AttackRange) {
		if attackRange == hero.Position {
			//攻击
			hero = attack(anemyHero, hero)
			isAnemyHeroAttack = true
			break
		}
	}

	if !isAnemyHeroAttack {
		anemyHero = move(anemyHero, SearchPath(anemyHero, hero))
	}

	call(hero, anemyHero)
}

// 循环
func loop(hero model.Hero, anemyHero model.Hero) {
	loop(hero, anemyHero)
}

// 攻击
func attack(hero model.Hero, anemyHero model.Hero) model.Hero {
	fmt.Printf("英雄%v  开始攻击,原始HP:%v \n", hero.Id, anemyHero.HP)
	damage := hero.Attack - anemyHero.Armor
	anemyHero.HP = anemyHero.HP - damage
	fmt.Printf("英雄%v  开始攻击,被攻击后HP：%v \n", hero.Id, anemyHero.HP)
	return anemyHero
}

// 移动
func move(hero model.Hero, position [2]int) model.Hero {
	fmt.Printf("英雄%v  开始移动：初始位置:%v 移动后位置:%v \n", hero.Id, hero.Position, position)
	hero.Position = position
	return hero
}

//英雄可攻击范围
func getAttackRange(p [2]int, attackRange int) [][2]int {
	attackRanges := [][2]int{
		{p[0] + attackRange, p[1]},
		{p[0] - attackRange, p[1]},
		{p[0], p[1] + attackRange},
		{p[0], p[1] - attackRange},
	}
	return attackRanges
}

func searchTarget(hero model.Hero, anemyHeros []*model.Hero) model.Hero {
	var h model.Hero
	p := hero.Position
	attackRanges := getAttackRange(p, hero.AttackRange)
	//var anemyHerosMap map[float64]Hero //在有效攻击范围的英雄

	for _, anemyHero := range anemyHeros {
		for _, attackRange := range attackRanges {
			if attackRange == anemyHero.Position {
				// TODO 多个敌方英雄的情况
				//distance := getDistance(attackRange, anemyHero.Position)
				//anemyHerosMap[distance] = *anemyHero
				return *anemyHero
			}
		}
	}

	// if len(anemyHerosMap) > 1 {

	// }
	return h
}

//SearchPath is 寻路
func SearchPath(hero model.Hero, anemyHero model.Hero) [2]int {
	p := hero.Position
	availablepaths := getAttackRange(p, hero.AttackRange)
	// TODO 过滤棋牌以外的位置
	var pathMap map[float64][2]int
	pathMap = make(map[float64][2]int)

	var keys []float64

	for _, availablepath := range availablepaths {
		path := getDistance(availablepath, anemyHero.Position)
		keys = append(keys, path)
		pathMap[path] = availablepath
	}

	sort.Float64s(keys)
	return pathMap[keys[0]]
}

//距离
func getDistanceByHero(hero model.Hero, anemyHero model.Hero) float64 {
	a := math.Abs(float64(hero.Position[0] - anemyHero.Position[0]))
	b := math.Abs(float64(hero.Position[1] - anemyHero.Position[1]))

	return math.Sqrt(float64(a*a + b*b))
}

//生成距离
func getDistance(src [2]int, dest [2]int) float64 {
	a := math.Abs(float64(src[0] - dest[0]))
	b := math.Abs(float64(src[1] - dest[1]))
	return math.Sqrt(float64(a*a + b*b))
}
