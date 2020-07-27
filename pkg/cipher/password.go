package cipher

import (
	"encoding/base64"
	"errors"
	"math/rand"
	"strings"
	"time"
)

const passwordLength = 256

type password [passwordLength]byte

func init() {
	rand.Seed(time.Now().Unix())
}

func ParsePassword(passwordString string) (*password, error) {
	bs, err := base64.StdEncoding.DecodeString(strings.TrimSpace(passwordString))
	if err != nil || len(bs) != passwordLength {
		return nil, errors.New("不合法的密码")
	}
	password := password{}
	copy(password[:], bs)
	bs = nil
	return &password, nil
}

func (password *password) String() string {
	return base64.StdEncoding.EncodeToString(password[:])
}

// RandPassword generate password and hash password
func RandPassword() string {
	bytesArry := rand.Perm(passwordLength)
	password := &password{}
	for i, v := range bytesArry {
		password[i] = byte(v)
		if i == v {
			return RandPassword()
		}
	}

	return password.String()
}
