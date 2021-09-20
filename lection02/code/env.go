//nolint: unused
package main

import (
	"flag"
	"fmt"
	"os"
)

func userInput() {
	fmt.Println("Enter login:")
	var login string
	_, _ = fmt.Scan(&login) // same as fmt.Fscan(os.Stdin, &login)

	fmt.Println("Enter animal and random integer:")
	var (
		animal string
		number int
	)
	_, _ = fmt.Scanf("%s %d", &animal, &number)

	fmt.Printf("user %s likes %s and his number is %d", login, animal, number)
}

var verbose = "VERBOSE"

func env() {
	s, ok := os.LookupEnv(verbose)
	fmt.Println(s, ok)

	s = os.Getenv(verbose)
	fmt.Println(s)

	all := os.Environ()
	fmt.Println(all)
}

func flags() {
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "add expanded logs")

	var configFilePath = flag.String("config", "", "path to config file")
	flag.Parse()

	if verbose {
		fmt.Println("verbose is enabled")
	}

	if configFilePath != nil {
		fmt.Printf("path to file: %s\n", *configFilePath)
	}
}
