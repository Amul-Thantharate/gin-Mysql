package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `gorm:"unique" json:"password"`
}

func UserRoutes(r *gin.Engine, db *gorm.DB) {

	r.POST("/users", func(c *gin.Context) {
		CreateUser(c, db)
	})
	r.GET("/users", func(c *gin.Context) {
		GetUsers(c, db)
	})
	r.GET("/users/:id", func(c *gin.Context) {
		GetUserByID(c, db)
	})
	r.PUT("/users/:id", func(c *gin.Context) {
		UpdateUser(c, db)
	})
	r.DELETE("/users/:id", func(c *gin.Context) {
		DeleteUser(c, db)
	})
}

func CreateUser(c *gin.Context, db *gorm.DB) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	db.Create(&user)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "✅ User created successfully!", "user": user})
}

func GetUsers(c *gin.Context, db *gorm.DB) {
	var user []User
	db.Find(&user)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "📋 List of users", "users": user})
}

func GetUserByID(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var user User
	if err := db.First(&user, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Error": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "👤 User found", "user": user})
}

func UpdateUser(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var user User
	if err := db.First(&user, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Error": err.Error()})
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "⚠️" + err.Error()})
		return
	}
	db.First(&user)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "🔄 User updated successfully!", "user": user})
}

func DeleteUser(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	if err := db.Delete(&User{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "❌ User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "🗑️ User deleted successfully!"})
}
