package bank

import (
	"errors"
)

// CardBIN is bank card bin
type CardBIN struct {
	Bin    string
	Bank   string
	Type   string
	Length int
}

// BankName return CardBIN bankname
func (cb CardBIN) BankName() string {
	return BankNameMap[cb.Bank]
}

// Bank bank
type Bank struct {
	initialized bool
	*node
}

type node struct {
	binByte  rune
	children []*node
	cardBINs []*CardBIN
}

// DefaultBank initialized by Default BINS
var defaultBank *Bank

func init() {
	defaultBank = New(CardBINs[:])
}

// New returns a bank point that is initialized by []*CardBIN
func New(cb []*CardBIN) *Bank {
	b := new(Bank)
	root := new(node)
	for _, c := range cb {
		root.Insert(c)
	}
	b.node = root
	b.initialized = true
	return b
}

// Get fetch card BIN via card
func Get(card string) (*CardBIN, error) {
	return defaultBank.Get(card)
}

func (in *node) child(c rune) (*node, bool) {
	for _, n := range in.children {
		if n.binByte == c {
			return n, true
		}
	}
	return nil, false
}

func (in *node) Insert(cb *CardBIN) {
	insert(in, []rune(cb.Bin), cb)
}

func (in *node) Get(card string) (c *CardBIN, e error) {
	l := len(card)
	if l < 10 || l > 19 {
		e = errors.New("cardNo: PARAM_ILLEGAL")
		return
	}
	for _, r := range card {
		n, ok := in.child(r)
		if !ok {
			break
		}
		in = n
		for _, cb := range n.cardBINs {
			if l == cb.Length {
				c = cb
				break
			}
		}
	}
	if c == nil {
		e = errors.New("cardNo: CARD_BIN_NOT_MATCH")
	}
	return
}

func insert(n *node, bin []rune, cb *CardBIN) {
	l := len(bin)
	if l == 0 {
		return
	}
	var nn *node
	if ns, ok := n.child(bin[0]); ok {
		nn = ns
	} else {
		nn = new(node)
		nn.binByte = bin[0]
		n.children = append(n.children, nn)
	}
	if l == 1 {
		nn.cardBINs = append(nn.cardBINs, cb)
		return
	}
	insert(nn, bin[1:], cb)
}
