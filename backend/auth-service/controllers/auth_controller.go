package controllers

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rahulkale/auth-service/config"
	"github.com/rahulkale/auth-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" // <--- This was the missing line!
	"golang.org/x/crypto/bcrypt"
)

// Register handles user sign-up
func Register(c *gin.Context) {
	var user models.User
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1. Parse JSON Input
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Check if user already exists
	collection := config.DB.Database("auth_db").Collection("users")
	count, err := collection.CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking user"})
		return
	}
	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	// 3. Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}
	user.Password = string(hashedPassword)

	// 4. Set extra fields
	user.ID = primitive.NewObjectID() // This needs the primitive package
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// 5. Insert into Database
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": user})
}

// Login handles user authentication
func Login(c *gin.Context) {
	var input models.User
	var foundUser models.User
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1. Get Email & Password from user
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Find user in Database
	collection := config.DB.Database("auth_db").Collection("users")
	err := collection.FindOne(ctx, bson.M{"email": input.Email}).Decode(&foundUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// 3. Check Password (Compare Hash)
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// 4. Generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": foundUser.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	})

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default_secret"
	}

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// 5. Send Token back
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user":  gin.H{"id": foundUser.ID, "email": foundUser.Email, "name": foundUser.Name},
	})
}
