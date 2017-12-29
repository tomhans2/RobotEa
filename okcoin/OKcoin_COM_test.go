package okcoin

import (
	"github.com/openbtc/RobotEa"
	"net/http"
	"testing"
)

var okcom = NewCOM(http.DefaultClient, "", "")

func TestOKCoinCOM_API_GetTicker(t *testing.T) {
	ticker, _ := okcom.GetTicker(goex.BTC_CNY)
	t.Log(ticker)
}