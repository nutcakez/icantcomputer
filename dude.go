package main

type dude struct {
	dmg, hp, attackRange int
	posX, posY           int
}

func (d *dude) TakeDmg(dmgTaken int) {
	d.hp -= dmgTaken
}

func (d *dude) Attack(target Unit) {
	target.TakeDmg(d.dmg)
}

func (d *dude) GetRange() int {
	return d.attackRange
}
