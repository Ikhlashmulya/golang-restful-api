package model

type WebResponse struct {
	Code   int    `json:"code,omitempty"`
	Status string `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
	Data   any    `json:"data,omitempty"`
}
