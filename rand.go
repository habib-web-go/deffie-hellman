package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"log"
	"math/big"
)

func randomString(length int) string {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	randomString := base64.URLEncoding.EncodeToString(buffer)
	randomString = randomString[:length]
	return randomString
}

func randomUint() uint64 {
	var randomNumber uint64
	for {
		err := binary.Read(rand.Reader, binary.BigEndian, &randomNumber)
		if err == nil {
			return randomNumber
		}
	}
}

func generatePrime(bits int) (*big.Int, error) {
	prime, err := rand.Prime(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	return prime, nil
}

func generateGenerator(p *big.Int) *big.Int {
	two := big.NewInt(2)
	for g := big.NewInt(2); g.Cmp(p) < 0; g.Add(g, big.NewInt(1)) {
		exp := new(big.Int).Sub(p, big.NewInt(1))
		exp.Div(exp, two)
		remainder := new(big.Int).Exp(g, exp, p)
		if remainder.Cmp(big.NewInt(1)) != 0 {
			return g
		}
	}
	return nil
}
func initDeffieHellman() (uint64, uint64) {
	for {
		p, err := generatePrime(256)
		if err != nil || !p.IsUint64() {
			continue
		}
		g := generateGenerator(p)
		if g != nil && g.IsUint64() {
			return p.Uint64(), g.Uint64()
		}
	}
}
