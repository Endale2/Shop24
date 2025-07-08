package controllers

import (
	"net/http"

	"github.com/Endale2/DRPS/admin/models"
	"github.com/Endale2/DRPS/admin/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAdmins retrieves all staff/admin users
func GetAdmins(c *gin.Context) {
	admins, err := repositories.ListAdmins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving admins"})
		return
	}
	c.JSON(http.StatusOK, admins)
}

// GetAdmin retrieves a single admin by ID
func GetAdmin(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}
	admin, err := repositories.GetAdminByID(objID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}
	c.JSON(http.StatusOK, admin)
}

// CreateAdmin creates a new admin
func CreateAdmin(c *gin.Context) {
	var admin models.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	_, err := repositories.CreateAdmin(&admin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating admin"})
		return
	}
	c.JSON(http.StatusOK, admin)
}

// UpdateAdmin updates an existing admin by ID
func UpdateAdmin(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	err = repositories.UpdateAdmin(objID, bson.M(updates))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating admin"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Admin updated"})
}

// DeleteAdmin deletes an admin by ID
func DeleteAdmin(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}
	err = repositories.DeleteAdmin(objID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting admin"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Admin deleted"})
}
