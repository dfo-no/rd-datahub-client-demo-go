package internal

import (
	"strconv"
	"strings"
)

type Acatrans struct {
	TransDate   string `json:"trans_date,omitzero"`
	VoucherDate string `json:"voucher_date,omitzero"`
	Dim1        string `json:"dim_1,omitzero"`
	Dim2        string `json:"dim_2,omitzero"`
	Dim3        string `json:"dim_3,omitzero"`
	Dim4        string `json:"dim_4,omitzero"`
	Dim5        string `json:"dim_5,omitzero"`
	Dim6        string `json:"dim_6,omitzero"`
	Dim7        string `json:"dim_7,omitzero"`
	Att1ID      string `json:"att_1_id,omitzero"`
	Att2ID      string `json:"att_2_id,omitzero"`
	Att3ID      string `json:"att_3_id,omitzero"`
	Att4ID      string `json:"att_4_id,omitzero"`
	Att5ID      string `json:"att_5_id,omitzero"`
	Att6ID      string `json:"att_6_id,omitzero"`
	Att7ID      string `json:"att_7_id,omitzero"`
	AparType    string `json:"apar_type,omitzero"`
	Base        string `json:"base,omitzero"`
	Client      string `json:"client,omitzero"`
	VoucherType string `json:"voucher_type,omitzero"`
	Currency    string `json:"currency,omitzero"`
	Account     string `json:"account,omitzero"`
	Description string `json:"description,omitzero"`
	AccNo       string `json:"acc_no,omitzero"`
	AparID      string `json:"apar_id,omitzero"`
	HeadAccount string `json:"head_account,omitzero"`
	Status      string `json:"status,omitzero"`
	Value1      string `json:"value_1,omitzero"`
	Value2      string `json:"value_2,omitzero"`
	Value3      string `json:"value_3,omitzero"`
	Amount      string `json:"amount,omitzero"`
	Percentage  string `json:"percentage,omitzero"`
	CurAmount   string `json:"cur_amount,omitzero"`
	CashAmount  string `json:"cash_amount,omitzero"`
	AgrtID      int64  `json:"agrtid,omitzero"`
	RepPeriod   int    `json:"rep_period,omitzero"`
	Period      int    `json:"period,omitzero"`
	PayPeriod   int    `json:"pay_period,omitzero"`
	SequenceRef int    `json:"sequence_ref,omitzero"`
	SequenceNo  int    `json:"sequence_no,omitzero"`
	Number1     int    `json:"number_1,omitzero"`
	DCFlag      int    `json:"dc_flag,omitzero"`
	VoucherNo   int    `json:"voucher_no,omitzero"`
}

func (a *Acatrans) ToCSVString() string {
	sb := strings.Builder{}

	sb.WriteString(strings.TrimSpace(a.TransDate) + ";")
	sb.WriteString(strings.TrimSpace(a.VoucherDate) + ";")
	sb.WriteString(strings.TrimSpace(a.Dim1) + ";")
	sb.WriteString(strings.TrimSpace(a.Dim2) + ";")
	sb.WriteString(strings.TrimSpace(a.Dim3) + ";")
	sb.WriteString(strings.TrimSpace(a.Dim4) + ";")
	sb.WriteString(strings.TrimSpace(a.Dim5) + ";")
	sb.WriteString(strings.TrimSpace(a.Dim6) + ";")
	sb.WriteString(strings.TrimSpace(a.Dim7) + ";")
	sb.WriteString(strings.TrimSpace(a.Att1ID) + ";")
	sb.WriteString(strings.TrimSpace(a.Att2ID) + ";")
	sb.WriteString(strings.TrimSpace(a.Att3ID) + ";")
	sb.WriteString(strings.TrimSpace(a.Att4ID) + ";")
	sb.WriteString(strings.TrimSpace(a.Att5ID) + ";")
	sb.WriteString(strings.TrimSpace(a.Att6ID) + ";")
	sb.WriteString(strings.TrimSpace(a.Att7ID) + ";")
	sb.WriteString(strings.TrimSpace(a.AparType) + ";")
	sb.WriteString(strings.TrimSpace(a.Base) + ";")
	sb.WriteString(strings.TrimSpace(a.Client) + ";")
	sb.WriteString(strings.TrimSpace(a.VoucherType) + ";")
	sb.WriteString(strings.TrimSpace(a.Currency) + ";")
	sb.WriteString(strings.TrimSpace(a.Account) + ";")
	sb.WriteString(strings.TrimSpace(a.Description) + ";")
	sb.WriteString(strings.TrimSpace(a.AccNo) + ";")
	sb.WriteString(strings.TrimSpace(a.AparID) + ";")
	sb.WriteString(strings.TrimSpace(a.HeadAccount) + ";")
	sb.WriteString(strings.TrimSpace(a.Status) + ";")
	sb.WriteString(strings.TrimSpace(a.Value1) + ";")
	sb.WriteString(strings.TrimSpace(a.Value2) + ";")
	sb.WriteString(strings.TrimSpace(a.Value3) + ";")
	sb.WriteString(strings.TrimSpace(a.Amount) + ";")
	sb.WriteString(strings.TrimSpace(a.Percentage) + ";")
	sb.WriteString(strings.TrimSpace(a.CurAmount) + ";")
	sb.WriteString(strings.TrimSpace(a.CashAmount) + ";")
	sb.WriteString(strconv.FormatInt(a.AgrtID, 10) + ";")
	sb.WriteString(strconv.Itoa(a.RepPeriod) + ";")
	sb.WriteString(strconv.Itoa(a.Period) + ";")
	sb.WriteString(strconv.Itoa(a.PayPeriod) + ";")
	sb.WriteString(strconv.Itoa(a.SequenceRef) + ";")
	sb.WriteString(strconv.Itoa(a.SequenceNo) + ";")
	sb.WriteString(strconv.Itoa(a.Number1) + ";")
	sb.WriteString(strconv.Itoa(a.DCFlag) + ";")
	sb.WriteString(strconv.Itoa(a.VoucherNo) + ";")

	return sb.String()
}

type AcatransResponse struct {
	Metadata MetadataCursor `json:"metadata,omitzero"`
	Data     []*Acatrans    `json:"data,omitzero"`
}
