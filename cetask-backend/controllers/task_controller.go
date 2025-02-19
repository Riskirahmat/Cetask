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

func GetTasks(c *gin.Context) {
	taskCollection := db.TaskCollection
	checklistCollection := db.ChecklistCollection

	if taskCollection == nil || checklistCollection == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not initialized"})
		return
	}

	columnID := c.Param("column_id")
	columnObjID, err := primitive.ObjectIDFromHex(columnID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid column_id"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := taskCollection.Find(ctx, bson.M{"column_id": columnObjID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	defer cursor.Close(ctx)

	var tasksWithChecklist []gin.H
	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			continue
		}

		var checklists []models.Checklist
		checklistCursor, err := checklistCollection.Find(ctx, bson.M{"task_id": task.ID})
		if err == nil {
			checklistCursor.All(ctx, &checklists)
		}

		tasksWithChecklist = append(tasksWithChecklist, gin.H{
			"id":        task.ID.Hex(),
			"title":     task.Title,
			"desc":      task.Desc,
			"column_id": task.ColumnID.Hex(),
			"user_id":   task.UserID.Hex(),
			"checklist": checklists,
		})
	}

	c.JSON(http.StatusOK, tasksWithChecklist)
}

func GetTaskByID(c *gin.Context) {
	taskCollection := db.TaskCollection

	if taskCollection == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not initialized"})
		return
	}

	taskID := c.Param("task_id")

	taskObjID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task_id"})
		return
	}

	var task models.Task
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = taskCollection.FindOne(ctx, bson.M{"_id": taskObjID}).Decode(&task)

	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusOK, []models.Task{})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
    taskCollection := db.TaskCollection

    if taskCollection == nil {

        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not initialized"})
        return
    }

    var input struct {
        Title  string `json:"title" binding:"required"`
        Desc   string `json:"desc"`
        Status string `json:"status"`
    }

    columnID := c.Param("column_id")
    userID := c.GetString("userID")

    columnObjID, err := primitive.ObjectIDFromHex(columnID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid column_id"})
        return
    }

    userObjID, err := primitive.ObjectIDFromHex(userID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
        return
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    if input.Status == "" {
        input.Status = "TO DO"
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    taskCount, err := taskCollection.CountDocuments(ctx, bson.M{"column_id": columnObjID})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to determine task position"})
        return
    }

    newTask := models.Task{
        ID:       primitive.NewObjectID(),
        Title:    input.Title,
        Desc:     input.Desc,
        ColumnID: columnObjID,
        UserID:   userObjID,
        Status:   input.Status,
        Position: int(taskCount),
    }

    _, err = taskCollection.InsertOne(ctx, newTask)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
        return
    }

    c.JSON(http.StatusCreated, newTask)
}

func UpdateTask(c *gin.Context) {
	taskCollection := db.TaskCollection

	if taskCollection == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not initialized"})
		return
	}

	taskID := c.Param("task_id")

	taskObjID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task_id"})
		return
	}

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	update := bson.M{"$set": bson.M{"title": task.Title, "desc": task.Desc}}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = taskCollection.UpdateOne(ctx, bson.M{"_id": taskObjID}, update)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func UpdateTaskStatus(c *gin.Context) {
    taskID := c.Param("task_id")
    taskObjID, err := primitive.ObjectIDFromHex(taskID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task_id"})
        return
    }

    var input struct {
        Status string `json:"status"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    update := bson.M{"$set": bson.M{"status": input.Status}}
    _, err = db.TaskCollection.UpdateOne(ctx, bson.M{"_id": taskObjID}, update)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task status"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Task status updated"})
}

func UpdateTaskPosition(c *gin.Context) {
    taskID := c.Param("task_id")
    columnID := c.Param("column_id")

    taskObjID, err := primitive.ObjectIDFromHex(taskID)
    columnObjID, err2 := primitive.ObjectIDFromHex(columnID)

    if err != nil || err2 != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task_id or column_id"})
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
    _, err = db.TaskCollection.UpdateOne(ctx, bson.M{"_id": taskObjID, "column_id": columnObjID}, update)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task position"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Task position updated"})
}


func DeleteTask(c *gin.Context) {
	taskID := c.Param("task_id")
	taskObjID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task_id"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()


	_, err = db.ChecklistCollection.DeleteMany(ctx, bson.M{"task_id": taskObjID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete related checklists"})
		return
	}

	_, err = db.TaskCollection.DeleteOne(ctx, bson.M{"_id": taskObjID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task and related checklists deleted"})
}


func MoveTask(c *gin.Context) {
    taskCollection := db.TaskCollection

    if taskCollection == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not initialized"})
        return
    }

    var input struct {
        ColumnID string `json:"column_id"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    taskID := c.Param("task_id")

    taskObjID, err := primitive.ObjectIDFromHex(taskID)
    newColumnObjID, err2 := primitive.ObjectIDFromHex(input.ColumnID)
    if err != nil || err2 != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task_id or column_id"})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    taskCount, err := taskCollection.CountDocuments(ctx, bson.M{"column_id": newColumnObjID})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to determine new task position"})
        return
    }

    newPosition := int(taskCount)

    update := bson.M{"$set": bson.M{"column_id": newColumnObjID, "position": newPosition}}

    _, err = taskCollection.UpdateOne(ctx, bson.M{"_id": taskObjID}, update)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to move task"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Task moved successfully", "new_position": newPosition})
}
