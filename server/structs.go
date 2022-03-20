package main

type Error struct {
	Message string `json:"message"`
}

type Result struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Request struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Clock []struct {
		ClockId  string `json:"clock_id"`
		ClockNum int    `json:"clock_num"`
	} `json:"clock"`
}
