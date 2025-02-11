package internal

import "time"

type Metadata struct {
	CurrentPage  int `json:"current_page,omitempty"`
	PageSize     int `json:"page_size,omitempty"`
	FirstPage    int `json:"first_page,omitempty"`
	LastPage     int `json:"last_page,omitempty"`
	TotalRecords int `json:"total_records,omitempty"`
	NextCursor   int `json:"next_cursor"`
}

type AgltransactResponse struct {
	Data     []Agltransact `json:"data"`
	Metadata Metadata      `json:"metadata"`
}

type Agltransact struct {
	LastUpdate    time.Time `json:"last_update"`
	VoucherDate   time.Time `json:"voucher_date"`
	TransDate     time.Time `json:"trans_date"`
	ExtInvRef     string    `json:"ext_inv_ref"`
	UserID        string    `json:"user_id"`
	Att2ID        string    `json:"att_2_id"`
	Att3ID        string    `json:"att_3_id"`
	Att4ID        string    `json:"att_4_id"`
	Att5ID        string    `json:"att_5_id"`
	Att6ID        string    `json:"att_6_id"`
	Att7ID        string    `json:"att_7_id"`
	Client        string    `json:"client"`
	Base          string    `json:"base"`
	Currency      string    `json:"currency"`
	AparType      string    `json:"apar_type"`
	Description   string    `json:"description"`
	Dim1          string    `json:"dim_1"`
	Dim2          string    `json:"dim_2"`
	Dim3          string    `json:"dim_3"`
	Dim4          string    `json:"dim_4"`
	Dim5          string    `json:"dim_5"`
	Dim6          string    `json:"dim_6"`
	Dim7          string    `json:"dim_7"`
	Account       string    `json:"account"`
	ExtArchRef    string    `json:"ext_arch_ref"`
	ExtRef        string    `json:"ext_ref"`
	VoucherType   string    `json:"voucher_type"`
	Att1ID        string    `json:"att_1_id"`
	AparID        string    `json:"apar_id"`
	TaxSystem     string    `json:"tax_system"`
	TaxCode       string    `json:"tax_code"`
	Status        string    `json:"status"`
	CurAmount     float64   `json:"cur_amount"`
	OrderID       int       `json:"order_id"`
	VoucherNo     int       `json:"voucher_no"`
	Amount        float64   `json:"amount"`
	UpdateFlag    int       `json:"update_flag"`
	SequenceNo    int       `json:"sequence_no"`
	Value1        float64   `json:"value_1"`
	Period        int       `json:"period"`
	Value3        float64   `json:"value_3"`
	Value2        float64   `json:"value_2"`
	TransID       int       `json:"trans_id"`
	AgrtID        int       `json:"agrtid"`
	LineNo        int       `json:"line_no"`
	UnroCurAmount float64   `json:"unro_cur_amount"`
	FiscalYear    int       `json:"fiscal_year"`
	UnroAmount    float64   `json:"unro_amount"`
	DcFlag        int       `json:"dc_flag"`
	Number1       int       `json:"number_1"`
}

type Aagstd struct {
	Dim8       string  `json:"dim8"`
	Base       string  `json:"base"`
	Dim7       string  `json:"dim7"`
	Dim1       string  `json:"dim1"`
	Dim2       string  `json:"dim2"`
	Dim3       string  `json:"dim3"`
	Dim4       string  `json:"dim4"`
	Dim5       string  `json:"dim5"`
	Dim6       string  `json:"dim6"`
	Client     string  `json:"client"`
	PldAmount  float64 `json:"pld_amount"`
	CashAmount float64 `json:"cash_amount"`
	Amount     float64 `json:"amount"`
	PlbAmount  float64 `json:"plb_amount"`
	PlcAmount  float64 `json:"plc_amount"`
	PlaAmount  float64 `json:"pla_amount"`
	PleAmount  float64 `json:"ple_amount"`
	PlhAmount  float64 `json:"plh_amount"`
	PlgAmount  float64 `json:"plg_amount"`
	PlfAmount  float64 `json:"plf_amount"`
	PliAmount  float64 `json:"pli_amount"`
	AgrtID     int     `json:"agrtid"`
	Period     int     `json:"period"`
}

type AagstdResponse struct {
	Data     []Aagstd `json:"data"`
	Metadata Metadata `json:"metadata"`
}
