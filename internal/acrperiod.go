package internal

import (
	"time"

	"github.com/shopspring/decimal"
)

type Acrperiod struct {
	// LastUpdate Dato for siste oppdatering
	LastUpdate time.Time `json:"last_update"`
	// PayDate2 Betalingsdato 2 for lønn
	PayDate2 time.Time `json:"pay_date2"`
	// PayDate Betalingsdato for lønn
	PayDate time.Time `json:"pay_date"`
	// DateFrom Dato fra
	DateFrom time.Time `json:"date_from"`
	// DateTo Dato til
	DateTo time.Time `json:"date_to"`
	// Base U4 base
	Base string `json:"base"`
	// UserID Bruker-ID
	UserID string `json:"user_id"`
	// Password Passord
	Password string `json:"password"`
	// Status Status
	Status string `json:"status"`
	// Client Klientkode
	Client string `json:"client"`
	// Description Beskrivelse
	Description string `json:"description"`
	// PeriodID Periode-ID
	PeriodID string `json:"period_id"`
	// Value1 Fritt decimal-felt 1
	Value1 decimal.Decimal `json:"value_1"`
	// AgrtID Rad-ID
	AgrtID int64 `json:"agrtid"`
	// Period Periode
	Period int `json:"period"`
	// AccPeriod Regnskapsperiode i hovedbok
	AccPeriod int `json:"acc_period"`
	// ConsPeriod Konsolideringsperiode
	ConsPeriod int `json:"cons_period"`
	// FiscPeriod Regnskapsperiode
	FiscPeriod int `json:"fisc_period"`
	// FiscalYear Regnskapsår
	FiscalYear int `json:"fiscal_year"`
} // @name Acrperiod

type AcrperiodResponse struct {
	Data     []Acrperiod    `json:"data"`
	Metadata MetadataCursor `json:"metadata"`
}
