package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

func main() {
	svc := elbv2.New(session.New())

	input := &elbv2.DescribeLoadBalancersInput{
		LoadBalancerArns: []*string{},
	}

	result, _ := svc.DescribeLoadBalancers(input)

	var list_of_arns []*string
	for _, lb := range result.LoadBalancers {
		list_of_arns = append(list_of_arns, lb.LoadBalancerArn)
		fmt.Println(*lb.LoadBalancerName)
	}
}
