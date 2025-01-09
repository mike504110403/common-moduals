package countrylist

import (
	"encoding/xml"

	"github.com/valyala/fasthttp"
)

func GetCountry() (CountyItems, error) {
	res := CountyItems{}
	if body, err := sendGet("https://api.nlsc.gov.tw/other/ListCounty"); err != nil {
		return res, err
	} else {
		if err := xml.Unmarshal(body, &res); err != nil {
			return res, err
		}
	}
	return res, nil
}

func GetTown(country string) (TownItems, error) {
	res := TownItems{}
	if body, err := sendGet("https://api.nlsc.gov.tw/other/ListTown/" + country); err != nil {
		return res, err
	} else {
		if err := xml.Unmarshal(body, &res); err != nil {
			return res, err
		}
	}
	return res, nil
}

func sendGet(url string) ([]byte, error) {
	client := &fasthttp.Client{}

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod("GET")

	resp := fasthttp.AcquireResponse()
	if err := client.Do(req, resp); err != nil {
		return []byte{}, err
	} else {
		return resp.Body(), nil

	}
}
