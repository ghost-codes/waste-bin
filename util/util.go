package util

import (
	"errors"
	"ghost-codes/waste-bin/middleware"

	"github.com/gin-gonic/gin"
)

func GetUserFromCTX(ctx *gin.Context) (user *string, err error) {
	u := ctx.MustGet(middleware.UserPayloadKey).(*string)
	user = u
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if recover() != nil {
			err = errors.New("Unauthorized")
		}
	}()
	return
}
