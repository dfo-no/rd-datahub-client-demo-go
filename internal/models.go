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
	Metadata Metadata      `json:"metadata"`
	Data     []Agltransact `json:"data"`
}

type Agltransact struct {
	Account       string    `json:"account"`
	Amount        float64   `json:"amount"`
	AparID        string    `json:"apar_id"`
	AparType      string    `json:"apar_type"`
	Att1ID        string    `json:"att_1_id"`
	Att2ID        string    `json:"att_2_id"`
	Att3ID        string    `json:"att_3_id"`
	Att4ID        string    `json:"att_4_id"`
	Att5ID        string    `json:"att_5_id"`
	Att6ID        string    `json:"att_6_id"`
	Att7ID        string    `json:"att_7_id"`
	Client        string    `json:"client"`
	CurAmount     float64   `json:"cur_amount"`
	Currency      string    `json:"currency"`
	DcFlag        int       `json:"dc_flag"`
	Description   string    `json:"description"`
	Dim1          string    `json:"dim_1"`
	Dim2          string    `json:"dim_2"`
	Dim3          string    `json:"dim_3"`
	Dim4          string    `json:"dim_4"`
	Dim5          string    `json:"dim_5"`
	Dim6          string    `json:"dim_6"`
	Dim7          string    `json:"dim_7"`
	ExtInvRef     string    `json:"ext_inv_ref"`
	FiscalYear    int       `json:"fiscal_year"`
	LastUpdate    time.Time `json:"last_update"`
	LineNo        int       `json:"line_no"`
	Number1       int       `json:"number_1"`
	OrderID       int       `json:"order_id"`
	Period        int       `json:"period"`
	SequenceNo    int       `json:"sequence_no"`
	Status        string    `json:"status"`
	TaxCode       string    `json:"tax_code"`
	TaxSystem     string    `json:"tax_system"`
	TransDate     time.Time `json:"trans_date"`
	TransID       int       `json:"trans_id"`
	UpdateFlag    int       `json:"update_flag"`
	UserID        string    `json:"user_id"`
	Value1        float64   `json:"value_1"`
	Value2        float64   `json:"value_2"`
	Value3        float64   `json:"value_3"`
	VoucherDate   time.Time `json:"voucher_date"`
	VoucherNo     int       `json:"voucher_no"`
	VoucherType   string    `json:"voucher_type"`
	AgrtID        int       `json:"agrtid"`
	ExtRef        string    `json:"ext_ref"`
	ExtArchRef    string    `json:"ext_arch_ref"`
	UnroAmount    float64   `json:"unro_amount"`
	UnroCurAmount float64   `json:"unro_cur_amount"`
	Base          string    `json:"base"`
}
