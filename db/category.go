package db

import (
	"database/sql"
	"fmt"

	"github.com/danilocordeirodev/lamb-core/models"
)

func InsertCategory(c models.Category) (int64, error) {
	fmt.Println("Init InsertCategory")

	err := DbConnect()
	if err != nil {
		return 0, err
	}

	defer Db.Close()

	sqlInsert := "INSERT INTO category (Categ_Name, Categ_Path) VALUES ('" + c.CategName + "','" + c.CategPath + "')"

	var result sql.Result
	result, err = Db.Exec(sqlInsert)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	LastInsertId, err2 := result.LastInsertId()
	if err2 != nil {
		return 0, err2
	}

	fmt.Println("Insert Category > Successfully executed")

	return LastInsertId, nil
}