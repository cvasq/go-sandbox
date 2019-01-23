package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func main() {

	url := "wss://ws.blockchain.info/inv"

	log.Printf("Connecting to %v", url)

	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Dial Error:", err)
	}

	// Subscribe to unconfirmed transaction stream on Websocket conn
	opMessage := []byte(`{"op":"unconfirmed_sub"}`)
	err = ws.WriteMessage(websocket.TextMessage, opMessage)
	if err != nil {
		log.Println("Error subscribing to upstream socket:", err)
		ws, _, _ = websocket.DefaultDialer.Dial(url, nil)
		return
	}

	// Read data stream from upstream Websocket connection
	// Receives JSON objects and Decodes to Message struct type
	for {

		// The entire transaction message
		dataIn := &Message{}

		// A subset of the Message struct fields
		dataOut := &SummarizedTransaction{}

		err = ws.ReadJSON(dataIn)

		if err != nil {
			log.Println("Read Error:", err)
			return
		}

		//Calculate outputs total
		getTotalOutputs := func() float64 {
			totalout := 0.0
			for _, i := range dataIn.Transaction.Out {
				totalout += float64(i.Value)
			}
			return totalout / float64(100000000)
		}

		// Calculate inputs total
		getTotalInputs := func() float64 {
			totalout := 0.0
			for _, i := range dataIn.Transaction.Inputs {
				totalout += float64(i.PrevOut.Value)
			}
			return totalout / float64(100000000)
		}

		dataOut.Hash = dataIn.Transaction.Hash
		dataOut.TotalOut = getTotalOutputs()
		dataOut.TotalIn = getTotalInputs()
		dataOut.MinerFee = fmt.Sprintf("%.8f", dataOut.TotalIn-dataOut.TotalOut)

		dataOutBytes, err := json.Marshal(dataOut)
		if err != nil {
			log.Println(err)
			return
		}

		// Print summarized marshalled object
		fmt.Println(string(dataOutBytes))
	}

}
