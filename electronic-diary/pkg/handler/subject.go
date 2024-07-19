package handler

import (
	diary "github.com/biyoba1/electronic-diary"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateSubject(c *gin.Context) {
	var input diary.SubjectDiary
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid type body")
		return
	}
	id, err := h.services.Subject.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

type getAllSubjectResponse struct {
	Students []diary.StudentDiary `json:"students"`
	Subjects []diary.SubjectDiary `json:"subjects"`
}

func (h *Handler) GetAllSubject(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	students, subjects := h.services.Subject.GetAll(id)
	response := getAllSubjectResponse{
		Students: students,
		Subjects: subjects,
	}
	c.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateSubject(c *gin.Context) {
	subjectId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	var input diary.UpdateSubjectInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid type body")
		return
	}
	if err := h.services.Subject.Update(subjectId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) DeleteSubject(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	err = h.services.Subject.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
