package main

import (
	"bytes"
	"file-encryption-tool/filecrypt"
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
		fmt.Println("File path: ")
		var filePath string
		fmt.Scanln(&filePath)
		encryptHandle(filePath)

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
	filecrypt.Encrypt(filePath, password)
	fmt.Println("File successfully encrypting.")

}

func decryptHandle(filePath string) {
	if len(filePath) == 0 {
		fmt.Println("Missing path file")
	}
	if !validateFile(filePath) {
		fmt.Println("File not found")
	}

	fmt.Println("Enter Password : ")
	password, _ := term.ReadPassword(0)
	fmt.Println("Decrypting File...")
	filecrypt.Decrypt(filePath, password)
	fmt.Println("File successfully decrypting.")
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
