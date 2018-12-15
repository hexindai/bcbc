// Copyright 2018 Runrioter. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package bank

import "errors"

var (
	// ErrParamIllegal PARAM_ILLEGAL
	ErrParamIllegal = errors.New("bank: PARAM_ILLEGAL")
	// ErrCardBINNotMatch CARD_BIN_NOT_MATCH
	ErrCardBINNotMatch = errors.New("bank: CARD_BIN_NOT_MATCH")
)
