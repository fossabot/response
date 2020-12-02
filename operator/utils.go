package operator

import (
	"encoding/base64"
	"log"
	"math/rand"
)

// GenEncryptionKey generates an encryption key suitable of use as a cryptographic encryption key in Response.
func GenEncryptionKey() string {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatalf("unable to generate encryption key, this is weird: %s", err.Error())
	}

	return base64.StdEncoding.EncodeToString(key)
}
