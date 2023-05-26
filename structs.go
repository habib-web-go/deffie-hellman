package main

type clientHandShake struct {
	P                uint64 `json:"p"`
	G                uint64 `json:"g"`
	CurrentMessageId uint64 `json:"currentMessageId"`
}

type client struct {
	AuthKey          uint64 `json:"authKey"`
	CurrentMessageId uint64 `json:"currentMessageId"`
}
