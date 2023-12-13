package main

import (
	"bytes"
	"file-encryption-tool/helper"
	"fmt"
	"os"

	"golang.org/x/term"
)

func main() {
	helper.WelcomeMessage()
	fmt.Println("\nSelect Choice")
	var choice string
	fmt.Scanln(&choice)

	switch choice {
	case "help":
		helper.HelpMessage()

	case "encrypt":

	}

}

func encryptHandle(filePath string) {
	if len(filePath) == 0 {
		fmt.Println("Missing path file")
	}
	if !validateFile(filePath) {
		fmt.Println("File not found")
	}

	password := getPassword()
	fmt.Println("Encrypting File...")

}

func validateFile(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true

}

func validatePassword(password1 []byte, password2 []byte) bool {
	if !bytes.Equal(password1, password2) {
		return false
	}
	return true
}

func getPassword() []byte {
	fmt.Println("Enter password : ")
	password, _ := term.ReadPassword(0)
	fmt.Print("\n Confirm Password : ")
	password2, _ := term.ReadPassword(0)

	if !validatePassword(password, password2) {
		fmt.Println("Passwords do not match please try again")
		return getPassword()
	}

	return password

}
