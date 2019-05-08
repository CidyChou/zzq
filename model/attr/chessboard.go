package model

type chessboard struct {
	Heromap map[int]Hero
	legal   [][]int
}
