package main

import (
	"crypto/rand"
	"encoding/hex"
)

type IdType = string

func generateUID() IdType {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}
