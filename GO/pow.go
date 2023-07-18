package queueit

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

type Solution struct {
	Postfix int64  `json:"postfix"`
	Hash    string `json:"hash"`
}

func GetHash(input string, complexity, runs int) ([]Solution, error) {
	solutions := make([]Solution, runs)
	var postFix int64
	hash := sha256.New()
	var hashBytes [sha256.Size]byte

	inputBytes := []byte(input)
	postFixBytes := make([]byte, 0, 20) // all 64-bit integers will fit

	for i := 0; i < runs; i++ {
		for {
			hash.Reset()
			hash.Write(inputBytes)
			postFixBytes = strconv.AppendInt(postFixBytes[:0], postFix, 10)
			hash.Write(postFixBytes)
			hash.Sum(hashBytes[:0])

			if checkZeroPrefix(hashBytes, complexity) {
				solutions[i] = Solution{
					Postfix: postFix,
					Hash:    hex.EncodeToString(hashBytes[:]),
				}
				postFix++
				break
			}

			postFix++
		}
	}

	return solutions, nil
}

// checkZeroPrefix checks if the first characters are all zeros in hexadecimal presentation
func checkZeroPrefix(data [sha256.Size]byte, count int) bool {
	byteCount := (count + 1) / 2

	for i := 0; i < byteCount-1; i++ {
		if data[i] != 0x00 {
			return false
		}
	}

	if count%2 == 0 {
		if data[byteCount-1] != 0x00 {
			return false
		}
	} else {
		if (data[byteCount-1] & 0xf0) != 0x00 {
			return false
		}
	}

	return true
}
