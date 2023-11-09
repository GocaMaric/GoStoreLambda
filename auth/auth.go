package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TokenJSON struct {
	Sub       string
	Event_Id  string
	Token_use string
	Scope     string
	Auth_time int
	Iss       string
	Exp       int
	Iat       int
	Client_id string
	Username  string
}

func ValidToken(token string) (bool, error, string) {
	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		fmt.Println("The token is not valid")
		return false, nil, "The token is not valid"
	}

	userInfo, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		fmt.Println("Cannot decode part of the token :", err.Error())
		return false, err, err.Error()
	}

	var tkj TokenJSON
	err = json.Unmarshal(userInfo, &tkj)
	if err != nil {
		fmt.Println("Cannot decode structure of JSON", err.Error())
		return false, err, err.Error()
	}

	now := time.Now()
	tm := time.Unix(int64(tkj.Exp), 0)

	if tm.Before(now) {
		fmt.Println("Token expiration date" + tm.String())
		fmt.Println("Token expired !")
		return false, err, "Token expired !!"
	}
	return true, nil, string(tkj.Username)
}
