package main

import (
	"crypto/sha256"
	"log"
	"os"
	"strconv"

	"github.com/tyler-smith/go-bip39"
)

const (
	minDeriverIndex = 1
	maxDeriverIndex = 4096
)

func main() {
	brainMessage := readBrainMessage()
	deriverIndex := readDeriverIndex()

	mnemonic, err := newMnemonic(brainMessage, deriverIndex)
	if err != nil {
		log.Fatalf("fail to create mnemonic: %v", err)
	}

	log.Printf("your brain message is \"%s\", the deriver index is %d", brainMessage, deriverIndex)
	log.Printf("this is your mnemonic: \"%s\"", mnemonic)
}

func newMnemonic(brainMessage string, deriverIndex int) (string, error) {
	entropy := []byte(brainMessage)
	for i := 0; i < deriverIndex; i++ {
		hashData := sha256.Sum256(entropy)
		entropy = hashData[:]
	}

	return bip39.NewMnemonic(entropy)
}

func readBrainMessage() string {
	if len(os.Args) < 2 {
		log.Fatal("please input the secret message in your brain to generate brain mnemonic")
	}

	return os.Args[1]
}

func readDeriverIndex() int {
	if len(os.Args) < 3 {
		return 1
	}

	deriverIndex, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("fail to read deriver index: %v", err)
	}

	if deriverIndex < minDeriverIndex || deriverIndex > maxDeriverIndex {
		log.Fatalf("the deriver index should in range [%d, %d]", minDeriverIndex, maxDeriverIndex)
	}

	return deriverIndex
}
