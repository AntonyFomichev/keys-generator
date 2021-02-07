package main

import (
	"log"
	"math/big"
	"math/rand"
	"time"
)

var one = big.NewInt(1)

func makeBigInt(number string) *big.Int {
	i, success := new(big.Int).SetString(number, 10)

	if !success {
		log.Fatal("Failed to create BigInt from string")
	}

	return i
}

func RandBigInt(min *big.Int, max *big.Int) *big.Int {
	source := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	return big.NewInt(0).Add(min, big.NewInt(0).Rand(source, big.NewInt(0).Add(big.NewInt(0).Sub(max, min), big.NewInt(1))))
}
