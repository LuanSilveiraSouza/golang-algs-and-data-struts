package caesarCipher

import (
	"reflect"
	"testing"
)

func TestCaesarCipher(t *testing.T) {
	const message = "Golang is amazing!"
	const shift = 5
	const expectedMessage = "Ltqfsl nx frfensl!"

	var encryptedMessage = Encrypt(message, shift)

	if !reflect.DeepEqual(expectedMessage, encryptedMessage) {
		t.Fatal("Encrypt method not worked as expected")
	}

	var decryptedMessage = Decrypt(expectedMessage, shift)

	if !reflect.DeepEqual(message, decryptedMessage) {
		t.Fatal("Encrypt method not worked as expected")
	}
}
