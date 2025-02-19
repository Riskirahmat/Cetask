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
	"go.mongodb.org/mongo-driver/mongo"
)

func GetColumns(c *gin.Context) {
	columnCollection := db.ColumnCollection
	projectCollection := db.ProjectCollection

	if columnCollection == nil || projectCollection == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not initialized"})
		return
	}

	projectID := c.Param("project_id")
	userID := c.GetString("userID")

	projectObjID, err := primitive.ObjectIDFromHex(projectID)
	userObjID, err2 := primitive.ObjectIDFromHex(userID)

	if err != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var project models.Project
	err = projectCollection.FindOne(ctx, bson.M{
		"_id":     projectObjID,
		"user_id": userObjID,
	}).Decode(&project)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access to project"})
		return
	}

	cursor, err := columnCollection.Find(ctx, bson.M{"project_id": projectObjID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch columns"})
		return
	}
	defer cursor.Close(ctx)

	var columns []models.Column
	for cursor.Next(ctx) {
		var column models.Column
		if err := cursor.Decode(&column); err != nil {
			continue
		}
		columns = append(columns, column)
	}

	if len(columns) == 0 {
		c.JSON(http.StatusOK, []models.Column{}) 
		return
	}

	c.JSON(http.StatusOK, columns)
}

func GetColumnByID(c *gin.Context) {
	columnCollection := db.ColumnCollection

	if columnCollection == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not initialized"})
		return
	}

	projectID := c.Param("project_id")
	columnID := c.Param("column_id")

	projectObjID, err := primitive.ObjectIDFromHex(projectID)
	columnObjID, err2 := primitive.ObjectIDFromHex(columnID)

	if err != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var column models.Column
	err = columnCollection.FindOne(ctx, bson.M{
		"_id":        columnObjID,
		"project_id": projectObjID,
	}).Decode(&column)

	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusOK, []models.Column{}) 
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch column"})
		return
	}

	c.JSON(http.StatusOK, column)
}

func CreateColumn(c *gin.Context) {
	columnCollection := db.ColumnCollection

	if columnCollection == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not initialized"})
		return
	}

	var column models.Column
	projectID := c.Param("project_id")
	userID := c.GetString("userID")

	projectObjID, err := primitive.ObjectIDFromHex(projectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project_id"})
		return
	}

	if err := c.ShouldBindJSON(&column); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	column.ID = primitive.NewObjectID()
	column.ProjectID = projectObjID
	column.UserID, _ = primitive.ObjectIDFromHex(userID)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = columnCollection.InsertOne(ctx, column)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create column"})
		return
	}

	c.JSON(http.StatusCreated, column)
}

func UpdateColumn(c *gin.Context) {
	columnCollection := db.ColumnCollection

	if columnCollection == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not initialized"})
		return
	}

	columnID := c.Param("column_id")
	projectID := c.Param("project_id")

	columnObjID, err := primitive.ObjectIDFromHex(columnID)
	projectObjID, err2 := primitive.ObjectIDFromHex(projectID)

	if err != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var column models.Column
	if err := c.ShouldBindJSON(&column); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	update := bson.M{"$set": bson.M{"name": column.Name}}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = columnCollection.UpdateOne(ctx, bson.M{
		"_id":        columnObjID,
		"project_id": projectObjID,
	}, update)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update column"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Column updated successfully"})
}

func UpdateColumnPosition(c *gin.Context) {
    columnID := c.Param("column_id")
    projectID := c.Param("project_id")

    columnObjID, err := primitive.ObjectIDFromHex(columnID)
    projectObjID, err2 := primitive.ObjectIDFromHex(projectID)

    if err != nil || err2 != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid column_id or project_id"})
        return
    }

    var input struct {
        NewPosition int `json:"position"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    update := bson.M{"$set": bson.M{"position": input.NewPosition}}
    _, err = db.ColumnCollection.UpdateOne(ctx, bson.M{"_id": columnObjID, "project_id": projectObjID}, update)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update column position"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Column position updated"})
}


func DeleteColumn(c *gin.Context) {
    columnID := c.Param("column_id")
    columnObjID, err := primitive.ObjectIDFromHex(columnID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid column_id"})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err = db.TaskCollection.DeleteMany(ctx, bson.M{"column_id": columnObjID})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tasks in column"})
        return
    }

    _, err = db.ColumnCollection.DeleteOne(ctx, bson.M{"_id": columnObjID})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete column"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Column and its tasks deleted successfully"})
}
