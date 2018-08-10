package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyAcctgPreference struct {
	gorm.Model
	PartyId                string
	FiscalYearStartMonth   string
	FiscalYearStartDay     string
	TaxFormId              string
	CogsMethodId           string
	BaseCurrencyUomId      string
	InvoiceSeqCustMethId   string
	InvoiceIdPrefix        string
	LastInvoiceNumber      string
	LastInvoiceRestartDate time.Time
	UseInvoiceIdForReturns string
	QuoteSeqCustMethId     string
	QuoteIdPrefix          string
	LastQuoteNumber        string
	OrderSeqCustMethId     string
	OrderIdPrefix          string
	LastOrderNumber        string
	RefundPaymentMethodId  string
	ErrorGlJournalId       string
	InvoiceSequenceEnumId  string
	OrderSequenceEnumId    string
	QuoteSequenceEnumId    string

	CreatedAt time.Time
	UpdatedAt time.Time
}
