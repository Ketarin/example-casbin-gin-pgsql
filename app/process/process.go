package process

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SomeProcess() gin.HandlerFunc {

	return func(c *gin.Context) {

		responseCode, responseMessage := doSomeProcess(c.GetHeader("AUTH_IDENTITY"))

		c.JSON(responseCode, responseMessage)
	}
}

func doSomeProcess(user string) (int, string) {
	// do stuff
	return 200, fmt.Sprintf("You can do it %s", user)
}
