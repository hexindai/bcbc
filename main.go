package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/fatih/color"
)

type cardBinCheckResponse struct {
	CardType  string         `json:"cardType"`
	Bank      string         `json:"bank"`
	Key       string         `json:"key"`
	Messages  []errorMessage `json:"messages"`
	Validated bool           `json:"validated"`
	Stat      string         `json:"stat"`
	CardNo    string
}

type errorMessage struct {
	Name       string `json:"name"`
	ErrorCodes string `json:"errorCodes"`
}

var card string

func init() {
	flag.StringVar(&card, "c", "", "Bank card number to be checked")
}

func main() {

	flag.Usage = usage
	flag.Parse()

	if card == "" {
		usage()
		os.Exit(2)
	}

	resp, err := http.Get(makeURL(card))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var cbc = cardBinCheckResponse{}
	err = json.NewDecoder(resp.Body).Decode(&cbc)
	if err != nil {
		log.Fatalln(err)
	}
	cbc.CardNo = card

	printCardBinCheckResult(cbc)
}

func makeURL(cardNo string) string {
	u := "https://ccdcapi.alipay.com/validateAndCacheCardInfo.json"
	v := url.Values{}
	v.Set("_input_charset", "utf-8")
	v.Set("cardBinCheck", "true")
	v.Set("cardNo", cardNo)
	return u + "?" + v.Encode()
}

func usage() {
	fmt.Fprintf(os.Stderr, `
Usage: bcbc -c=<card number>

Flags:
`)
	flag.PrintDefaults()
}

func printCardBinCheckResult(result cardBinCheckResponse) {
	if result.Stat == "ok" && result.Validated {
		fmt.Println("Validated:", color.GreenString("✔"))
		fmt.Println("No:", result.CardNo)
		fmt.Println("Type:", result.CardType)
		fmt.Println("Bank:", result.Bank)
	} else {
		fmt.Println("Validated:", color.RedString("✘"))
		fmt.Println("No:", result.CardNo)
		fmt.Println("Messages:")
		for _, msg := range result.Messages {
			fmt.Println(color.YellowString("!"), msg.Name, ":", msg.ErrorCodes)
		}
	}
}
