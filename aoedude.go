package main

// doing aoe damage around the target
type aoeDude struct {
	dmg, hp, attackRange, area int
	// putting something into the struct to figure out a way to get units around the target
	// 1 unitsInRange(int, int) []Unit func
	// 2 just embed the whole unitManager here yolo
}

func (d *aoeDude) TakeDmg(dmgTaken int) {
	d.hp -= dmgTaken
}

// other
// 3 pass ALLLLLLLL the targets, unit itself will figure out what to attack/who takes dmg
// 4 return with a value and unit manager will find the aoe targets, dont like it cuz Unit related logic
// bleeds into unitmanager
func (d *aoeDude) Attack(target Unit) {
	target.TakeDmg(d.dmg)
	// i want to do dmg to the rest T_T
}

func (d *aoeDude) GetRange() int {
	return d.attackRange
}
