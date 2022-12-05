package dto

type ResponseStatistic struct {
	Total int `json:"total"`
	Win   int `json:"win"`
	Lose  int `json:"lose"`
	Draw  int `json:"draw"`
}
