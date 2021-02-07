# Keys.lol generator

This repository contains the key generator for [Keys.lol](https://keys.lol)

I modified the automatic generation of wallets, checking the balance on them via the blockchain.info API and, when something found, writing non-empty wallets to keys.txt.

Note: This method was created for fun and my interest in Go. The probability that you will find a non-empty wallet is so small that a meteorite is more likely to hit your house.
Also keep in mind that this is the final version of the script and further actions with non-empty wallets, depending on your jurisdiction, can be considered theft.
If you find some non-empty wallet: rejoice, delete the keys.txt file and go buy a lottery ticket :)

## Building and installing

1. cd to `~/go/src/github.com/sjorso/keys-generator`
2. install required packages with `go get`
3. build the executable with `go build`
4. include the executable in `$PATH`: `sudo cp keys-generator /usr/local/bin`

## Usage

For generating keys, run:

```bash
keys-generator
```

## License

This project is open-sourced software licensed under the [MIT license](http://opensource.org/licenses/MIT)
