package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"hrms/model"
	"hrms/resource"
	"log"
)

func AddAuthorityDetail(c *gin.Context, dto *model.AddAuthorityDetailDTO) error {
	var detail model.AuthorityDetail
	Transfer(&dto, &detail)
	if err := resource.HrmsDB(c).Create(&detail).Error; err != nil {
		log.Printf("AddAuthorityDetail err = %v", err)
		return err
	}
	return nil
}

func UpdateAuthorityDetailById(c *gin.Context, dto *model.UpdateAuthorityDetailDTO) error {
	var detail model.AuthorityDetail
	Transfer(&dto, &detail)
	if err := resource.HrmsDB(c).Where("id = ?", detail.ID).
		Updates(&detail).Error; err != nil {
		log.Printf("UpdateAuthorityDetailById err = %v", err)
		return err
	}
	return nil
}

func GetAuthorityDetailByUserTypeAndModel(c *gin.Context, detail *model.GetAuthorityDetailDTO) (string, error) {
	var authorityDetail model.AuthorityDetail
	if err := resource.HrmsDB(c).Where("user_type = ? and model = ?", detail.UserType, detail.Model).
		Find(&authorityDetail).Error; err != nil {
		log.Printf("GetAuthorityDetailByUserTypeAndModel err = %v", err)
		return "", err
	}
	return authorityDetail.AuthorityContent, nil
}

func GetAuthorityDetailListByUserType(c *gin.Context, userType string, start int, limit int) ([]*model.AuthorityDetail, int64, error) {
	var authorityDetailList []*model.AuthorityDetail
	var err error
	if start == -1 && limit == -1 {
		// 不加分页
		err = resource.HrmsDB(c).Where("user_type = ?", userType).Find(&authorityDetailList).Error
	} else {
		// 加分页
		err = resource.HrmsDB(c).Where("user_type = ?", userType).Offset(start).Limit(limit).Find(&authorityDetailList).Error
	}
	if err != nil {
		return nil, 0, err
	}
	var total int64
	resource.HrmsDB(c).Model(&model.AuthorityDetail{}).Count(&total)
	return authorityDetailList, total, nil
}

func SetAdminByStaffId(c *gin.Context, staffId string) error {
	authority := model.Authority{
		UserType: "sys",
	}
	if err := resource.HrmsDB(c).Where("staff_id = ?", staffId).Updates(&authority).Error; err != nil {
		log.Printf("SetAdminByStaffId err = %v", err)
		return err
	}
	return nil
}

func SetNormalByStaffId(c *gin.Context, staffId string) error {
	authority := model.Authority{
		UserType: "normal",
	}
	if err := resource.HrmsDB(c).Where("staff_id = ?", staffId).Updates(&authority).Error; err != nil {
		log.Printf("SetNormalByStaffId err = %v", err)
		return err
	}
	return nil
}

func GetNotificationByTitle(c *gin.Context, noticeTitle string, start int, limit int) ([]*model.Notification, int64, error) {
	var notifications []*model.Notification
	var err error
	if start == -1 && limit == -1 {
		// 不加分页
		if noticeTitle != "all" {
			err = resource.HrmsDB(c).Where("notice_title like ?", "%"+noticeTitle+"%").Order("date desc").Find(&notifications).Error
		} else {
			err = resource.HrmsDB(c).Order("date desc").Find(&notifications).Error
		}

	} else {
		// 加分页
		if noticeTitle != "all" {
			err = resource.HrmsDB(c).Where("notice_title like ?", "%"+noticeTitle+"%").Order("date desc").Offset(start).Limit(limit).Find(&notifications).Error
		} else {
			err = resource.HrmsDB(c).Order("date desc").Find(&notifications).Error
		}
	}
	if err != nil {
		return nil, 0, err
	}
	var total int64
	resource.HrmsDB(c).Model(&model.Notification{}).Count(&total)
	return notifications, total, nil
}

func CreateNotification(c *gin.Context, dto *model.NotificationDTO) error {
	var notification model.Notification
	Transfer(&dto, &notification)
	notification.NoticeId = RandomID("notice")
	notification.Date = Str2Time(dto.Date, 0)
	// 富文本内容base64编码(前端实现)
	//notification.NoticeContent = base64.StdEncoding.EncodeToString([]byte(dto.NoticeContent))
	if err := resource.HrmsDB(c).Create(&notification).Error; err != nil {
		log.Printf("CreateNotification err = %v", err)
		return err
	}
	return nil
}

func DelNotificationById(c *gin.Context, notice_id string) error {
	if err := resource.HrmsDB(c).Where("notice_id = ?", notice_id).Delete(&model.Notification{}).Error; err != nil {
		log.Printf("DelNotificationById err = %v", err)
		return err
	}
	return nil
}

func UpdateNotificationById(c *gin.Context, dto *model.NotificationEditDTO) error {
	var notification model.Notification
	Transfer(&dto, &notification)
	notification.Date = Str2Time(dto.Date, 0)
	if err := resource.HrmsDB(c).Where("id = ?", notification.ID).
		Updates(&notification).Error; err != nil {
		log.Printf("UpdateNotificationById err = %v", err)
		return err
	}
	return nil
}

func CreateSalary(c *gin.Context, dto *model.SalaryCreateDTO) error {
	var total int64
	resource.HrmsDB(c).Model(&model.Salary{}).Where("staff_id = ?", dto.StaffId).Count(&total)
	if total != 0 {
		return errors.New(fmt.Sprintf("该员工薪资数据已经存在"))
	}
	var salary model.Salary
	Transfer(&dto, &salary)
	salary.SalaryId = RandomID("salary")
	if err := resource.HrmsDB(c).Create(&salary).Error; err != nil {
		log.Printf("CreateSalary err = %v", err)
		return err
	}
	return nil
}

func DelSalaryBySalaryId(c *gin.Context, salaryId string) error {
	if err := resource.HrmsDB(c).Where("salary_id = ?", salaryId).Delete(&model.Salary{}).
		Error; err != nil {
		log.Printf("DelSalaryBySalaryId err = %v", err)
		return err
	}
	return nil
}

func UpdateSalaryById(c *gin.Context, dto *model.SalaryEditDTO) error {
	var salary model.Salary
	Transfer(&dto, &salary)
	if err := resource.HrmsDB(c).Where("id = ?", salary.ID).Updates(&salary).
		Error; err != nil {
		log.Printf("UpdateSalaryById err = %v", err)
		return err
	}
	return nil
}

func GetSalaryByStaffId(c *gin.Context, staffId string, start int, limit int) ([]*model.Salary, int64, error) {
	var salarys []*model.Salary
	var err error
	if start == -1 && limit == -1 {
		// 不加分页
		if staffId != "all" {
			err = resource.HrmsDB(c).Where("staff_id = ?", staffId).Find(&salarys).Error
		} else {
			err = resource.HrmsDB(c).Find(&salarys).Error
		}

	} else {
		// 加分页
		if staffId != "all" {
			err = resource.HrmsDB(c).Where("staff_id = ?", staffId).Offset(start).Limit(limit).Find(&salarys).Error
		} else {
			err = resource.HrmsDB(c).Find(&salarys).Error
		}
	}
	if err != nil {
		return nil, 0, err
	}
	var total int64
	resource.HrmsDB(c).Model(&model.Salary{}).Count(&total)
	return salarys, total, nil
}
