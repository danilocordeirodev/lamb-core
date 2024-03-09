package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/danilocordeirodev/lamb-core/models"
	"github.com/danilocordeirodev/lamb-core/secretmanager"
)



var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretmanager.GetSecret(os.Getenv(("SecretName")))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Database sucessfully connected")
	return nil
}

func ConnStr(keys models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = keys.Username
	authToken = keys.Password
	dbEndpoint = keys.Host
	dbName = "lamb"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dsn)
	return dsn
}

func UserIsAdmin(userUUID string)(bool, string) {
	fmt.Println("Initiate UserIsAdmin")

	err := DbConnect()

	if err != nil {
		return false, err.Error()
	}

	defer Db.Close()

	sql := "SELECT 1 FROM users WHERE User_UUID='"+userUUID+"' AND UserStatus = 0"

	rows, err := Db.Query(sql)
	if err != nil {
		return false, err.Error()
	}

	var value string
	rows.Next()
	rows.Scan(&value)

	fmt.Println("UserIsAdmin > Sucessfully executed - value: " + value)

	if value == "1" {
		return true, ""
	}

	return false, "User is not admin"
}