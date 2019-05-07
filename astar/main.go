package main

import (
	"fmt"

	astar "dc-sz/dc/zzq/astar/index"
)

func main() {
	legal := [][]int{
		{1, 1, 1, 1},
		{0, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}
	start := astar.Point{0, 0}
	end := astar.Point{3, 3}
	lenth, path := astar.Search(legal, start, end)

	fmt.Println("最优步数:", lenth)

	for p := path.Front(); p != nil; p = p.Next() {
		fmt.Println(p.Value)
	}
}
