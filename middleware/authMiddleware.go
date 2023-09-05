package middleware

import (
	"fmt"
	db "ghost-codes/waste-bin/db/models"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

const (
	authorizationKey               = "Authorization"
	authorizationBearerType        = "bearer"
	UserPayloadKey                 = "user"
	DriverPayloadKey               = "driver"
	UserTypeClient          string = "client"
	UserTypeKey                    = "userType"
)

func UserAuthMiddleware(store db.Store, firebaseAuth *auth.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationKey)
		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, newErrorJson(fmt.Errorf("Unauthorized")))
			return
		}
		fields := strings.Fields(authorizationHeader)

		if len(fields) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, newErrorJson(fmt.Errorf("Unauthorized")))
			return
		}

		if strings.ToLower(fields[0]) != authorizationBearerType {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, newErrorJson(fmt.Errorf("Unauthorized")))
			return
		}

		token, err := firebaseAuth.VerifyIDToken(ctx, fields[1])

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, newErrorJson(err))
			return
		}

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, newErrorJson(err))
			return
		}

		ctx.Set(UserPayloadKey, token.UID)

		ctx.Next()
	}

}

func newErrorJson(err error) gin.H {
	return gin.H{
		"message": err.Error(),
	}
}
