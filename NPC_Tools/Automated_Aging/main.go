package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var globalBonus int

//name, age, decInt, decPer, decStr, decSta, decPre, decCom, decDex, decQui
func ageCalculation(age, bonus string) string {
	personalBonus, _ := strconv.Atoi(bonus)

	rollMult := 1
	ageRoll := diceRoll(10)
	for ageRoll == 1 {
		rollMult++
		ageRoll = diceRoll(10)
	}
	ageTotal := ageRoll*rollMult + globalBonus + personalBonus

	//fmt.Print("Die roll: ", ageRoll, "\n")
	//fmt.Print("Age Multiplier: ", rollMult, "\n")
	fmt.Print("Age roll: ", ageTotal, "\n")
	//fmt.Print("Global Bonuses: ", globalBonus, "\n")
	//fmt.Print("Personal Bonuses: ", personalBonus, "\n")
	return "yeah, I guess"
}

func menu(year int) {
	fmt.Print("The current year is ", year, ". What would you like to do? \n")
	fmt.Print("1) : View NPC Ages and statistics \n")
	fmt.Print("2) Advance to year ", year+1, "\n")
	fmt.Print("3) Exit\n")
	proceed := ""
	fmt.Scanln(&proceed)
	switch {
	case proceed == "1":
		fmt.Print("You selected 1\n")
	case proceed == "2":
		fmt.Print("You selected 2\n")
	default:
		fmt.Print("Goodbye. \n")
	}

}

func diceRoll(diceType int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(diceType) + 1
}

func main() {
	//file, err := os.Open("ages.txt")
	file, err := os.Open("Keras Nisi Aging Chart Test - Sheet1.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strSplit := strings.Split(scanner.Text(), ",")
		if strSplit[0] == "Living Modifiers" {
			globalBonus, _ = strconv.Atoi(strSplit[1])
		} else if strSplit[0] == "Year" {
			year, _ := strconv.Atoi(strSplit[1])
			strSplit[1] = strconv.Itoa(year + 1)
			menu(year)
		} else {
			ageCalculation(strSplit[2], strSplit[3])
			//fmt.Print(strSplit)
			//fmt.Print(strSplit[2])
			//fmt.Print(strSplit[3])

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
