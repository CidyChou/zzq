package model

//Hero is 英雄属性
type Hero struct {
	Id          int
	HP          int    // 血量
	MP          int    // 魔法值
	Level       int    // 等级
	Attack      int    // 攻击力
	Armor       int    // 护甲
	Pace        int    // 移动速度
	AttackRange int    // 攻击范围
	Position    [2]int // 位置
	//*HeroFSM
}

//IHero is 英雄接口
type IHero interface {
	Move(enemy Hero) error           //　移动
	Attack(enemy Hero) (Hero, error) // 攻击
	Skill(enemy Hero) (Hero, error)  //技能释放
}
