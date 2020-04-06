package main

import (
	"crypto/sha256"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
)

const (
	mnemonicWordLength = 12
	minDeriverIndex    = 1
	maxDeriverIndex    = 4096
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

	seed := bip39.NewSeed(mnemonic, "")
	bitcoinAddress, err := calcBitcoinAddress(seed)
	if err != nil {
		log.Fatalf("fail to calculate bitcoin address: %v", err)
	}

	log.Printf("bitcoin mainnet adress is: %s", bitcoinAddress)
}

func calcBitcoinAddress(seed []byte) (string, error) {
	key, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return "", err
	}

	// follow by https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki
	mainAddressPaths := []uint32{hdkeychain.HardenedKeyStart + 44, hdkeychain.HardenedKeyStart + 0, hdkeychain.HardenedKeyStart + 0, 0, 0}
	for _, path := range mainAddressPaths {
		if key, err = key.Child(path); err != nil {
			return "", err
		}
	}

	address, err := key.Address(&chaincfg.MainNetParams)
	if err != nil {
		return "", err
	}

	return address.String(), nil
}

func newMnemonic(brainMessage string, deriverIndex int) (string, error) {
	entropy := []byte(brainMessage)
	for i := 0; i < deriverIndex; i++ {
		hashData := sha256.Sum256(entropy)
		entropy = hashData[:]
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", err
	}

	words := strings.Split(mnemonic, " ")
	return strings.Join(words[:mnemonicWordLength], " "), nil
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
