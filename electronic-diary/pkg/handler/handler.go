package handler

import (
	"github.com/biyoba1/electronic-diary/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

//@student
//student
//getall
//getbyid/:id
//update/:id
//delete/:id

//@subject
//subject
//getall
//getbyid:/id
//update:/id
//delete:

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	student := router.Group("/student")
	{
		student.POST("/", h.CreateStudent)
		student.GET("/", h.GetAllStudent)
		student.GET("/:id", h.GetStudentById)
		student.PUT("/:id", h.UpdateStudent)
		student.DELETE("/:id", h.DeleteStudent)
	}
	subject := router.Group("/subject")
	{
		subject.POST("/", h.CreateSubject)
		subject.GET("/", h.GetAllSubject)
		subject.PUT("/:id", h.UpdateSubject)
		subject.DELETE("/:id", h.DeleteSubject)
	}
	return router
}
