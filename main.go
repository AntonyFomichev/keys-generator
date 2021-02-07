package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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

	if err != nil {
        log.Fatal(err)
  	}

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
				time.Sleep(1000 * time.Millisecond)
				printBitcoinKeys(keysPerPage)
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

	if resp.StatusCode != 200 {
		panic(resp.Status)
	}

	fmt.Print("StatusCode: " + strconv.Itoa(resp.StatusCode) + " ")

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
