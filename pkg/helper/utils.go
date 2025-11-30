package helper

import (
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password using bcrypt.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPasswordHash compares a password with a hash.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func StringToUUID(id string) uuid.UUID {
	return uuid.MustParse(id)
}

func UUIDToString(id uuid.UUID) string {
	return id.String()
}

func GenerateAccountNumber(isMerchant bool, userID string) string {

	// Convert UUID to a string and take the last 6 characters
	uuidStr := UUIDToString(StringToUUID(userID))
	shortUUID := uuidStr[len(uuidStr)-6:]
	// Remove all non-numeric characters from shortUUID
	reg := regexp.MustCompile("[^0-9]+")
	shortUUID = reg.ReplaceAllString(shortUUID, "")

	for len(shortUUID) < 6 {
		shortUUID += "0"
	}

	// Get current Unix timestamp
	unixTime := time.Now().Unix()
	unixTimeStr := strconv.FormatInt(unixTime, 10)

	// Take the last 4 digits of the Unix timestamp
	var timestampSuffix string
	if len(unixTimeStr) >= 4 {
		timestampSuffix = unixTimeStr[len(unixTimeStr)-4:]
	} else {
		timestampSuffix = unixTimeStr
		for len(timestampSuffix) < 4 {
			timestampSuffix = "0" + timestampSuffix
		}
	}

	var userCode string
	if isMerchant {
		userCode = "30"
	} else {
		userCode = "50"
	}

	// Generate a random 4-digit number
	randomNumber := rand.Intn(9000) + 1000 // Generates a number between 1000 and 9999

	// Combine them to create a unique account number
	accountNumber := userCode + shortUUID + strconv.Itoa(randomNumber) + timestampSuffix

	return accountNumber

}
