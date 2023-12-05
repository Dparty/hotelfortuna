package utils

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"strconv"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

type IdGenertor struct {
	ID uint
}

func init() {
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
}

func GenerteId() uint {
	id := node.Generate().Int64()
	return uint(id)
}

func StringToUint(s string) uint {
	u, _ := strconv.ParseUint(s, 10, 64)
	return uint(u)
}

func UintToString(i uint) string {
	return strconv.FormatUint(uint64(i), 10)
}

const saltSize = 16

func generateRandomSalt(saltSize int) []byte {
	var salt = make([]byte, saltSize)
	_, err := rand.Read(salt[:])
	if err != nil {
		panic(err)
	}
	return salt
}

func HashPassword(password string, salt []byte) string {
	var sha512Hasher = sha512.New()
	sha512Hasher.Write(append([]byte(password), salt...))
	return hex.EncodeToString(sha512Hasher.Sum(nil))
}

func HashWithSalt(password string) (string, []byte) {
	var salt = generateRandomSalt(saltSize)
	hashed := HashPassword(password, salt)
	return hashed, salt
}

func PasswordsMatch(
	hashedPassword,
	currPassword string,
	salt []byte) bool {
	var currPasswordHash = HashPassword(currPassword, salt)
	return hashedPassword == currPasswordHash
}
