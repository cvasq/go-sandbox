package main

type Message struct {
	Op          string `json:"op"`
	Transaction struct {
		Hash     string  `json:"hash"`
		TotalOut float64 `json:"totalOut,omitempty"`
		Inputs   []struct {
			PrevOut struct {
				Addr    string `json:"addr"`
				N       int    `json:"n"`
				Script  string `json:"script"`
				Spent   bool   `json:"spent"`
				TxIndex int    `json:"tx_index"`
				Type    int    `json:"type"`
				Value   int    `json:"value"`
			} `json:"prev_out"`
			Script   string `json:"script"`
			Sequence int    `json:"sequence"`
		} `json:"inputs"`
		LockTime int `json:"lock_time"`
		Out      []struct {
			Addr    string `json:"addr"`
			N       int    `json:"n"`
			Script  string `json:"script"`
			Spent   bool   `json:"spent"`
			TxIndex int    `json:"tx_index"`
			Type    int    `json:"type"`
			Value   int    `json:"value"`
		} `json:"out"`
		RelayedBy string `json:"relayed_by"`
		Size      int    `json:"size"`
		Time      int    `json:"time"`
		TxIndex   int    `json:"tx_index"`
		Ver       int    `json:"ver"`
		VinSz     int    `json:"vin_sz"`
		VoutSz    int    `json:"vout_sz"`
	} `json:"x"`
}

type SummarizedTransaction struct {
	Hash     string  `json:"hash"`
	TotalOut float64 `json:"totalOut,omitempty"`
	TotalIn  float64 `json:"totalIn,omitempty"`
	MinerFee string  `json:"minerFee,omitempty"`
}
