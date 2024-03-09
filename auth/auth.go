package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TokenJSON struct {
	Sub      string
	EventID  string `json:"event_id"`
	TokenUse string `json:"token_use"`
	Scope    string
	AuthTime int `json:"auth_time"`
	Iss      string
	Exp      int
	Iat      int
	ClientID string `json:"client_id"`
	Username string `json:"cognito:username"`
}

func TokenValidation(token string) (bool, string, error) {
	parts := strings.Split(token, ".")

	if len(parts) < 2 {
		fmt.Println("Invalid token format")
		return false, "", fmt.Errorf("invalid token format")
	}

	userInfo, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		fmt.Println("Failed to decode token: ", err.Error())
		return false, "", fmt.Errorf("failed to decode token: %s", err.Error())
	}

	var tokenJ TokenJSON
	err = json.Unmarshal(userInfo, &tokenJ)
	if err != nil {
		fmt.Println("Error to decode JSON structure", err.Error())
		return false, "", fmt.Errorf("failed to unmarshal JSON: %s", err.Error())
	}

	now := time.Now()
    tm := time.Unix(int64(tokenJ.Exp), 0)
 
    if tm.Before(now) {
		fmt.Println("Token has expired")
        return false, "", fmt.Errorf("token has expired")
    }
 
    return true, tokenJ.Username, nil
}
