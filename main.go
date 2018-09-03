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

type jsonOutput struct {
	Bin    string `json:"bin"`
	Bank   string `json:"bank"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Length int    `json:"length"`
}

var card string
var output string

func init() {
	flag.StringVar(&card, "c", "", "Bank card number to be checked")
	flag.StringVar(&output, "o", "text", "Output format: `text, json`; Default: text")
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

	if output == "json" {
		printCardBinCheckResultJSON(cbc)
	} else {
		printCardBinCheckResultText(cbc)
	}
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
Usage: bcbc -c=<card number> [-o <text|json>]

Flags:
`)
	flag.PrintDefaults()
}

func printCardBinCheckResultText(result cardBinCheckResponse) {
	if result.Stat == "ok" && result.Validated {
		fmt.Println("Validated:", color.GreenString("✔"))
		fmt.Println("No:", result.CardNo, "(", length(result.CardNo), "位", ")")
		fmt.Println("Type:", result.CardType)
		fmt.Println("Bank:", result.Bank)
		fmt.Println("BankName:", bankName(result.Bank))
	} else {
		fmt.Println("Validated:", color.RedString("✘"))
		fmt.Println("No:", result.CardNo)
		fmt.Println("Messages:")
		for _, msg := range result.Messages {
			fmt.Println(color.YellowString("!"), msg.Name, ":", msg.ErrorCodes)
		}
	}
}

func printCardBinCheckResultJSON(result cardBinCheckResponse) {
	if result.Stat == "ok" && result.Validated {
		output := jsonOutput{
			Bin:    cardBin(result.CardNo),
			Bank:   result.Bank,
			Name:   bankName(result.Bank),
			Type:   result.CardType,
			Length: length(result.CardNo),
		}
		b, err := json.Marshal(output)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(b))
	} else {
		output := struct {
			Validated bool
			No        string
			Messages  string
		}{
			Validated: false,
			No:        result.CardNo,
			Messages:  inlineMessages(result.Messages),
		}
		b, err := json.Marshal(output)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(b))
	}
}

func length(s string) int {
	return len([]rune(s))
}

func bankName(bank string) string {
	if name, ok := bankNameMap[bank]; ok {
		return name
	}
	return "Unknown"
}

func cardBin(cardNo string) string {
	if length(cardNo) < 6 {
		return ""
	}
	return cardNo[:6]
}

func inlineMessages(msgs []errorMessage) string {
	var m string
	for _, msg := range msgs {
		m += fmt.Sprintf("%v: %v;", msg.Name, msg.ErrorCodes)
	}
	return m
}
