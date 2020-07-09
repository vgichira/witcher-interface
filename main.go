package main

import "fmt"

type (
	character struct {
		name     string
		level    uint64
		attacks  map[string]int
		vitality int
	}

	witcher struct {
		character
		school string
	}

	monster struct {
		character
		monsterType string
	}

	combatCharacters interface {
		attack(attackName string) int
		takeDamage(damage int)
		getName() string
		getVitality() int
	}
)

func doDamage(attacker, victim combatCharacters, attackName string) {
	// we get the damage inflicted by an attacker using the attackName
	damageInflicted := attacker.attack(attackName)
	// we pass the damage inflicted to the victim character
	victim.takeDamage(damageInflicted)

	fmt.Println(attacker.getName(), "attacked", victim.getName(), "did", damageInflicted, "damage")

	fmt.Println(victim.getName(), "now has", victim.getVitality(), "health remaining")
}

// WITCHER

func (witch witcher) attack(attackName string) int {
	return witch.attacks[attackName]
}

func (witch *witcher) takeDamage(damage int) {
	witch.vitality = witch.vitality - damage
}

func (witch witcher) getName() string {
	return witch.name
}

func (witch witcher) getVitality() int {
	return witch.vitality
}

// MONSTER

func (monst monster) attack(attackName string) int {
	return monst.attacks[attackName]
}

func (monst *monster) takeDamage(damage int) {
	monst.vitality = monst.vitality - damage
}

func (monst monster) getName() string {
	return monst.name
}

func (monst monster) getVitality() int {
	return monst.vitality
}

func main() {
	geralt := witcher{
		character: character{
			name:  "Geralt",
			level: 112,
			attacks: map[string]int{
				"swordHeavy": 1250,
				"swordQuick": 750,
			},
			vitality: 3000,
		},
		school: "Wolf",
	}

	ekkimmara := monster{
		character: character{
			name:  "Ekkimmara",
			level: 114,
			attacks: map[string]int{
				"heavySlash": 1440,
				"rush":       800,
			},
			vitality: 6500,
		},
	}

	doDamage(&ekkimmara, &geralt, "heavySlash")
}
