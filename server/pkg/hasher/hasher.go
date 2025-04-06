package hasher

import (
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

type Hasher struct {
	pepper    string // Secret string, same for all passwords
	time      uint32
	memory    uint32
	threads   uint8
	keyLength uint32
}

func NewHasher(pepper string) *Hasher {
	return &Hasher{
		pepper:    pepper,
		time:      1,
		memory:    64 * 1024,
		threads:   2,
		keyLength: 32,
	}
}

func (h Hasher) Get(salt, password string) string {
	passwordWithPepper := password + h.pepper

	hash := argon2.IDKey([]byte(passwordWithPepper), []byte(salt), h.time, h.memory, h.threads, h.keyLength)

	return base64.StdEncoding.EncodeToString(hash)
}
