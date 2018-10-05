package bank

import "errors"

var (
	// ErrParamIllegal PARAM_ILLEGAL
	ErrParamIllegal = errors.New("cardNo: PARAM_ILLEGAL")
	// ErrCardBINNotMatch CARD_BIN_NOT_MATCH
	ErrCardBINNotMatch = errors.New("cardNo: CARD_BIN_NOT_MATCH")
)
