package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func Prompt(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	r := scanner.Text()
	return r
}

func PromptMulti(prompt string) (items []string) {
	fmt.Println(prompt)
	fmt.Println("(Enter a blank line to continue)")
	for {
		fmt.Print("> ")
		scanner.Scan()
		inputted := scanner.Text()

		if inputted == "" {
			return
		}

		items = append(items, inputted)
	}

}

func PromptSelect(prompt string, options []string, allowNone bool) (selected int, chosenItem string) {
	fmt.Println(prompt)
	for i, v := range options {
		fmt.Printf("  %d: %s\n", i+1, v)
	}
	for {
		fmt.Print("> ")
		scanner.Scan()
		userInput := scanner.Text()

		if userInput == "" && allowNone {
			return -1, ""
		}

		userSelected, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Not an integer.")
			continue
		}

		userSelected -= 1

		if userSelected < 0 || userSelected >= len(options) {
			fmt.Println("Out of bounds.")
			continue
		}

		selected = userSelected
		chosenItem = options[userSelected]

		return
	}
}
