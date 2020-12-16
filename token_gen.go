package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	sex := randomSex()
	// true female false male

	// strength := randomAttr()
	// dexterity := randomAttr()
	// constitution := randomAttr()
	// intelligence := randomAttr()
	// wisdom := randomAttr()
	// charisma := charisma()

	classStruct := map[string]string{"name": randomClass()}
	race := map[string]string{"name": randomRace(sex), "benefits": ""}
	attributes := map[string]int{"strength": randomAttr(), "dexterity": randomAttr(), "constitution": randomAttr(), "intelligence": randomAttr(), "wisdom": randomAttr(), "charisma": randomAttr()}

	fmt.Println(classStruct)
	fmt.Println(race)
	fmt.Println(attributes)
	fmt.Println(sex)
}

func random(min, max int) int {

	return min + rand.Intn(max-min)
}

func randomClass() string {
	value := random(1, 15)
	switch value {
	case 1:
		return "Arcanista"
	case 2:
		return "Barbáro"
	case 3:
		return "Bardo"
	case 4:
		return "Bucaneiro"
	case 5:
		return "Caçador"
	case 6:
		return "Cavaleiro"
	case 7:
		return "Clérigo"
	case 8:
		return "Druida"
	case 9:
		return "Guerreiro"
	case 10:
		return "Inventor"
	case 11:
		return "Ladino"
	case 12:
		return "Lutador"
	case 13:
		return "Nobre"
	case 14:
		return "Paladino"
	default:
		return randomClass()
	}
}

func randomRace(sex bool) string {
	value := 0
	if sex {
		value = random(-2, 16)
	} else {
		value = random(1, 17)
	}
	switch value {
	case -2:
		return "Sereia"
	case -1:
		return "Medusa"
	case 0:
		return "Humano"
	case 1:
		return "Anão"
	case 2:
		return "Dahllan"
	case 3:
		return "Elfo"
	case 4:
		return "Goblin"
	case 5:
		return "Lefou"
	case 6:
		return "Minotauro"
	case 7:
		return "Qareen"
	case 8:
		return "Golen"
	case 9:
		return "Hynne"
	case 10:
		return "Kliren"
	case 11:
		return "Osteon"
	case 12:
		return "Sílfide"
	case 13:
		return "Anggelus"
	case 14:
		return "Sulfure"
	case 15:
		return "Trog"
	case 16:
		return "Tritão"
	default:
		return "Humano"
	}
}

func randomSex() bool {
	if random(1, 3) == 1 {
		return true
	}
	return false
}

func randomAttr() int {

	rolls := []int{random(1, 7), random(1, 7), random(1, 7), random(1, 7)}
	sort.Ints(rolls)
	total := attrSumOfRoll(rolls)
	if total < 8 {
		total = random(8, 20)
	}
	return total

}
func attrSumOfRoll(slice []int) int {
	result := 0
	for i := 1; i <= 3; i++ {
		result += slice[i]
	}
	return result

}
