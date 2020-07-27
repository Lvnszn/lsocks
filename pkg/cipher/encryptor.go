package cipher

type Cipher struct {
	encodePassword *Password
	decodePassword *Password
}

func (c *Cipher) encode(bs []byte) {
	for i, v := range bs {
		bs[i] = c.encodePassword[v]
	}
}

func (c *Cipher) decode(bs []byte) {
	for i, v := range bs {
		bs[i] = c.decodePassword[v]
	}
}

// NewCipher new cipher struct for encode/decode msg
func NewCipher(encodePassword *Password) *Cipher {
	decodePassword := new(Password)
	for i, v := range encodePassword {
		encodePassword[i] = v
		decodePassword[v] = byte(i)
	}

	return &Cipher{
		encodePassword: encodePassword,
		decodePassword: decodePassword,
	}
}
