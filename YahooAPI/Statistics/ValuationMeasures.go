package Statistics

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"stocks/Presets"
	"stocks/YahooAPI/RespContracts"
)

type marketCap struct {
	Raw     uint64
	Fmt     string
	LongFmt string
}
func (m *marketCap) decode(rawData map[string]interface{}) error {
	if _, ok := rawData[m.paramName()]; !ok {
		return fmt.Errorf("expected %s in %v", m.paramName(), rawData)
	}
	return mapstructure.Decode(rawData[m.paramName()], m)
}
func (m *marketCap) paramName() string { return "marketCap" }

type enterpriseValue struct {
	Raw     uint64
	Fmt     string
	LongFmt string
}
func (e *enterpriseValue) decode(rawData map[string]interface{}) error { return mapstructure.Decode(rawData[e.paramName()], e) }
func (e *enterpriseValue) paramName() string { return "enterpriseValue" }

type trailingEPs struct {
	Raw     uint64
	Fmt     string
}
func (t *trailingEPs) decode(mapData Presets.MapDataType) error { return mapstructure.Decode(mapData[t.paramName()], t) }
func (t *trailingEPs) paramName() string { return "trailingEps" }

type ValuationMeasures struct {
	MarketCap       marketCap
	EnterpriseValue enterpriseValue
	TrailingEPs     trailingEPs
	ForwardEPs      float32
	PEGRatio        float32
	PriceSales      float32
	PriceBook       float32
	EntValueRevenue float32
	EntValueEBITDA  float32
}

func (v *ValuationMeasures) RequestParams() []string { return []string{"defaultKeyStatistics", "price"} }
func (v *ValuationMeasures) ResponseContract() RespContracts.ResponseContract { return &RespContracts.QuoteSummary{} }
func (v *ValuationMeasures) decodePrice(mapData Presets.MapDataType) error { return v.MarketCap.decode(mapData) }
func (v *ValuationMeasures) decodeDefaultKeyStatistic(mapData Presets.MapDataType) (err error) {
	err = v.EnterpriseValue.decode(mapData)
	if err != nil {
		return
	}
	err = v.TrailingEPs.decode(mapData)
	if err != nil {
		return
	}
	return
}
func (v *ValuationMeasures) Decode(rawData Presets.RawDataType) (err error) {
	contract := v.ResponseContract()
	err = contract.Check(rawData)
	if err != nil {
		return
	}
	for _, data := range contract.List() {
		for _, param := range v.RequestParams() {
			value, ok := data[param]
			if !ok {
				return fmt.Errorf("expected %s in response", param)
			}
			switch param {
			case "price":
				err = v.decodePrice(value.(map[string]interface{}))
			case "defaultKeyStatistics":
				err = v.decodeDefaultKeyStatistic(value.(map[string]interface{}))
			}
			if err != nil {
				return
			}
		}
	}
	return
}
