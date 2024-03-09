package routers

import (
	"encoding/json"
	"strconv"

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