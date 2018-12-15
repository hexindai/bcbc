// Copyright 2018 Runrioter. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package bank

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
	l, e := validate(card)
	if e != nil {
		return nil, e
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
		e = ErrCardBINNotMatch
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

// validate check card and return the number of valid card numbers
func validate(card string) (int, error) {
	l := 0
	for _, r := range card {
		if r < '0' || r > '9' {
			return l, ErrParamIllegal
		}
		l++
	}
	if l < 10 || l > 19 {
		return l, ErrParamIllegal
	}
	return l, nil
}
