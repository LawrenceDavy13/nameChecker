package main

import (
	"flag" // Process command-line flags
	"fmt"
	"log"
	"regexp" // Vaildate regular expression matches
)

const UsernameRegex string = `^@?(\w){1,15}$`

func main() {

	// Pass in username
	var userInput string // Contains string of user input from the command line
	flag.StringVar(&userInput, "username", "Aiwass", "The Great Checker")
	flag.Parse()

	// Print result
	fmt.Println("The Great Checker is here")
	fmt.Println("Checking username..., \"", userInput, "\", resulted in: ", CheckUser(userInput))
}



// Check username
func CheckUser(username string) bool {
	validation := false
	r, err := regexp.Compile(UsernameRegex)
	if err != nil {
		log.Fatalln(err)
	}

	validation = r.MatchString(username)
	return validation
}

