package controllers

import (
	"api-pos/app/models"
	"api-pos/config"
	"api-pos/service"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// LoginResponse token response
type LoginResponse struct {
	ID       uint
	Name     string
	Username string
	Email    string
	Role     string
	Type     string
	// Token        string `json:"token"`
	AccessToken string `json:"access_token"`
	// RefreshToken string `json:"refresh_token"`
}

const (
	name     = "name"
	username = "username"
	email    = "email"
	role     = "role"
)

func Login(c *gin.Context) {
	var u models.LoginCredentials
	if err := c.BindJSON(&u); err != nil {
		panic(err)
	}
	var user models.User
	err := config.DB.Where("email = ?", u.Email).First(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user not found"})
		return
	}

	// cek password
	match := service.CheckPasswordHash(u.Password, user.Password)
	if match == false {
		c.JSON(http.StatusBadRequest, gin.H{"message": "password wrong"})
		return
	}

	// create session login users
	session := sessions.Default(c)
	session.Set(name, user.Name)
	session.Set(username, user.Username)
	session.Set(email, user.Email)
	session.Set(role, user.Role)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	// generate token login
	jwtWrapper := service.JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}
	signedToken, err := jwtWrapper.GenerateToken(user.Email, uint(user.ID))
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"msg": "error signing token",
		})
		c.Abort()
		return
	}

	tokenResponse := LoginResponse{
		ID:          user.ID,
		Name:        user.Name,
		Username:    user.Username,
		Email:       user.Email,
		Role:        user.Role,
		Type:        user.Type,
		AccessToken: signedToken,
	}
	c.JSON(http.StatusOK, gin.H{"message": "login successfully", "data": tokenResponse})
}

func Register(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		panic(err)
	}
	hashedPassword, err := service.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, "encryption failed")
		return
	}
	user.Password = hashedPassword
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "register successfully")
	}
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(name)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete(name)
	session.Delete(username)
	session.Delete(email)
	session.Delete(role)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func GetSession(c *gin.Context) {
	session := sessions.Default(c)
	name := session.Get(name)
	username := session.Get(username)
	email := session.Get(email)
	role := session.Get(role)
	c.JSON(http.StatusOK, gin.H{"name": name, "username": username, "email": email, "role": role})
}

func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}

func GenerateJWT(email, role string, c *gin.Context) (string, error) {
	var mySigningKey = []byte("secretkey")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "Something Went Wrong"})
		return "", err
	}
	return tokenString, nil
}
