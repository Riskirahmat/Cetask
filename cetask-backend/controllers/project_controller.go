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

func GetProjects(c *gin.Context) {
	projectCollection := db.ProjectCollection

	if projectCollection == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not initialized"})
		return
	}

	userID := c.GetString("userID")

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := projectCollection.Find(ctx, bson.M{"user_id": userObjID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch projects"})
		return
	}
	defer cursor.Close(ctx)

	var projects []models.Project
	for cursor.Next(ctx) {
		var project models.Project
		if err := cursor.Decode(&project); err != nil {
			continue
		}
		projects = append(projects, project)
	}
	c.JSON(http.StatusOK, projects)
}

func GetProjectByID(c *gin.Context) {
	projectCollection := db.ProjectCollection

	if projectCollection == nil {
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

	var project models.Project
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = projectCollection.FindOne(ctx, bson.M{
		"_id":     projectObjID,
		"user_id": userObjID,
	}).Decode(&project)

	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusOK, []models.Project{})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch project"})
		return
	}

	c.JSON(http.StatusOK, project)
}

func CreateProject(c *gin.Context) {
	projectCollection := db.ProjectCollection

	if projectCollection == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not initialized"})
		return
	}

	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userID := c.GetString("userID")

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	project.ID = primitive.NewObjectID()
	project.UserID = userObjID

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = projectCollection.InsertOne(ctx, project)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
		return
	}

	c.JSON(http.StatusCreated, project)
}

func UpdateProject(c *gin.Context) {
	projectCollection := db.ProjectCollection

	if projectCollection == nil {

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

	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	update := bson.M{"$set": bson.M{"name": project.Name}}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = projectCollection.UpdateOne(ctx, bson.M{
		"_id":    projectObjID,
		"user_id": userObjID,
	}, update)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project updated successfully"})
}

func DeleteProject(c *gin.Context) {
    projectID := c.Param("project_id")
    projectObjID, err := primitive.ObjectIDFromHex(projectID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project_id"})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err = db.ColumnCollection.DeleteMany(ctx, bson.M{"project_id": projectObjID})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete columns"})
        return
    }

    _, err = db.TaskCollection.DeleteMany(ctx, bson.M{"project_id": projectObjID})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tasks"})
        return
    }

    _, err = db.ProjectCollection.DeleteOne(ctx, bson.M{"_id": projectObjID})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Project and related columns/tasks deleted"})
}
