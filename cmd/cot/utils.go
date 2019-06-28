package main

import (
	"encoding/hex"
	"math/rand"
	"time"
)

func randStr() string {
	randBytes := make([]byte, 8)
	rand.Read(randBytes)
	return hex.EncodeToString(randBytes)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
