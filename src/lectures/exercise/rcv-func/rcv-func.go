//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	health int
	maxHealth int
	energy int
	maxEnergy int
	name string
}

func (player *Player) updateHealth(health int) {
	player.health = health
	fmt.Println("Health:", player.health)
}

func (player *Player) updateEnergy(energy int) {
	player.energy = energy
	fmt.Println("Energy:", player.energy)
}

func main() {
	player := Player{100, 100, 100, 100, "Duane"}
	fmt.Println("Player:", player)

	player.updateHealth(80)
	player.updateEnergy(70)
}
