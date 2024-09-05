package handlers

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
)

// User representa a estrutura do usuário no banco de dados
type User struct {
    ID    uint   `json:"id" gorm:"primaryKey"`
    Name  string `json:"name"`
    Email string `json:"email" gorm:"unique"`
}

// CreateUser cria um novo usuário no banco de dados
func CreateUser(c *gin.Context, db *gorm.DB) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	
    // if email already exists, show an error: email already exists
	var existingUser User
	db.Where("email = ?", user.Email).First(&existingUser)
	if existingUser.Email != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}
	
	if err := db.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
        return
    }
    c.JSON(http.StatusOK, user)
}

// GetUser busca um usuário pelo ID
func GetUser(c *gin.Context, db *gorm.DB) {
    var user User
    if err := db.First(&user, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

// UpdateUser atualiza um usuário pelo ID
func UpdateUser(c *gin.Context, db *gorm.DB) {
    var user User
    if err := db.First(&user, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := db.Save(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update user"})
        return
    }
    c.JSON(http.StatusOK, user)
}

// DeleteUser deleta um usuário pelo ID
func DeleteUser(c *gin.Context, db *gorm.DB) {
    var user User
    if err := db.Delete(&user, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
