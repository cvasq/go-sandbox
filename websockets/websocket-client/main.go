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

	s, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Dial Error:", err)
	}

	subscribeLiveTransactions := func() {
		opMessage := `{"op":"unconfirmed_sub"}`
		err := s.WriteMessage(websocket.TextMessage, []byte(opMessage))
		if err != nil {
			log.Println("Error subscribing to upstream socket:", err)
			s, _, _ = websocket.DefaultDialer.Dial(url, nil)
			return
		}
	}

	subscribeLiveTransactions()

	// Read data from upstream Websocket
	// Parses JSON and unmarshalls to Message struct type
	for {

		// The entire transaction message
		dataIn := &Message{}

		// A subset of the Message struct fields
		dataOut := &SummarizedTransaction{}

		err = s.ReadJSON(dataIn)

		if err != nil {
			log.Println("read:", err)
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

		fmt.Println(string(dataOutBytes))
	}

}
