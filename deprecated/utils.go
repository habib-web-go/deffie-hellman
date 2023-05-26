package deprecated

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"net/http"
)

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func isValidUser(username string, password string) bool {
	hashedPassword, err := redisClient.HGet(usersRedisSet, username).Result()
	if err != nil {
		return false
	}
	return hashPassword(password) == hashedPassword
}

func createSessionToken(username string) (string, error) {
	token, err := generateToken()
	if err != nil {
		return "", err
	}

	err = redisClient.Set(token, username, 0).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}

func getUsernameBySessionToken(token string) (string, error) {
	username, err := redisClient.Get(token).Result()
	if err != nil {
		return "", err
	}
	return username, nil
}

func generateToken() (string, error) {
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}
	token := base64.URLEncoding.EncodeToString(tokenBytes)
	return token, nil
}

func checkMethodIsPost(r *http.Request) bool {
	if r.Method != http.MethodPost {
		return false
	}
	return true
}
