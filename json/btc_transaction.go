package main

import (
	"encoding/json"
	"fmt"
)

// Data from Blockchain.info Websocket API
// wss://ws.blockchain.info/inv

var raw_data = `{
  "op" : "utx",
  "x" : {
    "lock_time" : 0,
    "ver" : 1,
    "size" : 255,
    "inputs" : [ {
      "sequence" : 4294967295,
      "prev_out" : {
        "spent" : true,
        "tx_index" : 350847445,
        "type" : 0,
        "addr" : "1AQFpyLHdz6FFeXfn7TNQU9RVG3SbT2DyQ",
        "value" : 6393900,
        "n" : 0,
        "script" : "76a914671f13824da4f99d65a84a67b301248e778593a288ac"
      },
      "script" : "483045022100a991b1274fcd7750bdb5ddebda19c8a9c8bfef766d3d9b335ba113aecf4a5d8e022079227d0fd5ad053945a39d48b6c913348d4a44c5724e44beb554d6c8ad3c559a0141042cee9d6b0ba70c6ae9a4f55e083ac5662bca1f21937f2b958bfa16d3e6a2646034d2c72923679f3d3a47418951fe83b7c486f34843f77e5815d71f13dee3ea92"
    } ],
    "time" : 1527463563,
    "tx_index" : 350848289,
    "vin_sz" : 1,
    "hash" : "cae145333e1b99eb3998e4287d1ffa04c587ae100c7296fe3827dacf2dc85b12",
    "vout_sz" : 2,
    "relayed_by" : "0.0.0.0",
    "out" : [ {
      "spent" : false,
      "tx_index" : 350848289,
      "type" : 0,
      "addr" : "19uado8QQtcyb6YrDyxsVfytjR9oywyAip",
      "value" : 6133397,
      "n" : 0,
      "script" : "76a91461b2b67e83b29f1a294490793950190b5f943b4388ac"
    }, {
      "spent" : false,
      "tx_index" : 350848289,
      "type" : 0,
      "addr" : "bc1q0tfs6a27mhg53cm7c3u2z5vlrpqmpnvam4l00j",
      "value" : 259743,
      "n" : 1,
      "script" : "00147ad30d755eddd148e37ec478a1519f1841b0cd9d"
    } ]
  }
}`

type Message struct {
	Op          string `json:"op"`
	Transaction struct {
		Hash   string `json:"hash"`
		Inputs []struct {
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

func main() {
	textBytes := []byte(raw_data)

	var tx Message
	json.Unmarshal(textBytes, &tx)

	fmt.Println("Transaction OP:", tx.Op)
	fmt.Println("Transaction Hash:", tx.Transaction.Hash)
}
