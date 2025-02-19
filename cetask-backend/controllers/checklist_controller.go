package controllers

import (
	"cetask-backend/db"
	"cetask-backend/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetChecklists(c *gin.Context) {
	checklistCollection := db.ChecklistCollection
	taskID := c.Param("task_id")
	taskObjID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task_id"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := checklistCollection.Find(ctx, bson.M{"task_id": taskObjID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch checklists"})
		return
	}
	defer cursor.Close(ctx)

	var checklists []models.Checklist
	if err = cursor.All(ctx, &checklists); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding checklists"})
		return
	}

	c.JSON(http.StatusOK, checklists)
}

func CreateChecklist(c *gin.Context) {
	checklistCollection := db.ChecklistCollection
	taskID := c.Param("task_id")
	taskObjID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task_id"})
		return
	}

	var checklist models.Checklist
	if err := c.ShouldBindJSON(&checklist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	checklist.ID = primitive.NewObjectID()
	checklist.TaskID = taskObjID

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = checklistCollection.InsertOne(ctx, checklist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create checklist"})
		return
	}

	c.JSON(http.StatusCreated, checklist)
}

func UpdateChecklist(c *gin.Context) {
	checklistCollection := db.ChecklistCollection

	checklistID := c.Param("checklist_id")
	checklistObjID, err := primitive.ObjectIDFromHex(checklistID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid checklist_id"})
		return
	}

	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{"$set": updateData}
	_, err = checklistCollection.UpdateOne(ctx, bson.M{"_id": checklistObjID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update checklist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Checklist updated successfully"})
}

func DeleteChecklist(c *gin.Context) {
	checklistCollection := db.ChecklistCollection

	checklistID := c.Param("checklist_id")
	checklistObjID, err := primitive.ObjectIDFromHex(checklistID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid checklist_id"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = checklistCollection.DeleteOne(ctx, bson.M{"_id": checklistObjID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete checklist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Checklist deleted successfully"})
}
