package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
	"strings"
)

// Use the information from https://footlocker.queue-it.net/challengeapi/pow/challenge/{userID} to solve POW
func getHash(inputString string, complexity, runs int) ([]QueueItPowPostFix, error) {
	Solutions := []QueueItPowPostFix{}
	CurrentRuns := 0
	loop := true
	for postfix := 0; loop; postfix++ {
		shaObj := sha256.New()
		shaObj.Write([]byte(inputString))
		shaObj.Write([]byte(strconv.Itoa(postfix)))
		hash := hex.EncodeToString(shaObj.Sum(nil))

		if strings.HasPrefix(string(hash), strings.Repeat("0", complexity)) {
			Solutions = append(Solutions, QueueItPowPostFix{
				Postfix: postfix,
				Hash:    hash,
			})
			CurrentRuns++
		}
		if CurrentRuns == runs {
			return Solutions, nil
		}
	}
	return []QueueItPowPostFix{}, errors.New("error solving pow")
}

// Postfix struct
type QueueItPowPostFix struct {
	Postfix int    `json:"postfix"`
	Hash    string `json:"hash"`
}

// Solution to send in verify post
type QueueItSolution struct {
	Hash []QueueItPowPostFix `json:"hash"`
	Type string              `json:"type"`
}
