package middleware

// import (
// 	exceptions "go-chi-api/exception"
// 	helpers "go-chi-api/helper"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func Authenthication() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		verifyToken, err := helpers.VerifyToken(ctx)
// 		_ = verifyToken

// 		if err != nil {
// 			exceptions.Errors(ctx, http.StatusUnauthorized, "Unauthenthicated", "Unauthenthicated")
// 			return
// 		}
// 		ctx.Set("userData", verifyToken)
// 		ctx.Next()

// 	}
// }
