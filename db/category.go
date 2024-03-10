package db

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/danilocordeirodev/lamb-core/models"
	"github.com/danilocordeirodev/lamb-core/tools"
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

func UpdateCategory(c models.Category) (error) {
	fmt.Println("Init UpdateCategory")

	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	sql := "UPDATE category SET "

	if (len(c.CategName) > 0) {
		sql += " Categ_Name = '" + tools.RemoveScape(c.CategName) + "'"
	}

	if (len(c.CategPath) > 0) {
		if !strings.HasSuffix(sql, "SET ") {
			sql += ", "
		}
		sql += "Categ_Path = '" + tools.RemoveScape(c.CategPath) + "'"
	}

	sql += " WHERE Categ_Id = " + strconv.Itoa(c.CategID)

	
	_, err = Db.Exec(sql)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update Category > Successfully executed")

	return nil
}