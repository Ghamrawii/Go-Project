package middlewares

import (
	"fmt"
	"net/http"

	"example.com/events/utils"
	"github.com/gin-gonic/gin"
)

func Authenticated(context *gin.Context){
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized "})
		return
	}

	userId,err := utils.VaildToken(token)
	if err != nil {
		fmt.Println(err)
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Not authorized "})
		return
	}
	context.Set("userId",userId)
	context.Next()
}