package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

var db *gorm.DB

func main() {
	r := gin.Default()

	// Kết nối tới MySQL
	dsn := "root:passroot@tcp(127.0.0.1:3308)/student_management?parseTime=true&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Student{})
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/students-data", func(c *gin.Context) {
		var students []Student
		db.Find(&students)
		c.JSON(http.StatusOK, students)
	})

	// API thêm sinh viên
	r.POST("/student", func(c *gin.Context) {
		var student Student
		if err := c.ShouldBindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&student)
		c.JSON(http.StatusOK, student)
	})

	// API cập nhật sinh viên
	r.PUT("/student/:id", func(c *gin.Context) {
		var student Student
		id := c.Param("id")
		if err := db.First(&student, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
			return
		}

		if err := c.ShouldBindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Save(&student)
		c.JSON(http.StatusOK, student)
	})

	// API xóa sinh viên
	r.DELETE("/student/:id", func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&Student{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete student"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
	})

	// Chạy server trên cổng 8080
	r.Run(":8080")
}
