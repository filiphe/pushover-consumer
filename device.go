package main

type Device struct {
	EncOn                  bool   `json:"encryption_enabled"`
	DefSound               string `json:"default_sound"`
	AlwaysDefSound         bool   `json:"always_use_default_sound"`
	DefHighPrioSound       string `json:"default_high_priority_sound"`
	AlwaysDefHighPrioSound bool   `json:"always_use_default_high_priority_sound"`
}
