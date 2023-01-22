package RespContracts

import (
	"stocks/Presets"
)

type ResponseContract interface {
	Check(rawData Presets.RawDataType) error
	List() []map[string]interface{}
}
