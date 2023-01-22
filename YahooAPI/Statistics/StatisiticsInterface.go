package Statistics

import "stocks/YahooAPI/RespContracts"

type Statistic interface {
	ResponseContract() RespContracts.ResponseContract
	RequestParams() []string
}
