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

func DeleteCategory(id int) (error) {
	fmt.Println("Init DeleteCategory")

	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	sql := "DELETE FROM category WHERE Categ_id = " + strconv.Itoa(id)
	
	_, err = Db.Exec(sql)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Delete Category > Successfully executed")

	return nil
}

func SelectCategories(CategId int, Slug string) ([] models.Category, error) {
	fmt.Println("Init SelectCategories")

	var Categ []models.Category

	err := DbConnect()
	if err != nil {
		return Categ, err
	}

	defer Db.Close()

	query := "SELECT Categ_Id, Categ_Name, Categ_Path FROM category "

	if CategId > 0 {
		query += "WHERE Categ_Id = " + strconv.Itoa(CategId)
	} else {
		if len(Slug) > 0 {
			query += "WHERE Categ_Path LIKE '%" + Slug +"%'"
		}
	}

	var rows *sql.Rows
	rows, err = Db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return Categ, err
	}
	
	for rows.Next() {
		var c models.Category
		var categId sql.NullInt32
		var categName sql.NullString
		var categPath sql.NullString

		err := rows.Scan(&categId, &categName, &categPath)
		if err != nil {
			return Categ, err
		}

		c.CategID = int(categId.Int32)
		c.CategName = categName.String
		c.CategPath = categPath.String

		Categ = append(Categ, c)
	}

	fmt.Println("Select Category > Successfully executed")
	return Categ, nil
}