package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	ChannelName string `json:"channel_name"`
	UserName    string `json:"user_name"`
}

func HandleRequest(event MyEvent) (string, error) {
	res := fmt.Sprint("test")
	return res, nil
}

func main() {
	lambda.Start(HandleRequest)
}
