package v1

import "math"

func (hero *Hero) Move(enemy Hero) error {
	return nil
}

//寻路
func FindPath() {

}

//生成路径
func generatePath() {

}

//生成距离
func getDistance(src [2]int, dest [2]int) float64 {
	a := math.Abs(float64(src[0] - dest[0]))
	b := math.Abs(float64(src[1] - dest[1]))
	return math.Sqrt(float64(a*a + b*b))
}

//攻击
// func (hero *Hero) Attack(enemy Hero) error {

// 	return nil
// }


