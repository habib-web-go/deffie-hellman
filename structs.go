package main

type clientHandShake struct {
	P                string `json:"p"`
	G                uint64 `json:"g"`
	CurrentMessageId uint64 `json:"currentMessageId"`
}

type client struct {
	AuthKey          string `json:"authKey"`
	CurrentMessageId uint64 `json:"currentMessageId"`
}
