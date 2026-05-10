package internal

import (
	"time"

	"github.com/shopspring/decimal"
)

type Acrperiod struct {
	DateFrom    time.Time       `json:"date_from"`
	DateTo      time.Time       `json:"date_to"`
	LastUpdate  time.Time       `json:"last_update"`
	PayDate     time.Time       `json:"pay_date"`
	PayDate2    time.Time       `json:"pay_date2"`
	Base        string          `json:"base,omitzero"`
	Client      string          `json:"client,omitzero"`
	Description string          `json:"description,omitzero"`
	Password    string          `json:"password,omitzero"`
	PeriodID    string          `json:"period_id,omitzero"`
	Status      string          `json:"status,omitzero"`
	UserID      string          `json:"user_id,omitzero"`
	Value1      decimal.Decimal `json:"value_1"`
	AgrtID      int64           `json:"agrtid"`
	AccPeriod   int             `json:"acc_period"`
	ConsPeriod  int             `json:"cons_period"`
	FiscalYear  int             `json:"fiscal_year"`
	FiscPeriod  int             `json:"fisc_period"`
	Period      int             `json:"period"`
}

type AcrperiodResponse struct {
	Data     []Acrperiod `json:"data"`
	Metadata Metadata    `json:"metadata"`
}
