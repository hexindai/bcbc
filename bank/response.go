package bank

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/fatih/color"
)

type (
	// CardBinCheckResponse bank card check result
	CardBinCheckResponse struct {
		CardType  string         `json:"cardType"`
		Bank      string         `json:"bank"`
		Key       string         `json:"key"`
		Messages  []errorMessage `json:"messages"`
		Validated bool           `json:"validated"`
		Stat      string         `json:"stat"`
		CardBIN   string
	}

	errorMessage struct {
		Name       string `json:"name"`
		ErrorCodes string `json:"errorCodes"`
	}

	successJSONOutput struct {
		Bin    string `json:"bin"`
		Bank   string `json:"bank"`
		Name   string `json:"name"`
		Type   string `json:"type"`
		Length int    `json:"length"`
	}

	failJSONOutput struct {
		Validated bool   `json:"validated"`
		No        string `json:"no"`
		Messages  string `json:"messages"`
	}
	// ContentType CardBinCheckResponse content type
	ContentType int
)

const (
	// JSONContentType JSON content type
	JSONContentType ContentType = iota
	// TextContentType JSON content type
	TextContentType
)

func (cbcr CardBinCheckResponse) toJSON() string {
	var output interface{}
	if cbcr.Stat == "ok" && cbcr.Validated {
		output = successJSONOutput{
			Bin:    cbcr.CardBIN,
			Bank:   cbcr.Bank,
			Name:   bankName(cbcr.Bank),
			Type:   cbcr.CardType,
			Length: length(cbcr.Key),
		}
	} else {
		output = failJSONOutput{
			Validated: false,
			No:        cbcr.Key,
			Messages:  inlineMessages(cbcr.Messages),
		}
	}
	b, _ := json.Marshal(output)
	return string(b)
}

func (cbcr CardBinCheckResponse) toText() string {
	if cbcr.Stat == "ok" && cbcr.Validated {
		return fmt.Sprintf(`
Validated: %s
BIN: %s
No: %s (%v位)
Type: %s
Bank: %s
BankName: %s
`, color.GreenString("✔"), cbcr.CardBIN, cbcr.Key, length(cbcr.Key), cbcr.CardType, cbcr.Bank, bankName(cbcr.Bank))
	}
	var m string
	for _, msg := range cbcr.Messages {
		m += fmt.Sprintf("\n%s %v: %v;", color.YellowString("!"), msg.Name, msg.ErrorCodes)
	}
	return fmt.Sprintf(`
Validated: %s
No: %s
Messages %s
`, color.RedString("✘"), cbcr.Key, m)
}

// WriteResponse writes response to io.Writer
func (cbcr CardBinCheckResponse) WriteResponse(w io.Writer, c ContentType) error {
	var err error
	switch c {
	case JSONContentType:
		_, err = io.WriteString(w, cbcr.toJSON())
	case TextContentType:
		_, err = io.WriteString(w, cbcr.toText())
	default:
		err = errors.New("No supporting content type")
	}
	return err
}

// FetchCardBinCheckByCard get cardBinCheckResponse by card
func FetchCardBinCheckByCard(card string) CardBinCheckResponse {
	var cbcr CardBinCheckResponse

	cb, err := defaultCardBINStore.Get(card)
	if err != nil {
		return makeErrorCardBinCheckResponse(card, err)
	}

	cbcr.Stat = "ok"
	cbcr.Key = card
	cbcr.Validated = true
	cbcr.CardType = cb.Type
	cbcr.Bank = cb.Bank
	cbcr.CardBIN = cb.Bin

	return cbcr
}

func makeErrorCardBinCheckResponse(card string, err error) CardBinCheckResponse {
	var cbcr CardBinCheckResponse
	cbcr.Validated = false
	cbcr.Key = card

	ms := strings.SplitN(err.Error(), ": ", 2)
	var k, v string
	if len(ms) < 2 {
		k, v = "internal", ms[0]
	} else {
		k, v = ms[0], ms[1]
	}
	cbcr.Messages = append(cbcr.Messages, errorMessage{k, v})
	return cbcr
}

func bankName(b string) string {
	if name, ok := BankNameMap[b]; ok {
		return name
	}
	return "Unknown"
}

func length(s string) int {
	return len([]rune(s))
}

func inlineMessages(msgs []errorMessage) string {
	var m string
	for i, msg := range msgs {
		var p string
		if i != 0 {
			p = "; "
		}
		m += fmt.Sprintf("%s%v: %v", p, msg.Name, msg.ErrorCodes)
	}
	return m
}
