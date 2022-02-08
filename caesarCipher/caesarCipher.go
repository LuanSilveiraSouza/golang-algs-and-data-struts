package caesarCipher

import (
	s "strings"
	"unicode"
)

const alphabet string = "abcdefghijklmnopqrstuvwxyz"

func Encrypt(message string, shift int) string {
	var encryptedMessage = ""
	for _, value := range message {
		var encryptedValue = string(value)
		if unicode.IsLetter(value) {
			var encryptedIndex = s.Index(alphabet, s.ToLower(string(value))) + shift
			if encryptedIndex > len(alphabet) {
				encryptedIndex = encryptedIndex - len(alphabet)
			}
			encryptedValue = string(alphabet[encryptedIndex])
		}
		if string(value) == s.ToUpper(string(value)) {
			encryptedValue = s.ToUpper(encryptedValue)
		}
		encryptedMessage = s.Join([]string{encryptedMessage, encryptedValue}, "")
	}
	return encryptedMessage
}

func Decrypt(message string, shift int) string {
	var decryptedMessage = ""
	for _, value := range message {
		var encryptedValue = string(value)
		if unicode.IsLetter(value) {
			var encryptedIndex = s.Index(alphabet, s.ToLower(string(value))) - shift
			if encryptedIndex < 0 {
				encryptedIndex = len(alphabet) + encryptedIndex
			}
			encryptedValue = string(alphabet[encryptedIndex])
		}
		if string(value) == s.ToUpper(string(value)) {
			encryptedValue = s.ToUpper(encryptedValue)
		}
		decryptedMessage = s.Join([]string{decryptedMessage, encryptedValue}, "")
	}
	return decryptedMessage
}
