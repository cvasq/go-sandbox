// Schedules downtime in Datadog for a single host
// DD_API_KEY and DD_APP_KEY environment variables must be set

package main

import (
	"fmt"
	"gopkg.in/zorkian/go-datadog-api.v2"
	"log"
	"os"
	"time"
)

type ddClient struct {
	client *datadog.Client
}

func NewClient() *ddClient {
	apiKey := os.Getenv("DD_API_KEY")
	appKey := os.Getenv("DD_APP_KEY")
	c := datadog.NewClient(apiKey, appKey)
	if c == nil {
		log.Println("Could not establish Datadog Client")
		os.Exit(1)
	}
	return &ddClient{client: c}
}

func (c ddClient) ScheduleDowntime(scope string, timeInMinutes int) {
	scopeInput := []string{scope}
	start := time.Now().Unix()
	// Convert minutes to seconds
	end := start + int64(timeInMinutes)*60

	downtime, err := c.client.CreateDowntime(&datadog.Downtime{
		Message: datadog.String("Go Script Scheduled Downtime"),
		Scope:   scopeInput,
		End:     datadog.Int(int(end)),
	})
	if err != nil {
		log.Fatal("Could not schedule downtime", err)
	} else {
		log.Println("Scheduled Downtime for:", scope)
		log.Println("Datadog Downtime ID:", downtime.GetId())
	}
}

func main() {
	// Create new datadog client
	ddClient := NewClient()

	// Node to schedule downtime for
	nodeFQDN := "ip-10-240-241-124.us-west-2.compute.internal"

	// 24 hrs - in mins
	downtimeMinutes := 1440

	// Format 'scope' of downtime request. Single host example
	downtimeScope := fmt.Sprintf("%v%v", "host:", nodeFQDN)

	// Schedule downtime
	ddClient.ScheduleDowntime(downtimeScope, downtimeMinutes)
}
