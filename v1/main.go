package v1

var (
	// 玩家A 英雄A初始位置
	p_a = [2]int{4, 1}
	// 玩家B 英雄B初始位置
	p_b = [2]int{4, 4}

	//初始化英雄A
	hero_a = &Hero{700, 0, 1, 100, 5, 1, 1, p_a}
	//初始化英雄B
	hero_b = &Hero{500, 0, 1, 80, 5, 1, 1, p_b}

	player_a = []Hero{*hero_a}
	player_b = []Hero{*hero_b}
)

//棋盘
//英雄切片

// func main() {
// 	// 初始化棋盘
// 	chessboard := NewChessboard(hero_a)

// 	// 初始化状态机
// 	//chessboard.

// 	//得出结果JSON

// 	//返回状态(移动?攻击?)

// }

//实例化棋盘
// func NewChessboard(initPlayer_a player_a, initPlayer_b player_b) *Chessboard {

// }

//实例化英雄
func NewHero(initState Hero) *Hero {
	return &Hero{
		HeroFSM: New(initState),
	}
}
