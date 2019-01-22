package auth

import (
	"../../database/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"

	//jwt "github.com/dgrijalva/jwt-go"
)


func hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateToken(c *gin.Context, userId uint) string {
	db := c.MustGet("db").(*gorm.DB)
	u1 := uuid.Must(uuid.NewV4())

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"token": u1,
	})

	token := models.Token{
		UserID:userId,
		Token:u1.String(),
		Expiry:time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	db.Create(&token)
	tokenString, _ := jwtToken.SignedString([]byte("your-256-bit-secret"))
	return tokenString

}

func register(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)

	type RequestBody struct {
		Username string `json:"username" binding:"required"`
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
		Password string `json:"password" binding:"required"`
	}

	var body RequestBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{
		FirstName: body.FirstName,
		LastName:body.LastName,
		Username:body.Username,
	}

	password, err := hash(body.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Password = password

	var count int
	db.Model(&models.User{}).Where("username = ?", user.Username).Count(&count)
	if count != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User Registered"})
}


func login(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.Where("username = ?", body.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username not found!"})
		return
	}

	if !checkHash(body.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username & Password not matched!"})
		return
	}

	token := generateToken(c, user.ID)
	c.JSON(http.StatusOK, gin.H{"token" : token, "user": user.Serialize()})
}
