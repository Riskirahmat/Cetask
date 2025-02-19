package routes

import (
	"cetask-backend/controllers"
	"cetask-backend/middleware"
	"github.com/gin-gonic/gin"
)

func ProjectRoutes(router *gin.Engine) {
	projectGroup := router.Group("/projects")
	projectGroup.Use(middleware.AuthMiddleware()) 
	{
		// CRUD Project
		projectGroup.GET("/", controllers.GetProjects)
		projectGroup.POST("/", controllers.CreateProject)
		projectGroup.GET("/:project_id", controllers.GetProjectByID)
		projectGroup.PUT("/:project_id", controllers.UpdateProject)
		projectGroup.DELETE("/:project_id", controllers.DeleteProject)

		// CRUD Columns
		projectGroup.GET("/:project_id/columns", controllers.GetColumns)
		projectGroup.GET("/:project_id/columns/:column_id", controllers.GetColumnByID)
		projectGroup.POST("/:project_id/columns", controllers.CreateColumn)
		projectGroup.PUT("/:project_id/columns/:column_id", controllers.UpdateColumn)
		projectGroup.PUT("/:project_id/columns/:column_id/position", controllers.UpdateColumnPosition)
		projectGroup.DELETE("/:project_id/columns/:column_id", controllers.DeleteColumn)

		// CRUD Tasks
		projectGroup.GET("/:project_id/columns/:column_id/tasks", controllers.GetTasks)
		projectGroup.GET("/:project_id/columns/:column_id/tasks/:task_id", controllers.GetTaskByID)
		projectGroup.POST("/:project_id/columns/:column_id/tasks", controllers.CreateTask)
		projectGroup.PUT("/:project_id/columns/:column_id/tasks/:task_id", controllers.UpdateTask)
		projectGroup.PUT("/:project_id/columns/:column_id/tasks/:task_id/move", controllers.MoveTask)
		projectGroup.PUT("/:project_id/columns/:column_id/tasks/:task_id/status", controllers.UpdateTaskStatus)
		projectGroup.PUT("/:project_id/columns/:column_id/tasks/:task_id/position", controllers.UpdateTaskPosition)
		projectGroup.DELETE("/:project_id/columns/:column_id/tasks/:task_id", controllers.DeleteTask)

		// CRUD Checklists
		projectGroup.GET("/:project_id/columns/:column_id/tasks/:task_id/checklists", controllers.GetChecklists)
		projectGroup.POST("/:project_id/columns/:column_id/tasks/:task_id/checklists", controllers.CreateChecklist)
		projectGroup.PUT("/:project_id/columns/:column_id/tasks/:task_id/checklists/:checklist_id", controllers.UpdateChecklist)
		projectGroup.DELETE("/:project_id/columns/:column_id/tasks/:task_id/checklists/:checklist_id", controllers.DeleteChecklist)
	}
}
