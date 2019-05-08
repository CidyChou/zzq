package model

//Hero is 英雄属性
type Hero struct {
	ID          int    // Id
	Type        int    // 英雄类型 1:我方 2：敌方 3：其他
	HP          int    // 血量
	Attack      int    // 攻击力
	Armor       int    // 护甲
	Pace        int    // 移动速度
	AttackRange int    // 攻击范围
	Position    [2]int // 位置
	Point
}

//IHero is 英雄接口
type IHero interface {
	Move(enemy Hero) error           //　移动
	Attack(enemy Hero) (Hero, error) // 攻击
	Skill(enemy Hero) (Hero, error)  //技能释放
}
