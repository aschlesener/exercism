package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

const testVersion = 1

// PrivateKey returns a number that is greater than 1 and less than p
func PrivateKey(p *big.Int) *big.Int {
	// handle range requirement and generate random number within that range
	num := big.NewInt(0)
	num = num.Add(p, big.NewInt(-2))
	randNum, _ := rand.Int(rand.Reader, num)

	result := randNum.Add(randNum, big.NewInt(2))
	return result
}

// PublicKey returns a public key A based on the formula
// A = g^a mod p
// where g is the generator, a is a secret key, and p is a prime
func PublicKey(a *big.Int, p *big.Int, g int64) *big.Int {
	bigG := big.NewInt(g)
	bigG.Exp(bigG, a, p)
	return bigG
}

// SecretKey returns a secret key s based on the formula
// s = B^a mod p
// where B is a public key, a is a private key, and p is a prime
func SecretKey(a *big.Int, B *big.Int, p *big.Int) *big.Int {
	key := big.NewInt(0)
	key.Exp(B, a, p)
	return key
}

// NewPair generates a private and public key given a prime number and generator
func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	privateKey := PrivateKey(p)
	publicKey := PublicKey(privateKey, p, g)
	return privateKey, publicKey
}
