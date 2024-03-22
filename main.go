package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

func main() {
	r := gin.Default()

	//r.Use(AuthMiddleware())

	r.POST("/deneme", func(ctx *gin.Context) {
		var req LoginForm
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	})

	r.GET("/deneme", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "DenemeDenemeDeneme")
	})

	r.Run(":8080")
}

func AuthMiddleware() gin.HandlerFunc {

	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("your-secret-key"), nil
		})

		if err != nil || token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			context.Abort()
			return
		}
		context.Next()
	}

}
