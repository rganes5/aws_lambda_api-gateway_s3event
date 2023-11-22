package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Handler)

}

type InputEvent struct {
	Link string `json:"link"`
	Key  string `json:"key"`
}

func Handler(event InputEvent) {
	fmt.Println("Function invoked!")
	log.Println("Yay")
}
