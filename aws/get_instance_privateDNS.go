package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"os"
)

type API struct {
	client *ec2.EC2
}

func NewClient() *API {
	// Load session from shared config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := ec2.New(sess)

	// Return new EC2 client
	return &API{
		client: client,
	}

}

// GetInstancePrivateDNS takes an instance ID and returns the private DNS name
func (a *API) GetInstancePrivateDNS(instanceID string) string {

	params := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceID),
		},
	}

	result, err := a.client.DescribeInstances(params)

	var privateDNSname string

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	} else {
		for _, reservation := range result.Reservations {
			for _, instance := range reservation.Instances {
				privateDNSname = *instance.PrivateDnsName
			}
		}
	}
	return privateDNSname

}

func main() {

	ec2Client := NewClient()
	instanceID := "i-0af910242f1ed8e67"

	privateDNS := ec2Client.GetInstancePrivateDNS(instanceID)

	fmt.Println("Instance ID:", instanceID)
	fmt.Println("Private DNS:", privateDNS)
}
