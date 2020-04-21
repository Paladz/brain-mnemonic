package main

import (
	"crypto/ecdsa"
	"crypto/sha256"

	"log"
	"os"
	"strconv"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
)

const (
	minDeriverIndex = 1
	maxDeriverIndex = 4096
)

var (
	// define by https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki
	bitcoinDeriverPath = []uint32{49, 0, 0, 0, 0}
	etherumDeriverPath = []uint32{44, 60, 0, 0, 0}
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

	// print main address for different chains
	echoAddresses(mnemonic)
}

func echoAddresses(mnemonic string) {
	seed := bip39.NewSeed(mnemonic, "")
	bitcoinAddress, err := calcBitcoinAddress(seed)
	if err != nil {
		log.Fatalf("fail to calculate bitcoin address: %v", err)
	}

	etherumAddress, err := calcEtherumAddress(seed)
	if err != nil {
		log.Fatalf("fail to calculate etherum address: %v", err)
	}

	log.Printf("bitcoin mainnet adress is: %s", bitcoinAddress)
	log.Printf("etherum mainnet adress is: %s", etherumAddress)

}

func calcBitcoinAddress(seed []byte) (string, error) {
	key, err := calcDeriverKey(seed, bitcoinDeriverPath)
	if err != nil {
		return "", err
	}

	// generate the P2SH-P2WPKH Address, 0x0014 is the hard code protocol
	script := addressPubKeyHash.ScriptAddress()
	script = append([]byte{0x00, 0x14}, script...)
	addressScriptHash, err := btcutil.NewAddressScriptHash(script, &chaincfg.MainNetParams)
	if err != nil {
		return "", err
	}

	return addressScriptHash.String(), nil
}

func calcDeriverKey(seed []byte, paths []uint32) (*hdkeychain.ExtendedKey, error) {
	key, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}

	for i, path := range paths {
		// follow by bip44 rule
		if i < 3 {
			path += hdkeychain.HardenedKeyStart
		}

		if key, err = key.Child(path); err != nil {
			return nil, err
		}
	}

	return key, nil
}

func calcEtherumAddress(seed []byte) (string, error) {
	key, err := calcDeriverKey(seed, etherumDeriverPath)
	if err != nil {
		return "", err
	}

	pubkey, err := key.ECPubKey()
	if err != nil {
		return "", err
	}

	address := crypto.PubkeyToAddress(ecdsa.PublicKey(*pubkey))
	return address.Hex(), nil
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
