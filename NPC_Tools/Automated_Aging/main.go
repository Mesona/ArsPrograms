package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var appearantAge, year, yearInt, yearBorn, effectiveAge int
var globalBonus, decrepitude, personalModifiers int
var name, virtues string
var sfb, ua, aq, crisis = "n", "n", "n", ""
var advanceYear, view, older = "n", "n", "n"
var filename = "Aging Chart Template.csv"

func ageCalculation() {
	rollMult := 1
	ageRoll := diceRoll(10)
	for ageRoll == 1 {
		rollMult++
		ageRoll = diceRoll(10)
	}
	agePenalty := effectiveAge / 10
	ageTotal := ageRoll*rollMult + globalBonus + personalModifiers + agePenalty

	//fmt.Print("Die roll: ", ageRoll, "\n")
	//fmt.Print("Age Multiplier: ", rollMult, "\n")
	fmt.Print("Age roll for ", name, ": ", ageTotal, "\n")
	//fmt.Print("Global Bonuses: ", globalBonus, "\n")
	//fmt.Print("Personal Bonuses: ", personalBonus, "\n")
	//return "yeah, I guess"
	switch {
	case ageTotal <= 2:
		fmt.Print(name, " has experienced no adverse effects of age.\n")
	case ageTotal <= 9:
		fmt.Print(name, " looks a little older.\n")
		older = "y"
		appearantAge++
	case ageTotal <= 12:
		older = "y"
		fmt.Print(name, " looks a little older.\n")
		appearantAge++
		decrepitude++
	case ageTotal == 13:
		older = "y"
		fmt.Print(name, " has experienced a crisis.\n")
		appearantAge++
		switch {
		case decrepitude < 5:
			decrepitude = 5
		case decrepitude < 15:
			decrepitude = 15
		case decrepitude < 30:
			decrepitude = 30
		case decrepitude < 50:
			decrepitude = 50
		case decrepitude < 75:
			decrepitude = 75
		}
		crisisCalculation()
	case ageTotal <= 17:
		older = "y"
		fmt.Print(name, " looks a little older.\n")
		appearantAge++
		decrepitude++
	case ageTotal <= 21:
		older = "y"
		fmt.Print(name, " looks a little older.\n")
		appearantAge++
		decrepitude++
		decrepitude++
	case ageTotal > 21:
		older = "y"
		fmt.Print(name, " has experienced a crisis.\n")
		appearantAge++
		switch {
		case decrepitude < 5:
			decrepitude = 5
		case decrepitude < 15:
			decrepitude = 15
		case decrepitude < 30:
			decrepitude = 30
		case decrepitude < 50:
			decrepitude = 50
		case decrepitude < 75:
			decrepitude = 75
		}
		crisisCalculation()
	}
}

func crisisCalculation() {
	decrepitudeScore := math.Floor((math.Sqrt(8*(float64(decrepitude)/5)+1) - 1) / 2)
	fmt.Print("\nDECREPTIUDE SCORE: ", decrepitudeScore, "\n")
	crisisRoll := diceRoll(10)
	agePenalty := effectiveAge / 10
	crisisTotal := crisisRoll + agePenalty + int(decrepitudeScore)
	switch {
	case crisisTotal <= 14:
		if ua == "y" {
			crisis = "Crisis prevented due to Unaging"
		} else {
			crisis = "Minor sickness, a season of work is lost"
		}
	case crisisTotal == 15:
		crisis = "Minor illness. Stamina roll of 3+ or CrCo20 to survive"
	case crisisTotal == 16:
		crisis = "Serious illness. Stamina roll of 6+ or CrCo25 to survive"
	case crisisTotal == 17:
		crisis = "Major illness. Stamina roll of 9+ or CrCo30 to survive"
	case crisisTotal == 18:
		crisis = "Critical illness. Stamina roll of 12+ or CrCo35 to survive"
	case crisisTotal >= 19:
		crisis = "Terminal illness. CrCo40 required to survive"

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
	//file, err := os.Open("Keras Nisi Aging Chart Test - Sheet1.csv")
	file, err := os.Open(filename)

	//input, err := ioutil.ReadFile(file)
	//input, err := ioutil.ReadFile("Keras Nisi Aging Chart Test - Sheet1.csv")
	input, err := ioutil.ReadFile(filename)

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
	for _, line := range lines {
		// Displays aging sheet for viewing
		words := strings.Split(string(line), ",")
		sfb, ua, aq, crisis = "n", "n", "n", ""
		older = "n"

		var arrayInt []int
		if len(words) > 1 {
			arrayInt = make([]int, len(words[2:]))
			for i := range arrayInt {
				arrayInt[i], _ = strconv.Atoi(words[i+1])
			}
		}

		if words[0] == "Living Modifiers" {
			words[0] = "\nLiving Modifiers"
			globalBonus, _ = strconv.Atoi(words[1])
			output := strings.Join(words, ",")
			if _, err = fo.WriteString(output); err != nil {
				panic(err)
			}
			if _, err = fo.WriteString("\n"); err != nil {
				panic(err)
			}
			continue
		} else if words[0] == "Year" {
			if advanceYear == "y" {
				grabYear++
				words[1] = strconv.Itoa(grabYear)
			}
			output := strings.Join(words, ",")
			if _, err = fo.WriteString(output); err != nil {
				panic(err)
			}
			continue
		} else if words[0] == "Name" {
			output := strings.Join(words, ",")
			if _, err = fo.WriteString(output); err != nil {
				panic(err)
			}
			if _, err = fo.WriteString("\n"); err != nil {
				panic(err)
			}
			continue
		} else if line != "" {
			name = words[0]
			virtues = words[1]
			if strings.Contains(virtues, "FB") {
				if strings.Contains(virtues, "SFB") {
					personalModifiers = -3
					sfb = "y"
				} else {
					personalModifiers = -1
				}
			}
			if strings.Contains(virtues, "UA") {
				ua = "y"
			}
			if strings.Contains(virtues, "AQ") {
				aq = "y"
			}
			yearBorn = arrayInt[1]
			effectiveAge = arrayInt[2]
			fmt.Print("\nEffective Age of ", name, " is: ", effectiveAge, "\n\n")
			words[3] = strconv.Itoa(effectiveAge)
			appearantAge = arrayInt[3]
			decrepitude = arrayInt[4]

			if advanceYear == "y" {
				effectiveAge++
				words[3] = strconv.Itoa(effectiveAge)
				if effectiveAge >= 35 && sfb == "n" || effectiveAge >= 50 {
					fmt.Print("Current Age: ", effectiveAge, "\n")
					ageCalculation()
					if aq == "y" {
						ageCalculation()
						fmt.Print(name, " appears to be aging quicker \n")
						effectiveAge++
					}
					if older == "y" {
						oldAge, _ := strconv.Atoi(words[4])
						oldAge++
						if aq == "y" {
							oldAge++
						}
						words[3] = strconv.Itoa(effectiveAge)
						words[4] = strconv.Itoa(oldAge)
						words[5] = strconv.Itoa(decrepitude)
						words[6] = crisis
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
	rename := os.Rename("output.csv", filename)
	if rename != nil {
		panic(rename)
	}
}
