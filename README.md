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
this is your mnemonic: "siren hand term grab dignity entire bike grace fuel document grace drip"
```

create your 73th brain mnemonic with the same secret message
```
go run mnemonic.go "whatever in your mind" 73

your brain message is "whatever in your mind", the deriver index is 73
this is your mnemonic: "nephew eager harsh nerve layer clock obtain task diary stove morning stem"
```
