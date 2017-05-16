package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
}


type APIActCardResultWrapper struct {
	Activities []*APIActivityCard `json:"activities,omitempty"`
	FailedIDs  []int              `json:"failed_ids,omitempty"`
}

type APIActivityCard struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	SubTitle    string  `json:"subtitle"`
	URLSeo      string  `json:"url_seo"`
	CityID      int     `json:"city_id"`
	CityName    string  `json:"city_name"`
	Instance    int     `json:"instance"`
	ImageURL    string  `json:"image_url"`
	VideoURL    string  `json:"video_url"`
	Participate int     `json:"participate"`
	Score       float64 `json:"score"`
	MarketPrice string  `json:"market_price"`
	SellPrice   string  `json:"sell_price"`
	StartTime   string  `json:"start_time"`
	SoldOut     bool    `json:"sold_out"`
}

//GError 错误信息，succes为true时，值都为空
type GError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

//KlookAPIJSONFormat 返回的常用格式
type KlookAPIJSONFormat struct {
	GError  `json:"error,omitempty"`
	XStruct interface{} `json:"result"`
	Success bool        `json:"success"`
}





func main() {
	var s map[string]interface{}
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
	fmt.Println(s["servers"])
	fmt.Println(s["servers"])

	bs := []byte = [[120 141 127 119 765 523 23 1659 835 0 357]  "GPS":  , ]
	fmt.Println(string(bs))


	var apiActCardsWrapper *APIActCardResultWrapper
	var rspStruct = &KlookAPIJSONFormat{XStruct: &apiActCardsWrapper}

	json.Unmarshal(bs, &rspStruct)




}