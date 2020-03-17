package models

import "time"

type LinkResponse struct {
	Link string `json:"link"`
}

type LinkPair [2]string

type LinkPairs []LinkPair

type AllLinks struct {
	Total  int64     `json:"total"`
	Result LinkPairs `json:"result"`
	//Result []Link `json:"result"`
}

type Link struct {
	Id        int64
	Url       string
	Short_url string
	Created   time.Time
	Expires   time.Time
}
