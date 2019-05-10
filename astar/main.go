package main

import (
	"fmt"

	astar "dc-sz/dc/zzq/astar/index"
	model "dc-sz/dc/zzq/model/attr"
)

func main() {
	legal := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
	}
	start := model.Point{0, 0}
	end := model.Point{7, 3}
	lenth, path := astar.Search(legal, start, end)

	fmt.Println("最优步数:", lenth)

	for p := path.Front(); p != nil; p = p.Next() {
		fmt.Println(p.Value)
	}
}
