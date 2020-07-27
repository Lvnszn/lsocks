package cipher

type Cipher struct {
	encodePassword *password
	decodePassword *password
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
func NewCipher(encodePassword *password) *Cipher {
	decodePassword := new(password)
	for i, v := range encodePassword {
		encodePassword[i] = v
		decodePassword[v] = byte(i)
	}

	return &Cipher{
		encodePassword: encodePassword,
		decodePassword: decodePassword,
	}
}
