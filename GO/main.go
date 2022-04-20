package main

import (
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"strings"
)

type HashedSolution struct {
	Hash    string `json:"hash"`
	Postfix int    `json:"postfix"`
}

func getHash(input string, paddingLength int) string {
	var list []HashedSolution
	padding := strings.Repeat("0", paddingLength)

	for postFix := 0; ; postFix++ {
		str := input + strconv.Itoa(postFix)
		hash := sha256.New()
		hash.Write([]byte(str))
		encodedHash := hex.EncodeToString(hash.Sum(nil))

		if strings.HasPrefix(encodedHash, padding) {
			for i := 0; i < 10; i++ {
				list = append(list, HashedSolution{
					Postfix: postFix,
					Hash:    encodedHash,
				})
			}

			jsonBytes, _ := json.Marshal(list)

			return b64.StdEncoding.EncodeToString(jsonBytes)
		}
	}
}
