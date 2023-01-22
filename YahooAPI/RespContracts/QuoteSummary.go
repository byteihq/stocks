package RespContracts

import (
	"encoding/json"
	"stocks/Presets"
)

type quoteSummaryContext struct {
	Result []map[string]interface{} `json:"result"`
	Error error `json:"error"`
}

type quoteSummaryImpl struct {
	QuoteSummaryImpl quoteSummaryContext `json:"quoteSummary"`
}

type QuoteSummary struct {
	quoteSummary quoteSummaryImpl
}

func (q *QuoteSummary) Check(rawData Presets.RawDataType) (err error) {
	err = json.Unmarshal(rawData, &q.quoteSummary)
	if err != nil {
		return
	}
	err = q.quoteSummary.QuoteSummaryImpl.Error
	return
}

func (q *QuoteSummary) List() []map[string]interface{} {
	return q.quoteSummary.QuoteSummaryImpl.Result
}
