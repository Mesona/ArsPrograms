package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/robertkrimen/otto"
)

func jsDownload(nameType string) {
	url := "http://www.fantasynamegenerators.com/scripts/" + nameType + ".js"
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		//fmt.Printf("%s\n", string(contents))

		// open output file
		fo, err := os.Create("nameGenerator.js")
		if err != nil {
			panic(err)
		}

		// close fo on exit and check for its returned error
		defer func() {
			if err := fo.Close(); err != nil {
				panic(err)
			}
		}()

		// writes the downloaded .js file to nameGenerator.js
		if _, err = fo.WriteString(string(contents)); err != nil {
			panic(err)
		}
		runNameGenerator(string(contents))
	}
}

func runNameGenerator(script string) {
	//var vm otto.Otto
	vm := otto.New()
	vm.Run(script)
}

func main() {
	fmt.Print("What type of name would you like generated?\n")
	fmt.Print("1) Ancient Greek\n")
	fmt.Print("2) Irish\n")
	fmt.Print("3) Other\n")
	fmt.Print("4) Exit\n")
	proceed := ""
	fmt.Scanln(&proceed)
	switch {
	case proceed == "1":
		jsDownload("ancientGreekNames")
	case proceed == "2":
		jsDownload("irishNames")
	case proceed == "3":
		jsDownload("blahblhblah")
	default:
		fmt.Print("Goodbye\n")
		os.Exit(1)
	}

	//runNameGenerator()
}
