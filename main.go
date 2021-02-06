package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	keysPerPage := 128
	printBitcoinKeys(keysPerPage)
}

func printBitcoinKeys(keysPerPage int) {
	compressedKeys := ""
	bitcoinKeys := generateBitcoinKeys(keysPerPage)

	length := len(bitcoinKeys)

	file, err := os.Create("keys.txt")

	file.Truncate(0)

	if err != nil {
        log.Fatal(err)
  	}

	defer file.Close()

	for i, key := range bitcoinKeys {
		compressedKeys += key.compressed

		if i != length - 1 {
			compressedKeys += ","
		}
	}

	findedBalance := checkBtcBalanceWallet(compressedKeys)

	if findedBalance != "" {
		for i, key := range bitcoinKeys {
			if key.compressed == findedBalance {
				fmt.Println("Wallet found! " + key.private)
				fmt.Fprintln(file, i, key)
			}
		}
	} else {
		fmt.Println(time.Now().UTC().Format("15:04:05") + " Nothing found..." + "\n")
		time.Sleep(1000 * time.Millisecond)
		printBitcoinKeys(keysPerPage)
	}
}

func checkBtcBalanceWallet(compressed string) string {
	resp, err := http.Get("https://blockchain.info/balance?cors=true&active=" + compressed)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode == 429 {
		panic("Too many requests")
	}

	type Data struct {
		final_balance int
		n_tx int
		total_received int
	}

	var result map[string]Data


	json.NewDecoder(resp.Body).Decode(&result)

	for i, key := range result {
		if key.final_balance != 0 {
			return i
		}
	}

	return ""
}
