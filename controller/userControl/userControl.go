package usercontrol

import (
	"fmt"
	"go/token"
	"strings"
	"net/http"

	// "time"

	jwt "github.com/Mzammilishtiaq/gotask1/ulits/jwt"
	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt/v5"
)

type UserController struct{}

func (c *UserController) InitUserControllerRoutes(router *gin.Engine) {
	users := router.Group("/user")
	users.POST("/login", c.Login)
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "Missing Authorization Header"}`))
		return
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	tokenString = string.TrimPrefix(tokenString, "Bearer ")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "Invalid Authorization token"}`))
		return
	}
	fmt.Println(w, "welcome to are authoried")
}

func (c *UserController) Login(ctx *gin.Context) {
	type loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var login loginRequest

	if err := ctx.BindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	email := "ali@gmail.com"
	password := "12345"

	if login.Email == email && login.Password == password {

		tokenString, err := jwtInstance.CreateToken(login.Email)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"email":    login.Email,
			"password": login.Password,
			"token":    tokenString,
		})
		fmt.Println("access given", tokenString)
		return
	}

	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}
