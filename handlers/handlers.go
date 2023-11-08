package handlers

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

func Handlers(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {

	fmt.Println("Going to process" + path + ">" + method)

	//id := request.PathParameters["id"]
	//idn, _ := strconv.Atoi(id)

	return 400, "Method Invalid"
}
