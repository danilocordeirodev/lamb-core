package handlers

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/danilocordeirodev/lamb-core/auth"
	"github.com/danilocordeirodev/lamb-core/routers"
)


func Handlers(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Processing "+path+" > "+method)

	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)

	isOk, statusCode, user := validAuth(path, method, headers)
	if !isOk {
		return statusCode, user
	}

	switch path[0:4] {
	case "user":
		return ProcessUsers(body, path, method, user, id, request)
	case "prod":
		return ProcessProducts(body, path, method, user, idn, request)
	case "stoc":
		return ProcessStocks(body, path, method, user, idn, request)
	case "addr":
		return ProcessAddresses(body, path, method, user, idn, request)
	case "cate":
		return ProcessCategories(body, path, method, user, idn, request)
	case "orde":
		return ProcessOrders(body, path, method, user, idn, request)
	}

	return 400, "Method Invalid" + path[0:4]
}

func validAuth(path string, method string, headers map[string]string)(bool, int, string) {
	if (path == "product" && method == "GET") ||
		(path == "category" && method == "GET") {
			return true, 200, ""
		}

	token := headers["authorization"]
	if len(token) == 0 {
		return false, 401, "Token required"
	}

	ok, msg, err := auth.TokenValidation(token)
	if !ok {
		if err != nil {
			fmt.Println("Token error: " + err.Error())
			return false, 401, err.Error()
		} else {
			fmt.Println("Token error msg: " + msg)
			return false, 401, msg
		}
	} 
	
	fmt.Println("Token Ok")
	return true, 200, msg

}

func ProcessUsers(body string, path string, method string, user string, id string, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"
}

func ProcessProducts(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"
}

func ProcessCategories(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "POST":
		return routers.InsertCategory(body, user)
	case "PUT":
		return routers.UpdateCategory(body, user, id)
	}
	
	return 400, "Method Invalid"
}

func ProcessStocks(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"
}

func ProcessAddresses(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"
}

func ProcessOrders(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"
}