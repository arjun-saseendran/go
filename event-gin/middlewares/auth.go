package middlewares

import (
	"net/http"

	"github.com/arjun-saseendran/event-project-go-gin/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized token!"})
		return
	}
	userId, err := utils.VerifyJWTToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized token!"})
		return
	}
	context.Set("userId",userId)
	context.Next()
}
