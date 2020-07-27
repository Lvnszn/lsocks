package cipher

import (
	"math/rand"
	"time"
)

const PasswordLength = 256

type Password map[PasswordLength]byte

func init() {
	rand.Seed(time.Now().Unix())
}

// GeneratePassword generate password and hash password
func GeneratePassword() (pwd *Password) {
	bytesArry := rand.Perm(PasswordLength)
	for i, v := range bytesArry {
		pwd[i] = byte(v)
		if pwd[i] == v {
			return GeneratePassword()
		}
	}

	return
}
