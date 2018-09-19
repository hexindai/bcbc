package bank

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

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
		CardNo    string
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
			Bin:    cardBin(cbcr.CardNo),
			Bank:   cbcr.Bank,
			Name:   bankName(cbcr.Bank),
			Type:   cbcr.CardType,
			Length: length(cbcr.CardNo),
		}
	} else {
		output = failJSONOutput{
			Validated: false,
			No:        cbcr.CardNo,
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
No: %s (%v位)
Type: %s
Bank: %s
BankName: %s
`, color.GreenString("✔"), cbcr.CardNo, length(cbcr.CardNo), cbcr.CardType, cbcr.Bank, bankName(cbcr.Bank))
	}
	var m string
	for _, msg := range cbcr.Messages {
		m += fmt.Sprintf("\n%s %v: %v;", color.YellowString("!"), msg.Name, msg.ErrorCodes)
	}
	return fmt.Sprintf(`
Validated: %s
No: %s
Messages %s
`, color.RedString("✘"), cbcr.CardNo, m)
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

	l := length(card)

	if l < 10 || l > 19 {
		err := errors.New("cardNo:PARAM_ILLEGAL")
		return makeErrorCardBinCheckResponse(card, err)
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", makeURL(card), nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.92 Safari/537.36")
	req.Header.Add("Host", "ccdcapi.alipay.com")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7")
	resp, err := client.Do(req)
	if err != nil {
		return makeErrorCardBinCheckResponse(card, err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&cbcr)
	if err != nil {
		return makeErrorCardBinCheckResponse(card, err)
	}
	cbcr.CardNo = card

	return cbcr
}

func makeErrorCardBinCheckResponse(card string, err error) CardBinCheckResponse {
	var cbcr CardBinCheckResponse
	cbcr.Validated = false
	cbcr.CardNo = card
	cbcr.Messages = append(cbcr.Messages, errorMessage{"http.Get", err.Error()})
	return cbcr
}

func makeURL(cardNo string) string {
	u := "https://ccdcapi.alipay.com/validateAndCacheCardInfo.json"
	v := url.Values{}
	v.Set("_input_charset", "utf-8")
	v.Set("cardBinCheck", "true")
	v.Set("cardNo", cardNo)
	return u + "?" + v.Encode()
}

func length(s string) int {
	return len([]rune(s))
}

func bankName(b string) string {
	if name, ok := BankNameMap[b]; ok {
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
