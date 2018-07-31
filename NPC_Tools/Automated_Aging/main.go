package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var year, yearInt, yearBorn, currentAge, personalModifiers int
var appearantAge, intAge, perAge, strAge, staAge, preAge, comAge, dexAge, quiAge, decrepitude int
var intelligence, perception, strength, stamina, presence, communication, dexterity, quickness int
var globalBonus, name string
var advanceYear, view, older = "n", "n", "n"

func ageCalculation() {
	rollMult := 1
	ageRoll := diceRoll(10)
	for ageRoll == 1 {
		rollMult++
		ageRoll = diceRoll(10)
	}
	bonusInt, _ := strconv.Atoi(globalBonus)
	ageTotal := ageRoll*rollMult + bonusInt + personalModifiers

	//fmt.Print("Die roll: ", ageRoll, "\n")
	//fmt.Print("Age Multiplier: ", rollMult, "\n")
	fmt.Print("Age roll for ", name, ": ", ageTotal, "\n")
	//fmt.Print("Global Bonuses: ", globalBonus, "\n")
	//fmt.Print("Personal Bonuses: ", personalBonus, "\n")
	//return "yeah, I guess"
	switch {
	case ageTotal <= 2:
		fmt.Print(name, " has experienced no adverse effects of age.\n")
	case ageTotal < 10:
		fmt.Print(name, " looks a little older.\n")
		older = "y"
		appearantAge++
	case ageTotal < 13:
		older = "y"
		appearantAge++
		decrepitude++
		// add 1 to any aging stat
	case ageTotal == 13:
		older = "y"
		appearantAge++
		// gain sufficient aging points in any 1 characteristic to reach
		// the next level of decrepitude and crisis
	case ageTotal == 14:
		older = "y"
		appearantAge++
		quiAge++
		decrepitude++
	case ageTotal == 15:
		older = "y"
		appearantAge++
		staAge++
		decrepitude++
	case ageTotal == 16:
		older = "y"
		appearantAge++
		perAge++
		decrepitude++
	case ageTotal == 17:
		older = "y"
		appearantAge++
		preAge++
		decrepitude++
	case ageTotal == 18:
		older = "y"
		appearantAge++
		staAge++
		strAge++
		decrepitude++
		decrepitude++
	case ageTotal == 19:
		older = "y"
		appearantAge++
		dexAge++
		quiAge++
		decrepitude++
		decrepitude++
	case ageTotal == 20:
		older = "y"
		appearantAge++
		comAge++
		preAge++
		decrepitude++
		decrepitude++
	case ageTotal == 21:
		older = "y"
		appearantAge++
		intAge++
		perAge++
		decrepitude++
		decrepitude++
	case ageTotal > 21:
		older = "y"
		appearantAge++
		// gain sufficient aging points in any 1 characteristic to reach
		// the next level of decrepitude and crisis

	default:
		fmt.Print("Bad news for ", name, "\n")
		older = "y"
	}
}

func menu(grabYear int) {
	fmt.Print("The current year is ", grabYear, ". What would you like to do? \n")
	fmt.Print("1) View NPC Ages and statistics \n")
	fmt.Print("2) Advance to year ", grabYear+1, "\n")
	fmt.Print("3) Exit\n")
	proceed := ""
	fmt.Scanln(&proceed)
	switch {
	case proceed == "1":
		fmt.Print("Loading Aging Sheet: \n\n")
		view = "y"
	case proceed == "2":
		fmt.Print("Advancing Character Ages: \n")
		advanceYear = "y"
	default:
		fmt.Print("Goodbye. \n")
		os.Exit(1)
	}

}

func diceRoll(diceType int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(diceType) + 1
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//file, err := os.Open("ages.txt")
	file, err := os.Open("Keras Nisi Aging Chart Test - Sheet1.csv")

	//input, err := ioutil.ReadFile(file)
	input, err := ioutil.ReadFile("Keras Nisi Aging Chart Test - Sheet1.csv")

	// close file on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	// open output file
	fo, err := os.Create("output.csv")
	if err != nil {
		panic(err)
	}

	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	grabYear, _ := strconv.Atoi(string(input[5:9]))
	menu(grabYear)

	// For viewing the aging sheet only
	if view == "y" {
		str := string(input)
		fmt.Println(str)
	}

	// extract each individual line
	lines := strings.Split(string(input), "\n")
	//Name,Year Born,Age,Personal Modifiers,Intelligence Aging,Perception Aging,Strength Aging,Stamina Aging,Prescence Aging,Communication Aging,Dexterity Aging,Quickness Aging,Decrepitude Score
	for _, line := range lines {
		// Displays aging sheet for viewing
		words := strings.Split(string(line), ",")
		//fmt.Print("!!! TEST !!!\n", words[1:], "\n!!! TEST !!!\n")

		arrayInt := make([]int, len(words[1:]))
		for i := range arrayInt {
			arrayInt[i], _ = strconv.Atoi(words[i+1])
		}
		fmt.Println("!!! TEST ARRAY !!!\n", arrayInt)

		if words[0] == "Living Modifiers" {
			words[0] = "\nLiving Modifiers"
			globalBonus = words[1]
			output := strings.Join(words, ",")
			if _, err = fo.WriteString(output); err != nil {
				panic(err)
			}
			if _, err = fo.WriteString("\n"); err != nil {
				panic(err)
			}
		} else if words[0] == "Year" {
			yearString := words[1]
			size := len(yearString)
			year, _ = strconv.Atoi(yearString[:size-1])
			if advanceYear == "y" {
				year++
				words[1] = strconv.Itoa(year)
			}
			output := strings.Join(words, ",")
			if _, err = fo.WriteString(output); err != nil {
				panic(err)
			}
		} else if words[0] == "Name" {
			continue
			//} else if words[12] != "" {
		} else if len(line) > 25 {
			// This line is meant to detect lines that are characters
			name = words[0]
			//yearBorn, _ = strconv.Atoi(words[1])
			yearBorn = arrayInt[0]
			currentAge = year - yearBorn
			//appearantAge = words[3]
			appearantAge = arrayInt[2]
			//personalModifiers, _ = strconv.Atoi(words[4])
			personalModifiers = arrayInt[3]
			//intAge = words[5]
			intAge = arrayInt[4]
			//perAge = words[6]
			perAge = arrayInt[5]
			//strAge = words[7]
			strAge = arrayInt[6]
			//staAge = words[8]
			staAge = arrayInt[7]
			//preAge = words[9]
			preAge = arrayInt[8]
			//comAge = words[10]
			comAge = arrayInt[9]
			//dexAge = words[11]
			dexAge = arrayInt[10]
			//quiAge = words[12]
			quiAge = arrayInt[11]
			//decrepitude = words[13]
			decrepitude = arrayInt[12]
			//fmt.Print("Testing: \n")
			//fmt.Print("Int Aging: ", intAge, "\n")
			//fmt.Print("perAge: ", perAge, "\n")
			//fmt.Print("Str Age:", strAge, "\n")
			//fmt.Print("staAge:", staAge, "\n")
			//fmt.Print("preAge: ", preAge, "\n")
			//fmt.Print("com age:", comAge, "\n")
			//fmt.Print("dexAge: ", dexAge, "\n")
			//fmt.Print("quiAge: ", quiAge, "\n")
			//fmt.Print("decrepitude: ", decrepitude, "\n")

			if advanceYear == "y" {
				if currentAge >= 35 {
					fmt.Print("Current Age: ", currentAge, "\n")
					ageCalculation()
					if older == "y" {
						oldAge, _ := strconv.Atoi(words[3])
						oldAge++
						words[3] = strconv.Itoa(oldAge)
					}
				}
			} else {
				fmt.Print(words)
			}

			output := strings.Join(words, ",")
			if _, err = fo.WriteString(output); err != nil {
				panic(err)
			}
			if _, err = fo.WriteString("\n"); err != nil {
				panic(err)
			}
		} else {
			fmt.Print("")
		}
	}
}
