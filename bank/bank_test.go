// Copyright 2018 Runrioter. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package bank

import "testing"

var cb103_19 = &CardBIN{Bin: "103", Bank: "ABC", Type: "DC", Length: 19}
var cb18572_18 = &CardBIN{Bin: "18572", Bank: "KMRCU", Type: "DC", Length: 18}

// same card bin, but different card lengths
var cb622309_19 = &CardBIN{Bin: "622309", Bank: "CZBANK", Type: "DC", Length: 19}
var cb622309_18 = &CardBIN{Bin: "622309", Bank: "ICBC", Type: "DC", Length: 18}

// same card bin, different card bin lengths, and different card lengths
var cb621260_16 = &CardBIN{Bin: "621260", Bank: "SPABANK", Type: "CC", Length: 16}
var cb621260_19 = &CardBIN{Bin: "621260", Bank: "CSRCB", Type: "DC", Length: 19}
var cb621260001_19 = &CardBIN{Bin: "621260001", Bank: "XFRCB", Type: "DC", Length: 19}
var cb621260002_19 = &CardBIN{Bin: "621260002", Bank: "ESRCB", Type: "DC", Length: 19}

var testCardBINs = [...]*CardBIN{
	cb103_19,
	cb18572_18,
	cb622309_19,
	cb622309_18,
	cb621260_16,
	cb621260_19,
	cb621260001_19,
	cb621260002_19,
}

var testAll = [...]struct {
	card string
	bin  *CardBIN
}{
	{"10312345678909876543", nil},
	{"1031234567890987654", cb103_19},
	{"103123456789098765", nil},

	{"1857212345678909876", nil},
	{"185721234567890987", cb18572_18},
	{"18572123456789098", nil},

	{"622309123456789098", cb622309_18},
	{"6223091234567890987", cb622309_19},

	{"6212601234567890", cb621260_16},
	{"6212601234567890987", cb621260_19},
	{"6212600011234567890", cb621260001_19},
	{"6212600021234567890", cb621260002_19},
	{"6212600031234567890", cb621260_19},
}

var testBank = New(testCardBINs[:])

func TestBankGet(t *testing.T) {
	for _, ti := range testAll {
		result, _ := testBank.Get(ti.card)
		if result != ti.bin {
			t.Errorf("bank.Get(%s) expect: %v, but got: %v\n", ti.card, ti.bin, result)
		}
	}
}
