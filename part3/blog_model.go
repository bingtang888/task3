package part3

import (
	"gorm.io/gorm"
)

// User 模型（用户）
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(64);not null;unique;index:uni_users_username"` // 唯一索引
	Posts    []Post `gorm:"foreignKey:UserID"`                                         // 一对多关联
}

// Post 模型（文章）
type Post struct {
	gorm.Model
	Title    string    `gorm:"type:varchar(255);not null"`
	Content  string    `gorm:"type:text;not null"`
	UserID   uint      `gorm:"not null"`
	User     User      `gorm:"foreignKey:UserID"`
	Comments []Comment `gorm:"foreignKey:PostID"`
}

// Comment 模型（评论）
type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null"`
	PostID  uint   `gorm:"not null"`
	Post    Post   `gorm:"foreignKey:PostID"`
}

// InitBlogTables 初始化表结构
func InitBlogTables(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &Post{}, &Comment{})
}

// CreateTestData 添加测试数据（修复：物理删除+正确获取ID）
func CreateTestData(db *gorm.DB) error {
	// 1. 物理删除所有历史数据（彻底删除，解决软删除导致的唯一键冲突）
	// 顺序：评论→文章→用户（符合外键依赖）
	if err := db.Unscoped().Where("1 = 1").Delete(&Comment{}).Error; err != nil {
		return err
	}
	if err := db.Unscoped().Where("1 = 1").Delete(&Post{}).Error; err != nil {
		return err
	}
	if err := db.Unscoped().Where("1 = 1").Delete(&User{}).Error; err != nil {
		return err
	}

	// 2. 创建测试用户（物理删除后，用户名可重复创建）
	var testUser User
	testUser.Username = "tech_blogger"
	if err := db.Create(&testUser).Error; err != nil { // Create后直接获取ID（GORM会自动回写）
		return err
	}

	// 3. 创建测试文章（简化逻辑，Create后ID自动回写）
	var post1 Post
	post1.Title = "GORM关联查询实战"
	post1.Content = "本文介绍GORM的Preload、Joins等关联查询用法"
	post1.UserID = testUser.ID
	if err := db.Create(&post1).Error; err != nil {
		return err
	}

	var post2 Post
	post2.Title = "SQLx与GORM对比"
	post2.Content = "分析SQLx的类型安全映射与GORM的ORM便捷性"
	post2.UserID = testUser.ID
	if err := db.Create(&post2).Error; err != nil {
		return err
	}

	// 4. 创建测试评论（PostID已正确获取，无外键冲突）
	comments := []Comment{
		{Content: "太实用了，解决了我的嵌套查询问题！", PostID: post1.ID},
		{Content: "请问Preload可以嵌套多层吗？", PostID: post1.ID},
		{Content: "感谢分享，已收藏～", PostID: post1.ID},
		{Content: "SQLx的NamedExec也很好用", PostID: post2.ID},
		{Content: "期待更多对比内容！", PostID: post2.ID},
	}
	return db.Create(&comments).Error
}
