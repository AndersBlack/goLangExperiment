package main

import "fmt"

func listOfFun() {

	commands := [10]string{"test1", "test2", "3", "4", "5", "6", "7", "8", "9", "10"}

	fmt.Println(commands[3])

}

func main() {

	fmt.Println("Welcome to the test project for go by Anders Black. In order to partake in the fun you might wanna check out /help")
	var helpInput string
	fmt.Scanln(&helpInput)

	if helpInput == "/help" {
		listOfFun()
	} else {
		fmt.Println("Not strong with following simple directions, are we? It's cool, neither was i in kindergarden. Try again.")
		var retryHelp string
		fmt.Scanln(&retryHelp)
		if retryHelp == "/help" {
			listOfFun()
		} else {
			fmt.Print("I think you are out of pedagogical reach. In case of cat on keyboard, you can always run the program again.")
		}
	}
}
