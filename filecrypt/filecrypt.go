package filecrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

func Encrypt(source string, password []byte) {

	if _, err := os.Stat(source); os.IsNotExist(err) {
		fmt.Println(err.Error())
	}

	plaintext, err := ioutil.ReadFile(source)

	if err != nil {
		fmt.Println(err.Error())
	}

	key := password
	nonce := make([]byte, 12)

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err.Error())
	}

	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		fmt.Println(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	ciphertext = append(ciphertext, nonce...)

	f, err := os.Create(source)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = io.Copy(f, bytes.NewReader(ciphertext))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Decrypt(source string, password []byte) {

	if _, err := os.Stat(source); os.IsNotExist(err) {
		fmt.Println(err.Error())
	}

	ciphertext, err := ioutil.ReadFile(source)

	if err != nil {
		fmt.Println(err.Error())
	}

	key := password
	salt := ciphertext[len(ciphertext)-12:]
	str := hex.EncodeToString(salt)

	nonce, err := hex.DecodeString(str)

	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		fmt.Println(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext[:len(ciphertext)-12], nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	f, err := os.Create(source)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = io.Copy(f, bytes.NewReader(plaintext))
	if err != nil {
		fmt.Println(err.Error())
	}
}
