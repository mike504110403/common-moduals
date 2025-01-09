package countrylist

import "encoding/xml"

type (
	CountyItems struct {
		XMLName    xml.Name     `xml:"countyItems"`
		CountyLIst []CountyItem `xml:"countyItem"`
	}

	CountyItem struct {
		CountyCode   string `xml:"countycode"`
		CountyName   string `xml:"countyname"`
		CountyCode01 string `xml:"countycode01"`
	}

	TownItems struct {
		XMLName  xml.Name   `xml:"townItems"`
		TownList []TownItem `xml:"townItem"`
	}

	TownItem struct {
		TownCode string `xml:"towncode"`
		TownName string `xml:"townname"`
	}

	CountryTownSelectItem struct {
		Country string
		Code    string
		Towns   []TownSelectItem
	}
	TownSelectItem struct {
		Code string `json:"town_code"`
		Town string `json:"town_name"`
	}
)
