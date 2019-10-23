package models

import (
	strfmt "github.com/go-openapi/strfmt"
)

type Amount struct {

	// amount
	// Required: true
	Amount *string `json:"amount"`

	// currency
	// Required: true
	Currency *string `json:"currency"`
}

// MerchantOrderCurrencyQuoteFee merchant order currency quote fee
// swagger:model MerchantOrderCurrencyQuoteFee
type MerchantOrderCurrencyQuoteFee struct {

	// currency
	// Required: true
	Currency *string `json:"currency"`

	// value
	// Required: true
	Value *string `json:"value"`
}

type MerchantOrderCurrencyQuote struct {

	// amount
	// Required: true
	Amount *Amount `json:"amount"`

	// buy unit price
	// Required: true
	BuyUnitPrice *UnitPrice `json:"buyUnitPrice"`

	// fee
	// Required: true
	Fee *MerchantOrderCurrencyQuoteFee `json:"fee"`

	// quantity
	// Required: true
	Quantity *Quantity `json:"quantity"`

	// sell unit price
	SellUnitPrice *UnitPrice `json:"sellUnitPrice,omitempty"`
}

// MerchantOrderMchMargin merchant order mch margin
// swagger:model MerchantOrderMchMargin
type MerchantOrderMchMargin struct {

	// amount
	// Required: true
	Amount *float64 `json:"amount"`

	// currency
	// Required: true
	Currency *string `json:"currency"`
}

// PaymentDetail payment detail
// swagger:model PaymentDetail
type PaymentDetail struct {

	// bank detail
	// Required: true
	BankDetail *string `json:"bankDetail"`

	// bank name
	// Required: true
	BankName *string `json:"bankName"`

	// card number
	// Required: true
	CardNumber *string `json:"cardNumber"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Quantity Quantity includes quantity and currency. Client can quote an exchange based on currency and quantity
// swagger:model Quantity
type Quantity struct {

	// currency
	// Required: true
	Currency *string `json:"currency"`

	// quantity
	// Required: true
	Quantity *string `json:"quantity"`
}

// UnitPrice unit price
// swagger:model UnitPrice
type UnitPrice struct {

	// currency
	// Required: true
	Currency *string `json:"currency"`

	// price
	// Required: true
	Price *float64 `json:"price"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

type MerchantOrder struct {

	// amount
	// Required: true
	Amount *Amount `json:"amount"`

	// counter party request Id
	// Required: true
	CounterPartyRequestID *string `json:"counterPartyRequestId"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// currency quote
	// Required: true
	CurrencyQuote *MerchantOrderCurrencyQuote `json:"currencyQuote"`

	// date time that the trade is executed
	// Required: true
	// Format: date-time
	ExecutedAt *strfmt.DateTime `json:"executedAt"`

	// extra info
	ExtraInfo string `json:"extraInfo,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// mch margin
	// Required: true
	MchMargin *MerchantOrderMchMargin `json:"mchMargin"`

	// memo
	Memo string `json:"memo,omitempty"`

	// merchant Id
	// Required: true
	MerchantID *string `json:"merchantId"`

	// otc history
	OtcHistory string `json:"otcHistory,omitempty"`

	// pay fund detail
	PayFundDetail string `json:"payFundDetail,omitempty"`

	// payment detail
	PaymentDetail *PaymentDetail `json:"paymentDetail,omitempty"`

	// payment proof
	PaymentProof string `json:"paymentProof,omitempty"`

	// pickup Url
	PickupURL string `json:"pickupUrl,omitempty"`

	// quantity
	// Required: true
	Quantity *Quantity `json:"quantity"`

	// receive fund detail
	ReceiveFundDetail string `json:"receiveFundDetail,omitempty"`

	// side
	// Required: true
	Side *string `json:"side"`

	// status
	// Required: true
	Status *string `json:"status"`

	// ticker
	Ticker string `json:"ticker,omitempty"`

	// unit price
	// Required: true
	UnitPrice *UnitPrice `json:"unitPrice"`
}
