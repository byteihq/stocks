package YahooAPI

import (
	"errors"
	"io"
	"net/http"
	"stocks/Presets"
	"stocks/YahooAPI/Statistics"
	"strings"
)

func getRequest(stock string, statistic Statistics.Statistic) (buf Presets.RawDataType, err error) {
	url := "https://query2.finance.yahoo.com/v10/finance/quoteSummary/" + stock + "?modules=" + strings.Join(statistic.RequestParams(), ",")
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		return
	}
	buf, err = io.ReadAll(resp.Body)
	return
}

func ValuationMeasures(stock string) (statistic Statistics.ValuationMeasures, err error) {
	buf, err := getRequest(stock, &statistic)
	if err != nil {
		return
	}
	err = statistic.Decode(buf)
	return
}
