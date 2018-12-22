package main

import (
	"gopkg.in/go-playground/validator.v9"
	"time"
)

// not sure this would work. This package is just for storing struct representations
// of each httpHandler

type WorkingKeyFields struct {
	CommonFields
}

type CardTransferFields struct {
	CommonFields
	CardInfoFields
	AmountFields
	ToCard string `json:"toCard" binding:"required"`
}

type PurchaseFields struct {
	WorkingKeyFields
	CardInfoFields
	AmountFields
}

type ChangePin struct {
	WorkingKeyFields
	NewPin string `json:"newPIN" binding:"required"`
}

type CommonFields struct {
	SystemTraceAuditNumber int       `json:"systemTraceAuditNumber,omitempty" binding:"required"`
	TranDateTime           time.Time `json:"tranDateTime,omitempty" binding:"required,iso8601"`
	TerminalID             string    `json:"terminalId,omitempty" binding:"required,len=8"`
	ClientID               string    `json:"clientId,omitempty" binding:"required"`
}

type CardInfoFields struct {
	Pan     string `json:"PAN" binding:"required"`
	Pin     string `json:"PIN" binding:"required"`
	Expdate string `json:"expDate" binding:"required"`
}

type AmountFields struct {
	TranAmount       float32 `json:"tranAmount" binding:"required"`
	TranCurrencyCode string  `json:"tranCurrencyCode"`
}

func iso8601(fl validator.FieldLevel) bool {

	dateLayout := time.RFC3339
	_, err := time.Parse(dateLayout, fl.Field().String())
	if err != nil{
		return false
	}
	return true
}
