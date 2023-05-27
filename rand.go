package main

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
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

func initDeffieHellman() (string, uint64, error) {
	groupId, err := strconv.Atoi(os.Getenv("DH_GROUP_ID"))
	if err != nil {
		return "", 0, err
	}

	p, g, err := getGroupParams(groupId)
	if err != nil {
		return "", 0, err
	}

	return p.String(), g.Uint64(), nil
}

func createDeffieHellmanSharedKey(g uint64, p string, A string, b uint64) (string, string) {
	ABig, _ := new(big.Int).SetString(A, 10)
	pBig, _ := new(big.Int).SetString(p, 10)
	bBig := big.NewInt(int64(b))
	gBig := big.NewInt(int64(g))

	B := new(big.Int).Exp(gBig, bBig, pBig)
	sharedKey := new(big.Int).Exp(ABig, bBig, pBig)

	return B.String(), sharedKey.String()
}

func getGroupParams(groupID int) (*big.Int, *big.Int, error) {
	switch groupID {
	case 14:
		return getGroup14Params()
	case 0:
		return big.NewInt(23), big.NewInt(5), nil
	default:
		return nil, nil, fmt.Errorf("unsupported group ID")
	}
}

func getGroup14Params() (*big.Int, *big.Int, error) {
	p, success := new(big.Int).SetString(
		"FFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74"+
			"020BBEA63B139B22514A08798E3404DDEF9519B3CD3A431B302B0A6DF25F1437"+
			"4FE1356D6D51C245E485B576625E7EC6F44C42E9A63A3620FFFFFFFFFFFFFFFF", 16)
	if !success {
		return nil, nil, fmt.Errorf("failed to parse prime number (p)")
	}

	g := big.NewInt(2)

	return p, g, nil
}

func createSHA1(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	hashSum := hash.Sum(nil)
	hashString := hex.EncodeToString(hashSum)

	return hashString
}
