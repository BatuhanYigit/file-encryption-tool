package helper

import "fmt"

func WelcomeMessage() {
	fmt.Println("/**************************************************************/")
	fmt.Println("**********************File Encryption Tool**********************")
	fmt.Println("Use the command")
	fmt.Println("help\n encrypt\n decrypt")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("/**************************************************************/")
}

func HelpMessage() {
	fmt.Println("File encryption")
	fmt.Println("Simple file ecrypter for your day-today needs.")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("Select encrypt - path your file")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("Encrypts a file given password")
	fmt.Println("Select decrypt Tries to decrypt a file using a password")
	fmt.Println("help displays select help")
}
