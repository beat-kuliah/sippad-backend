package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
	"net/http"
)

var Currencies = map[string]string{
	"USD": "USD",
	"NGN": "NGN",
	"IDR": "IDR",
}

func IsValidCurrency(currency string) bool {
	if _, ok := Currencies[currency]; ok {
		return true
	}
	return false
}

func GetActiveUser(c *gin.Context) (int64, error) {
	value, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to access resource"})
		return 0, fmt.Errorf("Error occured")
	}

	userId, ok := value.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Encountered an issue"})
		return 0, fmt.Errorf("Error occured")
	}
	return userId, nil
}

func HandlerError(err error, c *gin.Context, gValid galidator.Validator) interface{} {
	if c.Request.ContentLength == 0 {
		return "Provide body"
	}

	if e, ok := err.(*json.UnmarshalTypeError); ok {
		if e.Field == "" {
			return "Provide a json body"
		}
		msg := fmt.Sprintf("Invalid for field: '%s'. Expected a value of type '%s'", e.Field, e.Type)
		return msg
	}

	return gValid.DecryptErrors(err)
}
