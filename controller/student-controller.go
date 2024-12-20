package controller

import (
	"log"
	"net/http"
	"strconv"
	"student-teacher-api/Request"
	"student-teacher-api/manager"

	"github.com/labstack/echo/v4"
)

// StudentController handles HTTP requests related to Students
type StudentController struct {
	Manager *manager.StudentManager
}

// GetStudents handler to fetch all Students
func (c *StudentController) GetStudents(ctx echo.Context) error {
	flagValue := ctx.QueryParam("flag")
	flag, err := strconv.ParseBool(flagValue)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid flag value. Accepted true or false.",
		})
	}
	studentResponses, err := c.Manager.GetStudents(flag)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, studentResponses)
}

// GetStudentByID handler to fetch a Student by ID
func (c *StudentController) GetStudentByID(ctx echo.Context) error {
	flagValue := ctx.QueryParam("flag")
	flag, err := strconv.ParseBool(flagValue)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid flag value. Accepted true or false.",
		})
	}
	id := ctx.Param("id")

	student, err := c.Manager.GetStudentByID(flag, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	log.Println("Returned student by ID")
	return ctx.JSON(http.StatusOK, student)
}

// CreateStudent handler to create a new Student
func (c *StudentController) CreateStudent(ctx echo.Context) error {
	flagValue := ctx.QueryParam("flag")
	flag, err := strconv.ParseBool(flagValue)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid flag value. Accepted values are true or false.",
		})
	}
	var req Request.StudentRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	result, err := c.Manager.CreateStudent(flag, req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	log.Println("Created new student")
	return ctx.JSON(http.StatusCreated, result)
}

// UpdateStudent handler to update an existing Student
func (c *StudentController) UpdateStudent(ctx echo.Context) error {
	flagValue := ctx.QueryParam("flag")
	flag, err := strconv.ParseBool(flagValue)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid flag value. Accepted true or false.",
		})
	}
	id := ctx.Param("id")
	var req Request.StudentRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	updatedStudent, err := c.Manager.UpdateStudent(flag, id, req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	log.Println("Updated student successfully")
	return ctx.JSON(http.StatusOK, updatedStudent)
}

// DeleteStudent handler to delete a Student
func (c *StudentController) DeleteStudent(ctx echo.Context) error {
	flagValue := ctx.QueryParam("flag")
	flag, err := strconv.ParseBool(flagValue)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid flag value. Accepted true or false.",
		})
	}
	id := ctx.Param("id")
	if id == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Student ID is required.",
		})
	}
	err = c.Manager.DeleteStudent(flag, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	log.Printf("Deleted student with ID %s successfully", id)
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Student deleted successfully",
	})
}
