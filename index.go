package main

import (
	"net/http"
	"strconv" // Import thư viện strconv để sử dụng Atoi

	"github.com/gin-gonic/gin"
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

var students = []Student{
	{ID: 1, Name: "Nguyen Van A", Age: 20, Email: "nguyenvana@gmail.com"},
	{ID: 2, Name: "Le Thi B", Age: 21, Email: "lethib@gmail.com"},
	{ID: 3, Name: "Tran Van C", Age: 22, Email: "tranvanc@gmail.com"},
}

func getStudents(c *gin.Context) {
	c.JSON(http.StatusOK, students)
}

func getStudentDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
		return
	}

	for _, student := range students {
		if student.ID == id {
			c.JSON(http.StatusOK, student)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Sinh viên không tồn tại"})
}

func addStudent(c *gin.Context) {
	var newStudent Student
	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newStudent.ID = len(students) + 1
	students = append(students, newStudent)
	c.JSON(http.StatusOK, newStudent)
}

func updateStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
		return
	}

	var updatedStudent Student
	if err := c.ShouldBindJSON(&updatedStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, student := range students {
		if student.ID == id {
			students[i].Name = updatedStudent.Name
			students[i].Age = updatedStudent.Age
			students[i].Email = updatedStudent.Email
			c.JSON(http.StatusOK, students[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Sinh viên không tồn tại"})
}

func deleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
		return
	}

	for i, student := range students {
		if student.ID == id {
			students = append(students[:i], students[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Sinh viên đã được xóa"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Sinh viên không tồn tại"})
}

func main() {
	r := gin.Default()

	r.GET("/students", getStudents)
	r.GET("/student-detail/:id", getStudentDetail)
	r.POST("/student", addStudent)
	r.PUT("/student/:id", updateStudent)
	r.DELETE("/student/:id", deleteStudent)

	// Chạy server
	r.Run(":8080") // Server sẽ chạy tại localhost:8080
}
