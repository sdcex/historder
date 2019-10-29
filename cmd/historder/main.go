package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/sdcex/historder/pkg/models"
	"github.com/sdcex/historder/pkg/operations"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

const (
	cfgName  = "historder"
	auth0URL = "https://sdce.au.auth0.com/oauth/token"
)

type authConfig struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type apiConfig struct {
	Base string `json:"base"`
	Path string `json:"path"`
}

type searchConfig struct {
	FromDate string   `json:"fromDate"`
	ToDate   string   `json:"toDate"`
	Side     string   `json:"side"`
	Ticker   string   `json:"ticker"`
	Status   []string `json:"status"`
}

func main() {
	err := os.Chdir(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	auth, api, search, err := getConfig()
	if err != nil {
		return
	}

	token, err := getToken(auth)
	if err != nil {
		return
	}
	var (
		pageIndex int64
		count     int
		titles    = []string{
			"QRID",
			"Status",
			"Ticker",
			"Side",
			"CreatedAt",
			"ExecutedAt",
			"PayFundDetail",
			"ReceiveFundDetail",
			"BuyUnitPrice",
			"SellUnitPrice",
			"FeeValue",
			"FeeCurrency",
			"Amount",
			"Quantity",
			"ExtraInfo"}
	)
	tb := models.NewTable(titles)
	for {
		orders, err := getRecord(api, search, token, 500, pageIndex)
		if err != nil || len(orders) == 0 {
			break
		}
		count += len(orders)
		log.Infof("collect %v records", count)
		for _, order := range orders {
			row := mapOrder(order)
			err = tb.AddRowList(row)
			if err != nil {
				log.Error(err)
			}
		}
		pageIndex++
	}
	err = tb.Save()
	if err != nil {
		log.Error(err)
	}
	log.Infof("%v rows are converted to csv", count)
	log.Info("Tool exit.")
}

func getConfig() (auth authConfig, api apiConfig, search searchConfig, err error) {
	v := viper.New()
	v.SetConfigFile(fmt.Sprintf("config/%s.yaml", cfgName))
	err = v.ReadInConfig()
	if err != nil {
		log.Error(err)
		return
	}
	err = v.UnmarshalKey("auth", &auth)
	if err != nil {
		log.Error(err)
		return
	}
	err = v.UnmarshalKey("api", &api)
	if err != nil {
		log.Error(err)
		return
	}
	err = v.UnmarshalKey("search", &search)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

func getToken(auth authConfig) (string, error) {
	log.Infof("auth config info: %v", auth)
	resp, err := http.PostForm(auth0URL, url.Values{
		"client_id":     {auth.ClientID},
		"client_secret": {auth.ClientSecret},
		"grant_type":    {"client_credentials"},
		"audience":      {"https://api.sdce.com.au"},
	})
	if err != nil {
		log.Errorf("failed to create req from url %v: %v", auth0URL, err)
		return "", err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("cannot Read response body %v", err)
		return "", err
	}
	dataMap := map[string]interface{}{}
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		log.Error(err)
		return "", err
	}
	token, ok := dataMap["access_token"].(string)
	if !ok {
		err = fmt.Errorf("cannot get access token, %+v", dataMap)
		log.Error(err)
		return "", err
	}
	return token, nil
}

func getRecord(api apiConfig, search searchConfig, token string, pageSz, pageIndex int64) ([]*models.MerchantOrder, error) {

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: time.Minute,
	}

	url := fmt.Sprintf("%v%v?pageSize=%v&pageIndex=%v", api.Base, api.Path, pageSz, pageIndex)
	if search.FromDate != "" {
		url += ("&fromDate=" + search.FromDate)
	}
	if search.ToDate != "" {
		url += ("&toDate=" + search.ToDate)
	}
	if search.Side != "" {
		url += ("&side=" + search.Side)
	}
	if search.Ticker != "" {
		url += ("&ticker=" + search.Ticker)
	}
	if len(search.Status) != 0 {
		url += "&status="
		for i, stts := range search.Status {
			if i != 0 {
				url += ","
			}
			url += stts
		}
	}
	log.Infof("GET %v", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("failed to create req from url %v: %v", url, err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("failed to get data from url %v: %v", url, err)
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("cannot Read response body %v", err)
		return nil, err
	}
	var body operations.GetMerchantOrdersOKBody
	err = json.Unmarshal(data, &body)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Infof("Receive %v items", len(body.Result))
	return body.Result, nil
}

func mapOrder(order *models.MerchantOrder) []string {
	data := []string{}
	data = append(data, *order.CounterPartyRequestID)
	data = append(data, *order.Status)
	data = append(data, order.Ticker)
	data = append(data, *order.Side)
	data = append(data, order.CreatedAt.String())
	data = append(data, order.ExecutedAt.String())
	data = append(data, order.PayFundDetail)
	data = append(data, order.ReceiveFundDetail)
	if order.CurrencyQuote.BuyUnitPrice != nil {
		data = append(data, strconv.FormatFloat(*order.CurrencyQuote.BuyUnitPrice.Price, 'f', 4, 64))
	} else {
		data = append(data, "")
	}
	if order.CurrencyQuote.SellUnitPrice != nil {
		data = append(data, strconv.FormatFloat(*order.CurrencyQuote.SellUnitPrice.Price, 'f', 4, 64))
	} else {
		data = append(data, "")
	}
	if order.CurrencyQuote.Fee != nil {
		data = append(data, *order.CurrencyQuote.Fee.Value)
		data = append(data, *order.CurrencyQuote.Fee.Currency)
	} else {
		data = append(data, "", "")
	}
	data = append(data, *order.CurrencyQuote.Amount.Amount)
	data = append(data, *order.CurrencyQuote.Quantity.Quantity)
	data = append(data, order.ExtraInfo)
	return data
}
