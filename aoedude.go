package main

// doing aoe damage around the target
type aoeDude struct {
	dmg, hp, attackRange, area int
}

func (d *aoeDude) TakeDmg(dmgTaken int) {
	d.hp -= dmgTaken
}

func (d *aoeDude) Attack(target Unit) {
	target.TakeDmg(d.dmg)
	// i want to do dmg to the rest T_T
}

func (d *aoeDude) GetRange() int {
	return d.attackRange
}
