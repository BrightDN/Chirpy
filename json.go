package main

type params struct {
	Body string `json:"body"`
}

type jsonResp struct {
	Error        string `json:"error,omitempty"`
	Cleaned_Body string `json:"cleaned_body,omitempty"`
	Valid        bool   `json:"valid,omitempty"`
}
