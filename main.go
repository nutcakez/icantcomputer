package main

type Battle struct {
	UnitManager
	// other components
}

type UnitManager struct {
	units []Unit
}

func (um *UnitManager) Update() {
	for _, v := range um.units {
		// being sloppy here, we got the target ok?
		v.Attack(um.units[0])
	}
}

type Unit interface {
	Attack(Unit)
	TakeDmg(int)
}

func main() {
	battle := Battle{
		UnitManager: UnitManager{
			units: []Unit{
				&dude{},
				&dude{},
				&dude{},
			},
		},
	}

	battle.Update()
}
