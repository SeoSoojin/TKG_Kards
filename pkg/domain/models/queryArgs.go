package models

type QueryArgs struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}
