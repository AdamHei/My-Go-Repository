package models

type ITicker interface {
	GetExchangeData() LimitedJson
}

type LimitedJson map[string]map[string]string
