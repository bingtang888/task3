package part1

import (
	"gorm.io/gorm"
)

// Student 学生结构体，对应数据库表
type Student struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"type:varchar(255);not null"`
	Age   int    `gorm:"not null"`
	Grade string `gorm:"type:varchar(255);not null"`
}

// StudentService 学生服务结构体，封装GORM操作
type StudentService struct {
	db *gorm.DB
}

// NewStudentService 创建学生服务实例
func NewStudentService(db *gorm.DB) *StudentService {
	return &StudentService{db: db}
}

// InitTable 初始化表结构（自动迁移）
func (s *StudentService) InitTable() error {
	// 自动根据Student结构体创建表
	return s.db.AutoMigrate(&Student{})
}

// CreateStudent 新增学生
func (s *StudentService) CreateStudent(student *Student) error {
	return s.db.Create(student).Error
}

// FindAdults 查询所有年龄大于18岁的学生
func (s *StudentService) FindAdults() ([]Student, error) {
	var students []Student
	result := s.db.Where("age > ?", 18).Find(&students)
	return students, result.Error
}

// UpdateGradeByName 根据姓名更新年级
func (s *StudentService) UpdateGradeByName(name, newGrade string) (int64, error) {
	result := s.db.Model(&Student{}).Where("name = ?", name).Update("grade", newGrade)
	return result.RowsAffected, result.Error
}

// DeleteByAgeLessThan 删除年龄小于指定值的学生
func (s *StudentService) DeleteByAgeLessThan(age int) (int64, error) {
	result := s.db.Where("age < ?", age).Delete(&Student{})
	return result.RowsAffected, result.Error
}
