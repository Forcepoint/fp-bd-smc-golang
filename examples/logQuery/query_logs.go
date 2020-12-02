package main

import (
	"encoding/json"
	"fmt"
	"github.cicd.cloud.fpdev.io/BD/fp-smc-golang/src/smc/monitoring"
	"github.cicd.cloud.fpdev.io/BD/fp-smc-golang/src/smc/monitoring/constants"
	"github.cicd.cloud.fpdev.io/BD/fp-smc-golang/src/smc/monitoring/logsocket"
	"github.com/sirupsen/logrus"
	"time"
)

type Request struct {
	Query  Query  `json:"query,omitempty"`
	Fetch  Fetch  `json:"fetch,omitempty"`
	Format Format `json:"format,omitempty"`
}

type Filter struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type Query struct {
	StartMs int16  `json:"start_ms,omitempty"`
	EndMs   int16  `json:"end_ms,omitempty"`
	Type    string `json:"type,omitempty"`
	Filter  Filter `json:"filter,omitempty"`
}

type Fetch struct {
	Quantity  int8 `json:"quantity,omitempty"`
	Backwards bool `json:"backwards,omitempty"`
}

type Format struct {
	Type        string               `json:"type,omitempty"`
	FieldFormat string               `json:"field_format,omitempty"`
	FieldIds    []constants.LogField `json:"field_ids,omitempty"`
}

func main() {
	socket := logsocket.New("ws://178.32.156.143:3080/6.6/monitoring/log/socket")

	defer socket.Close()

	socket.RequestHeader.Add("Cookie", "JSESSIONID=E8F2FE537287FA28389503AE4344F2EF")

	socket.OnTextMessage = func(message string, socket logsocket.Socket) {
		fmt.Println(message)
	}

	socket.Connect()

	r := Request{
		Query: Query{
			StartMs: 0,
			EndMs:   0,
			Type:    "stored",
			Filter: Filter{
				Type:  "translated",
				Value: "default_false($Action == 9)",
			},
		},
		Fetch: Fetch{
			Quantity:  100,
			Backwards: true,
		},
		Format: Format{
			Type:        "texts",
			FieldFormat: "pretty",
			FieldIds:    monitoring.DefaultLogFields,
		},
	}

	data, err := json.Marshal(r)

	if err != nil {
		logrus.Error(err)
	}

	socket.SendText(string(data))

	// Just for the example to wait until all data has been received from the SMC
	time.Sleep(10 * time.Second)
}
