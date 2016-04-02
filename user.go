package gosrp6

import (
	"crypto/rand"
	"crypto/sha256"
	"io"
)

//Identity is the client Identity, defined I in the spec.
type Identity string

//User holds the salt and verifier
type User struct {
	Salt     []byte
	Verifier []byte
	PrivKey  []byte
}

//Users is a map of the clients, keyed by Identity
type Users map[Identity]User

//CreateUser creates a user and stores it in the users varaible
func CreateUser(identity, password string) (Identity, User, error) {
	salt := make([]byte, 4096)
	rand.Read(salt)
	u := User{}
	u.Salt = salt
	x := createPrivateKey(string(salt), identity, password)
	u.PrivKey = x
	return "", u, nil
}

func createPrivateKey(salt, ident, pw string) []byte {
	privHash := sha256.New()
	io.WriteString(privHash, salt)
	io.WriteString(privHash, ident)
	io.WriteString(privHash, pw)
	return privHash.Sum(nil)
}
