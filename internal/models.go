package internal

import (
	"time"

	"github.com/shopspring/decimal"
)

type Metadata struct {
	PageSize     int `json:"page_size,omitempty"`
	TotalRecords int `json:"total_records,omitempty"`
	LastSeenID   int `json:"last_seen_id,omitempty"`
}

type AgltransactResponse struct {
	Data     []Agltransact `json:"data"`
	Metadata Metadata      `json:"metadata"`
}

type Agltransact struct {
	LastUpdate    time.Time       `json:"last_update"`
	VoucherDate   time.Time       `json:"voucher_date"`
	TransDate     time.Time       `json:"trans_date"`
	ExtInvRef     string          `json:"ext_inv_ref"`
	UserID        string          `json:"user_id"`
	Att1ID        string          `json:"att_1_id"`
	Att2ID        string          `json:"att_2_id"`
	Att3ID        string          `json:"att_3_id"`
	Att4ID        string          `json:"att_4_id"`
	Att5ID        string          `json:"att_5_id"`
	Att6ID        string          `json:"att_6_id"`
	Att7ID        string          `json:"att_7_id"`
	Client        string          `json:"client"`
	Base          string          `json:"base"`
	Currency      string          `json:"currency"`
	AparType      string          `json:"apar_type"`
	Description   string          `json:"description"`
	Dim1          string          `json:"dim_1"`
	Dim2          string          `json:"dim_2"`
	Dim3          string          `json:"dim_3"`
	Dim4          string          `json:"dim_4"`
	Dim5          string          `json:"dim_5"`
	Dim6          string          `json:"dim_6"`
	Dim7          string          `json:"dim_7"`
	Account       string          `json:"account"`
	ExtArchRef    string          `json:"ext_arch_ref"`
	ExtRef        string          `json:"ext_ref"`
	VoucherType   string          `json:"voucher_type"`
	AparID        string          `json:"apar_id"`
	TaxSystem     string          `json:"tax_system"`
	TaxCode       string          `json:"tax_code"`
	Status        string          `json:"status"`
	CurAmount     decimal.Decimal `json:"cur_amount"`
	Amount        decimal.Decimal `json:"amount"`
	Value1        decimal.Decimal `json:"value_1"`
	Value3        decimal.Decimal `json:"value_3"`
	Value2        decimal.Decimal `json:"value_2"`
	UnroCurAmount decimal.Decimal `json:"unro_cur_amount"`
	UnroAmount    decimal.Decimal `json:"unro_amount"`
	OrderID       int64           `json:"order_id"`
	VoucherNo     int64           `json:"voucher_no"`
	TransID       int64           `json:"trans_id"`
	AgrtID        int64           `json:"agrtid"`
	UpdateFlag    int             `json:"update_flag"`
	SequenceNo    int             `json:"sequence_no"`
	Period        int             `json:"period"`
	LineNo        int             `json:"line_no"`
	FiscalYear    int             `json:"fiscal_year"`
	DcFlag        int             `json:"dc_flag"`
	Number1       int             `json:"number_1"`
}

type Aagstd struct {
	Base       string          `json:"base,omitempty"`
	Dim1       string          `json:"dim1,omitempty"`
	Dim2       string          `json:"dim2,omitempty"`
	Dim3       string          `json:"dim3,omitempty"`
	Dim4       string          `json:"dim4,omitempty"`
	Dim5       string          `json:"dim5,omitempty"`
	Dim6       string          `json:"dim6,omitempty"`
	Dim7       string          `json:"dim7,omitempty"`
	Dim8       string          `json:"dim8,omitempty"`
	Client     string          `json:"client,omitempty"`
	PldAmount  decimal.Decimal `json:"pld_amount"`
	CashAmount decimal.Decimal `json:"cash_amount"`
	Amount     decimal.Decimal `json:"amount"`
	PlbAmount  decimal.Decimal `json:"plb_amount"`
	PlcAmount  decimal.Decimal `json:"plc_amount"`
	PlaAmount  decimal.Decimal `json:"pla_amount"`
	PleAmount  decimal.Decimal `json:"ple_amount"`
	PlhAmount  decimal.Decimal `json:"plh_amount"`
	PlgAmount  decimal.Decimal `json:"plg_amount"`
	PlfAmount  decimal.Decimal `json:"plf_amount"`
	PliAmount  decimal.Decimal `json:"pli_amount"`
	AgrtID     int64           `json:"agrtid"`
	Period     int             `json:"period"`
}

type AagstdResponse struct {
	Data     []Aagstd `json:"data"`
	Metadata Metadata `json:"metadata"`
}
