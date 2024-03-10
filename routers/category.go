package routers

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/danilocordeirodev/lamb-core/db"
	"github.com/danilocordeirodev/lamb-core/models"
)

func InsertCategory(body string , User string) (int, string) {
	var t models.Category

	err := json.Unmarshal([]byte(body), &t)

	if err != nil { 
		return 400, "Error in received data - category: " + err.Error()
	}

	if len(t.CategName) == 0 {
		return 400, "Name of category not informed"
	}

	if len(t.CategPath) == 0 {
		return 400, "Path of category not informed"
	}

	isAdmin, msg := db.UserIsAdmin(User)

	if !isAdmin {
		return 400, msg
	}

	result, errInsert := db.InsertCategory(t)
	if errInsert != nil {
		return 400, "Error to try insert category " + t.CategName + " > " + errInsert.Error()
	}

	return 200, "{ CategId: " + strconv.Itoa(int(result)) + "}"
}

func UpdateCategory(body string , User string, id int) (int, string) {
	var t models.Category

	err := json.Unmarshal([]byte(body), &t)

	if err != nil { 
		return 400, "Error in received data - category: " + err.Error()
	}

	if len(t.CategName) == 0 && len(t.CategPath) == 0 {
		return 400, "Name and Path of category not informed"
	}

	isAdmin, msg := db.UserIsAdmin(User)

	if !isAdmin {
		return 400, msg
	}


	t.CategID = id
	errUpdate := db.UpdateCategory(t)
	if errUpdate != nil {
		return 400, "Error to try update category " + strconv.Itoa(id) + " > " + errUpdate.Error()
	}

	return 200, "Update Ok"
}

func DeleteCategory(body string , User string, id int) (int, string) {
	
	if id == 0 { 
		return 400, "Category id not valid"
	}

	isAdmin, msg := db.UserIsAdmin(User)

	if !isAdmin {
		return 400, msg
	}


	err := db.DeleteCategory(id)
	if err != nil {
		return 400, "Error to try delete category " + strconv.Itoa(id) + " > " + err.Error()
	}

	return 204, "Delete Ok"
}

func SelectCategories(body string, request events.APIGatewayV2HTTPRequest) (int, string) {
	var err error
	var CategId int
	var Slug string

	if len(request.QueryStringParameters["categId"]) > 0 {
		CategId, err = strconv.Atoi(request.QueryStringParameters["categId"])
		if err != nil {
			return 500, "Failed to convert to int -> " + request.QueryStringParameters["categId"]
		}
	} else {
		if len(request.QueryStringParameters["slug"]) > 0 {
			Slug = request.QueryStringParameters["slug"]
		}
	}

	list, err2 := db.SelectCategories(CategId, Slug)
	if err2 != nil {
		return 400, "Failed to select categories > " + err2.Error()
	}

	Categ, err3 := json.Marshal(list)
	if err3 != nil {
		return 400, "Failed to convert to JSON > " + err3.Error()
	}

	return 200, string(Categ)
}