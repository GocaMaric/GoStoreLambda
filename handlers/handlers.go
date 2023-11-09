package handlers

import (
	"example.com/GoStoreLambda/auth"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"strconv"
)

func Handlers(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {

	fmt.Println("Going to process" + path + ">" + method)

	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)

	isOk, statusCode, user := validateAuthoratization(path, method, headers)
	if !isOk {
		return statusCode, user
	}

	switch path[0:4] {
	case "user":
		return ProcessUsers(body, path, method, user, id, request)

	case "prod":
		return ProcessProducts(body, path, method, user, idn, request)

	case "stoc":
		return ProcessStock(body, path, method, user, idn, request)

	case "addr":
		return ProcessAddress(body, path, method, user, idn, request)

	case "cate":
		return ProcessCategory(body, path, method, user, idn, request)

	case "orde":
		return ProcessOrder(body, path, method, user, idn, request)

	}
	return 400, "Method Invalid"
}

func validateAuthoratization(path string, method string, headers map[string]string) (bool, int, string) {
	if (path == "product" && method == "GET") ||
		(path == "category" && method == "GET") {
		return true, 200, ""
	}
	token := headers["authorization"]
	if len(token) == 0 {
		return false, 401, "Token is required"
	}

	everythingOK, err, msg := auth.ValidToken(token)
	if !everythingOK {
		if err != nil {
			fmt.Println("Token error" + err.Error())
			return false, 401, err.Error()
		} else {
			fmt.Println("Token error" + msg)
			return false, 401, msg
		}
	}

	fmt.Println("Token OK")
	return true, 200, msg

}

func ProcessUsers(body string, path string, method string, user string, id string, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"

}

func ProcessProducts(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"

}

func ProcessCategory(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"

}

func ProcessStock(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"

}
func ProcessAddress(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"

}

func ProcessOrder(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"

}
