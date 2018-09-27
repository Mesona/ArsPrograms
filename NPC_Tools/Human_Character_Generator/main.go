package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	intelligence, perception, strength, stamina   int
	presence, communication, dexterity, quickness int
)

func diceRoll(diceType int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(diceType) + 1
}

func menu() {
	fmt.Print("WELCOME MESSAGE\n\n")
	fmt.Print("1) Manual Generation\n")
	fmt.Print("2) Automatic Generation\n")
	fmt.Print("3) Adjust settings\n")
	fmt.Print("4) Exit\n")
	proceed := ""
	fmt.Scanln(&proceed)
	switch {
	case proceed == "1":
		fmt.Print("\n\nOption 1")
	case proceed == "2":
		fmt.Print("\n\nWhich type of NPC would you like to create?\n")
		fmt.Print("1) Combat grog\n")
		fmt.Print("2) Social grog\n")
		fmt.Print("3) Generic townsfolk\n")

		//For when I have more time, heroic characters are weird
		//fmt.Print("4) Heroic NPC\n")

		//Decided these will be their own, separate programs
		//fmt.Print("5) Faerie NPC\n")
		//fmt.Print("6) Faerie monster\n")
		//fmt.Print("7) Magical monser\n")
		//fmt.Print("8) Demonic monster\n")
		//fmt.Print("9) Angelic monster\n")
	case proceed == "3":
		fmt.Print("\n\nOption 3\n")
		fmt.Print("Change region\n")
		fmt.Print("Enable or disable content from beyond the Core rulebook\n")
		fmt.Print("More stuff, maybe\n")
	default:
		fmt.Print("Goodbye")
		os.Exit(1)
	}
}

func main() {
	menu()

}
