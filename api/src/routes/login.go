package routes

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"encoding/json"
	. "fmt"
)

type Account struct {
	UserID   string `json:"userId"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)

	if err != nil {
		Println(err.Error())
	}

	var account Account
	if err := json.Unmarshal([]byte(string(value)), &account); err != nil {
		Println(err.Error())
	}

	result := false
	if account.UserID == "userid" && account.Password == "password" {
		result = true
	}

	c.JSON(200, gin.H{ "result": result })
}