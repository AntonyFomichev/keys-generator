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
	file, err := os.Create("keys.txt")

	fmt.Fprintln(file, time.Now().UTC().Format("15:04:05") + "\n\n")

	if err != nil {
    log.Fatal(err)
  }

	for {
		res := printBitcoinKeys(keysPerPage)

		if (res) {
			break;
		}

		time.Sleep(1000 * time.Millisecond)
	}
}

func printBitcoinKeys(keysPerPage int)bool {
	compressedKeys := ""
	bitcoinKeys := generateBitcoinKeys(keysPerPage)

	length := len(bitcoinKeys)

	for i, key := range bitcoinKeys {
		compressedKeys += key.compressed

		if i != length - 1 {
			compressedKeys += ","
		}
	}

	foundedBalance := checkBtcBalanceWallet(compressedKeys)

	if foundedBalance != "" {
		for i, key := range bitcoinKeys {
			if key.compressed == foundedBalance {
				file, err := os.Open("keys.txt")

				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Wallet found! " + key.private)
				fmt.Fprintln(file, i, key)

				return true
			}
		}
	} else {
		fmt.Println(time.Now().UTC().Format("15:04:05") + " Nothing found..." + "\n")
		return false
	}

	return false
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
