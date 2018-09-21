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

// Bank bank
type bank struct {
	initialized bool
	*node
}

type node struct {
	binByte   rune
	nextNodes []*node
	cardBIN   *CardBIN
}

// Bank global bank
var defaultCardBINStore *bank

func init() {
	b := new(bank)
	rNode := new(node)
	for _, c := range CardBINs {
		rNode.Insert(c)
	}
	b.node = rNode
	b.initialized = true
	defaultCardBINStore = b
}

func (in *node) nextNode(c rune) (*node, bool) {
	for _, n := range in.nextNodes {
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
		if n, ok := in.nextNode(r); ok {
			if n.cardBIN != nil {
				if l == n.cardBIN.Length {
					c = n.cardBIN
				}
				break
			}
			in = n
			continue
		}
		break
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
	if ns, ok := n.nextNode(bin[0]); ok {
		nn = ns
	} else {
		nn = new(node)
		nn.binByte = bin[0]
		n.nextNodes = append(n.nextNodes, nn)
	}
	if l == 1 {
		nn.cardBIN = cb
		return
	}
	insert(nn, bin[1:], cb)
}
