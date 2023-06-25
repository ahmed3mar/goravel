package middleware

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/exceptions"
)

func CheckUserAuthorized() http.Middleware {
	return func(ctx http.Context) {

		// example to send non-authenticated user exception
		ctx.Request().AbortWithError(&exceptions.NotAuthorized{})

		ctx.Request().Next()
	}
}
