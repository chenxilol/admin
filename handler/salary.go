package handler

import (
	"github.com/gin-gonic/gin"
	"hrms/model"
	"hrms/service"
	"log"
)

func CreateSalary(c *gin.Context) {
	// 参数绑定
	var dto model.SalaryCreateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Printf("[CreateSalary] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5001,
			"result": err.Error(),
		})
		return
	}
	// 业务处理
	err := service.CreateSalary(c, &dto)
	if err != nil {
		log.Printf("[CreateSalary] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5002,
			"result": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 2000,
	})
}

func DelSalary(c *gin.Context) {
	// 参数绑定
	salaryId := c.Param("salary_id")
	// 业务处理
	err := service.DelSalaryBySalaryId(c, salaryId)
	if err != nil {
		log.Printf("[DelSalary] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5002,
			"result": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 2000,
	})
}

func UpdateSalaryById(c *gin.Context) {
	// 参数绑定
	var dto model.SalaryEditDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Printf("[UpdateSalaryById] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5001,
			"result": err.Error(),
		})
		return
	}
	// 业务处理
	err := service.UpdateSalaryById(c, &dto)
	if err != nil {
		log.Printf("[UpdateSalaryById] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5002,
			"result": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 2000,
	})
}

func GetSalaryByStaffId(c *gin.Context) {
	// 参数绑定
	staffId := c.Param("staff_id")
	start, limit := service.AcceptPage(c)
	// 业务处理
	list, total, err := service.GetSalaryByStaffId(c, staffId, start, limit)
	if err != nil {
		log.Printf("[GetSalaryByStaffId] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5000,
			"total":  0,
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 2000,
		"total":  total,
		"msg":    list,
	})
}