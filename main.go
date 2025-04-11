package main

import (
	// "naverdictionary/rest"
	"naverdictionary/telegram" // Import the telegram package

	"github.com/aws/aws-lambda-go/lambda"
)

// func main() {
// 	// Start the server
// 	rest.StartServer()
// }

// Start sets up the Lambda handler
func main() {
	lambda.Start(telegram.LambdaHandler)
}
