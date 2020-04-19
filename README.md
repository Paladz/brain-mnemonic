# brain-mnemonic
A command-line tool for using any message in your mind to generate 12 mnemonic words for blockchain wallet. This tool is designed for 3 intentions.
  1. Don't need to prevent any software/hardware wallet provider leave the back door while generating your mnemonic words.
  2. Don't worry about forgetting your mnemonic or the mnemonic got stolen。
  3. One brain message can generate unlimited mnemonics and each one is security isolation.

## Usage
```
go run mnemonic.go <message> <deriver-index>
```
**message**(_string_, _require_): The secret message in your brain, you can recover all your mnemonics as long as you remember this.
**deriver-index**(_int_, _option_): An int index for make one message generate different isolation mnemonics, the default is 1

## Example
create simple brain mnemonic
```
go run mnemonic.go "whatever in your mind"

your brain message is "whatever in your mind", the deriver index is 1
this is your mnemonic: "siren hand term grab dignity entire bike grace fuel document grace drip gate bench pioneer save absorb ostrich inmate grunt sea horse clinic subway"
bitcoin mainnet adress is: 1KB8WHSjrL146WVLdsabvYLkcbkPkNRwD8
etherum mainnet adress is: 0x16d03f4B955785A409db79c85Fec9eC367cCc761
```

create your 73th brain mnemonic with the same secret message
```
go run mnemonic.go "whatever in your mind" 73

your brain message is "whatever in your mind", the deriver index is 73
this is your mnemonic: "nephew eager harsh nerve layer clock obtain task diary stove morning stem install student wise survey decline shy neutral nation script enjoy tornado panic"
bitcoin mainnet adress is: 18HWNuEPyvPkTZYVergsZDS12kiJqm2q2g
etherum mainnet adress is: 0xe2ce59467F3C0AE2b65b8e1CC1e9a31F7De3C956
```

## Security
1. This is an open-source repo that you can trust and verify. We use go mod instead of go vendor to prevent anyone inserts back door in vendor files.
2. This project sticks to use go native library. The only 3rd parity library is [go-bip39](https://github.com/tyler-smith/go-bip39), But we use go mod to make sure we use the same version as [Ethereum](https://github.com/ethereum/go-ethereum/releases/tag/v1.9.12)(which already verify by thousands of developers).
3. The only weak point of this tool is using a weak message to generate mnemonics. Please don't use birthday, phone number, or anything easy to guess.
